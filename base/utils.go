package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
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
	if len(src) == 0 {
		return nil
	}
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

func ToPrettyJson(src interface{}) string {
	rs, err := json.MarshalIndent(src, "", "  ")
	if err != nil {
		return "{}"
	}
	return string(rs)
}

func Diff(src string, dest string) string {
	fmt.Println(dest, src)
	return string(_Diff("dest.txt", []byte(dest), "src.txt", []byte(src)))
}

func Recover() {
	if rec := recover(); rec != nil {
		i := 1
		_, file, line, ok := runtime.Caller(i)
		for ok && filepath.Base(file) == "utils.go" {
			i++
			_, file, line, _ = runtime.Caller(i)
		}
		panic(fmt.Sprintf("%s:%d %s", file, line, rec)) // ログ取得したらfiberのエラーハンドリングに任せる
	}
}
