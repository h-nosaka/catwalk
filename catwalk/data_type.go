package catwalk

import "fmt"

type DataType int

const (
	DataTypeUUID DataType = iota
	DataTypeInt16
	DataTypeInt8
	DataTypeInt32
	DataTypeInt64
	DataTypeUint16
	DataTypeUint8
	DataTypeUint32
	DataTypeUint64
	DataTypeFloat64
	DataTypeString
	DataTypeFixString
	DataTypeText4G
	DataTypeText16M
	DataTypeText64K
	DataTypeBlob4G
	DataTypeBlob16M
	DataTypeBlob64K
	DataTypeBytes
	DataTypeJson
	DataTypeTimestamp
	DataTypeDatetime
	DataTypeSmallSerial
	DataTypeSerial
	DataTypeBigSerial
)

var dataTypes = []DataType{
	DataTypeUUID,
	DataTypeInt8,
	DataTypeInt16,
	DataTypeInt32,
	DataTypeInt64,
	DataTypeUint8,
	DataTypeUint16,
	DataTypeUint32,
	DataTypeUint64,
	DataTypeFloat64,
	DataTypeString,
	DataTypeFixString,
	DataTypeText64K,
	DataTypeText16M,
	DataTypeText4G,
	DataTypeBlob64K,
	DataTypeBlob16M,
	DataTypeBlob4G,
	DataTypeBytes,
	DataTypeJson,
	DataTypeTimestamp,
	DataTypeDatetime,
	DataTypeSmallSerial,
	DataTypeSerial,
	DataTypeBigSerial,
}

// 指定の文字列からデータ型を取得
func FindDataType(src string, count int) DataType {
	for _, d := range dataTypes {
		if d.Mysql(count) == src || d.Postgres(0) == src {
			return d
		}
	}
	panic(fmt.Errorf("未定義のデータタイプ: %s", src))
}

// golangにおけるデータ型を生成
func (p DataType) String() string {
	switch p {
	case DataTypeUUID:
		return "string"
	case DataTypeInt8:
		return "int8"
	case DataTypeInt16:
		return "int16"
	case DataTypeInt32:
		return "int"
	case DataTypeInt64:
		return "int64"
	case DataTypeUint8:
		return "uint8"
	case DataTypeUint16:
		return "uint16"
	case DataTypeUint32:
		return "uint"
	case DataTypeUint64:
		return "uint64"
	case DataTypeFloat64:
		return "float64"
	case DataTypeString:
		return "string"
	case DataTypeFixString:
		return "string"
	case DataTypeText64K:
		return "string"
	case DataTypeText16M:
		return "string"
	case DataTypeText4G:
		return "string"
	case DataTypeBlob64K:
		return "[]byte"
	case DataTypeBlob16M:
		return "[]byte"
	case DataTypeBlob4G:
		return "[]byte"
	case DataTypeBytes:
		return "[]byte"
	case DataTypeJson:
		return "json.RawMessage"
	case DataTypeTimestamp:
		return "time.Time"
	case DataTypeDatetime:
		return "time.Time"
	case DataTypeSmallSerial:
		return "int16"
	case DataTypeSerial:
		return "int"
	case DataTypeBigSerial:
		return "int64"
	default:
		panic(fmt.Errorf("未定義のデータタイプ: %d", p))
	}
}

