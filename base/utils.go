package base

import (
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

func WriteFile(filename string, data []byte) {
	fp, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	if _, err := fp.Write(data); err != nil {
		panic(err)
	}
}
