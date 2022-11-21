package base

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"strings"
	"time"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/tidwall/gjson"
)

var ES *elasticsearch.Client

func ESInit() {
	insecure := GetEnv("ELASTICSEARCH_CERT", "0") == "1"
	rootCAs, _ := x509.SystemCertPool()
	esConf := elasticsearch.Config{
		Addresses: strings.Split(GetEnv("ELASTICSEACH_HOST", "es"), ","),
		Username:  GetEnv("ELASTICSEACH_USER", "docker"),
		Password:  GetEnv("ELASTICSEACH_PASSWORD", "password"),
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   GetEnvInt("ELASTICSEACH_MAXCONN", 10),
			ResponseHeaderTimeout: time.Second,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion:         tls.VersionTLS11,
				InsecureSkipVerify: insecure,
				RootCAs:            rootCAs,
			},
		},
	}
	file := GetEnvByte("ELASTIC_CAFILE")
	if file != nil && len(*file) > 0 {
		esConf.CACert = *file
	}
	es, err := elasticsearch.NewClient(esConf)
	if err != nil {
		panic(err)
	}
	ES = es
}

func ESIndexExist(index string) bool {
	rs, err := ES.Indices.Exists([]string{index})
	if err != nil {
		return false
	}
	return rs.StatusCode == 404
}

func ESIndexCreateIf(index string) error {
	if ESIndexExist(index) {
		body, _ := json.Marshal(map[string]interface{}{
			"settings": map[string]interface{}{
				"index": map[string]interface{}{
					"number_of_shards":   3,
					"number_of_replicas": 2,
				},
			},
		})
		if _, err := ES.Indices.Create(index, ES.Indices.Create.WithBody(bytes.NewReader(body))); err != nil {
			return err
		}
	}
	return nil
}

func ESIndexDrop(index string) error {
	if !ESIndexExist(index) {
		if _, err := ES.Indices.Delete([]string{index}); err != nil {
			return err
		}
	}
	return nil
}

func ESUpsert(index string, src interface{}) (*esapi.Response, error) {
	if err := ESIndexCreateIf(index); err != nil {
		return nil, err
	}
	doc := []byte(ToJson(src, "{}"))
	return ES.Index(index, bytes.NewReader(doc), ES.Index.WithRefresh("true"))
}

func ESSearch(result *[]interface{}, index string, query interface{}) (int, string, error) {
	var body []byte
	if reflect.ValueOf(query).Type().Name() == "string" {
		body = []byte(query.(string))
	} else {
		data := ToJson(query, "{}")
		body = []byte(data)
	}
	// indexが存在しない場合はSearchで404エラーになるため作成しておく
	if err := ESIndexCreateIf(index); err != nil {
		return 0, "", fmt.Errorf("ES error: %v", err)
	}
	rs, err := ES.Search(
		ES.Search.WithIndex(index),
		ES.Search.WithBody(bytes.NewReader(body)),
		ES.Search.WithFrom(0),
		ES.Search.WithSize(10),
		ES.Search.WithScroll(time.Minute),
	)
	if err != nil {
		return 0, "", fmt.Errorf("ES error: %v", err)
	}
	if rs.StatusCode != 200 {
		return 0, "", fmt.Errorf("ES error status code: %v", rs.StatusCode)
	}
	data := ReadBuffer(rs.Body)
	rs.Body.Close()
	hits := gjson.Get(data, "hits.hits").Array()
	for _, rec := range hits {
		var item map[string]interface{}
		if err := json.Unmarshal([]byte(rec.String()), &item); err != nil {
			Log.Error(err.Error())
		}
		*result = append(*result, item["_source"])
	}
	return len(hits), gjson.Get(data, "_scroll_id").String(), nil
}

func ESSearchNext(result *[]interface{}, scrollId string) error {
	rs, err := ES.Scroll(
		ES.Scroll.WithScrollID(scrollId),
		ES.Scroll.WithScroll(time.Minute),
	)
	if err != nil {
		return err
	}
	data := ReadBuffer(rs.Body)
	rs.Body.Close()
	hits := gjson.Get(data, "hits.hits").Array()
	for _, rec := range hits {
		var item map[string]interface{}
		if err := json.Unmarshal([]byte(rec.String()), &item); err != nil {
			Log.Error(err.Error())
		}
		*result = append(*result, item["_source"])
	}
	return nil
}

func ESSearchAll(result *[]interface{}, index string, query interface{}) error {
	hits, scrollId, err := ESSearch(result, index, query)
	if err != nil {
		return err
	}
	for hits > 0 && len(scrollId) > 0 {
		if err := ESSearchNext(result, scrollId); err != nil {
			return err
		}
	}
	return nil
}