// mysqlにおけるデータ型の文字列を生成
func (p DataType) Mysql(counts ...int) string {
	count := 0
	if len(counts) > 0 {
		count = counts[0]
	}
	switch p {
	case DataTypeUUID:
		return "uuid"
	case DataTypeInt8:
		if count > 0 {
			return fmt.Sprintf("tinyint(%d)", count)
		}
		return "tinyint"
	case DataTypeInt16:
		if count > 0 {
			return fmt.Sprintf("smallint(%d)", count)
		}
		return "smallint"
	case DataTypeInt32:
		if count > 0 {
			return fmt.Sprintf("int(%d)", count)
		}
		return "int"
	case DataTypeInt64:
		if count > 0 {
			return fmt.Sprintf("bigint(%d)", count)
		}
		return "bigint"
	case DataTypeUint8:
		if count > 0 {
			return fmt.Sprintf("tinyint(%d) unsigned", count)
		}
		return "tinyint unsigned"
	case DataTypeUint16:
		if count > 0 {
			return fmt.Sprintf("smallint(%d) unsigned", count)
		}
		return "smallint unsigned"
	case DataTypeUint32:
		if count > 0 {
			return fmt.Sprintf("int(%d) unsigned", count)
		}
		return "int unsigned"
	case DataTypeUint64:
		if count > 0 {
			return fmt.Sprintf("bigint(%d) unsigned", count)
		}
		return "bigint unsigned"
	case DataTypeFloat64:
		return "double"
	case DataTypeString:
		if count > 0 {
			return fmt.Sprintf("varchar(%d)", count)
		}
		return "varchar"
	case DataTypeFixString:
		if count > 0 {
			return fmt.Sprintf("char(%d)", count)
		}
		return "char"
	case DataTypeText64K:
		return "text"
	case DataTypeText16M:
		return "mediumtext"
	case DataTypeText4G:
		return "longtext"
	case DataTypeBlob64K:
		return "blob"
	case DataTypeBlob16M:
		return "mediumblob"
	case DataTypeBlob4G:
		return "longblob"
	case DataTypeBytes:
		return "binary"
	case DataTypeJson:
		return "json"
	case DataTypeTimestamp:
		return "timestamp"
	case DataTypeDatetime:
		return "datetime"
	case DataTypeSmallSerial:
		return "uuid"
	case DataTypeSerial:
		return "uuid"
	case DataTypeBigSerial:
		return "uuid"
	default:
		panic(fmt.Errorf("未定義のデータタイプ: %d", p))
	}
}

// postgresにおけるデータ型の文字列を生成
func (p DataType) Postgres(counts ...int) string {
	count := 0
	if len(counts) > 0 {
		count = counts[0]
	}
	switch p {
	case DataTypeUUID:
		return "uuid"
	case DataTypeInt8:
		if count > 0 {
			return fmt.Sprintf("smallint(%d)", count)
		}
		return "smallint"
	case DataTypeInt16:
		if count > 0 {
			return fmt.Sprintf("smallint(%d)", count)
		}
		return "smallint"
	case DataTypeInt32:
		if count > 0 {
			return fmt.Sprintf("int(%d)", count)
		}
		return "int"
	case DataTypeInt64:
		if count > 0 {
			return fmt.Sprintf("bigint(%d)", count)
		}
		return "bigint"
	case DataTypeUint8:
		return "int2"
	case DataTypeUint16:
		return "int2"
	case DataTypeUint32:
		return "int4"
	case DataTypeUint64:
		return "int8"
	case DataTypeFloat64:
		return "double"
	case DataTypeString:
		if count > 0 {
			return fmt.Sprintf("varchar(%d)", count)
		}
		return "varchar"
	case DataTypeFixString:
		if count > 0 {
			return fmt.Sprintf("character(%d)", count)
		}
		return "character"
	case DataTypeText64K:
		return "text"
	case DataTypeText16M:
		return "text"
	case DataTypeText4G:
		return "text"
	case DataTypeBlob64K:
		return "bytea"
	case DataTypeBlob16M:
		return "bytea"
	case DataTypeBlob4G:
		return "bytea"
	case DataTypeBytes:
		return "bytea"
	case DataTypeJson:
		return "json"
	case DataTypeTimestamp:
		return "timestamp"
	case DataTypeDatetime:
		return "timestamp"
	case DataTypeSmallSerial:
		return "smallserial"
	case DataTypeSerial:
		return "serial"
	case DataTypeBigSerial:
		return "bigserial"
	default:
		panic(fmt.Errorf("未定義のデータタイプ: %d", p))
	}
}
