package base

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strconv"
	"strings"
)

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func GetEnvByte(key string) *[]byte {
	value, ok := os.LookupEnv(key)
	if !ok {
		return nil
	}
	buf := []byte(strings.ReplaceAll(value, "\\n", "\n"))
	return &buf
}

func GetEnvBool(key string, fallback bool) bool {
	value, ok := os.LookupEnv(key)
	if ok {
		return value == "true" || value == "True" || value == "ok" || value == "OK" || value == "1"
	}
	return fallback
}

func GetEnvInt(key string, fallback int) int {
	value, ok := os.LookupEnv(key)
	if ok {
		rs, _ := strconv.Atoi(value)
		return rs
	}
	return fallback
}

func String(src string) *string {
	return &src
}

func Bool(src bool) *bool {
	return &src
}

func ReadBuffer(r io.Reader) string {
	var b bytes.Buffer
	if _, err := b.ReadFrom(r); err != nil {
		return ""
	}
	return b.String()
}

func ToJson(src interface{}, def string, pretty ...bool) string {
	var rs []byte
	var err error
	if len(pretty) > 0 && pretty[0] {
		rs, err = json.MarshalIndent(src, "", "\t")
	} else {
		rs, err = json.Marshal(src)
	}
	if err != nil {
		return def
	}
	return string(rs)
}

func ReJson(src string) string {
	var data interface{}
	if err := json.Unmarshal([]byte(src), &data); err != nil {
		Log.Error(err.Error())
		return src
	}
	return ToJson(data, src, true)
}
