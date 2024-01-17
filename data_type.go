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
	DataTypeMacaddr
	DataTypeArrayInt8
	DataTypeArrayString
	DataTypeArrayMacaddr
	DataTypeBps
	DataTypeMaskedString
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
	DataTypeMacaddr,
	DataTypeArrayInt8,
	DataTypeArrayString,
	DataTypeArrayMacaddr,
	DataTypeBps,
	DataTypeMaskedString,
}

// 指定の文字列からデータ型を取得
func FindDataType(src string, count int) DataType {
	defer func() {
		recover()
	}()
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
	case DataTypeMacaddr:
		return "string"
	case DataTypeArrayInt8:
		return "[]int"
	case DataTypeArrayString:
		return "[]string"
	case DataTypeArrayMacaddr:
		return "[]string"
	case DataTypeBps:
		return "*bps.Bps"
	case DataTypeMaskedString:
		return "maskedstring.MaskedString"
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
	case DataTypeBps:
		return fmt.Sprintf("varchar(%d)", 64)
	case DataTypeMaskedString:
		return fmt.Sprintf("varchar(%d)", count)
	case DataTypeMacaddr, DataTypeArrayInt8, DataTypeArrayString, DataTypeArrayMacaddr:
		panic(fmt.Errorf("サポートしないデータ型: %d, %s", p, p.String()))
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
		return "int2"
	case DataTypeInt16:
		return "int2"
	case DataTypeInt32:
		return "int4"
	case DataTypeInt64:
		return "int8"
	case DataTypeUint8:
		return "int2"
	case DataTypeUint16:
		return "int2"
	case DataTypeUint32:
		return "int4"
	case DataTypeUint64:
		return "int8"
	case DataTypeFloat64:
		return "decimal"
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
	case DataTypeMacaddr:
		return "macaddr"
	case DataTypeArrayInt8:
		return "_int4"
	case DataTypeArrayString:
		return "_varchar"
	case DataTypeArrayMacaddr:
		return "_macaddr"
	case DataTypeBps:
		return fmt.Sprintf("varchar(%d)", 64)
	case DataTypeMaskedString:
		return fmt.Sprintf("varchar(%d)", count)
	default:
		panic(fmt.Errorf("未定義のデータタイプ: %d", p))
	}
}

// Golang用のtype出力
func (p DataType) Go() string {
	switch p {
	case DataTypeUUID:
		return "DataTypeUUID"
	case DataTypeInt8:
		return "DataTypeInt8"
	case DataTypeInt16:
		return "DataTypeInt16"
	case DataTypeInt32:
		return "DataTypeInt32"
	case DataTypeInt64:
		return "DataTypeInt64"
	case DataTypeUint8:
		return "DataTypeUint8"
	case DataTypeUint16:
		return "DataTypeUint16"
	case DataTypeUint32:
		return "DataTypeUint32"
	case DataTypeUint64:
		return "DataTypeUint64"
	case DataTypeFloat64:
		return "DataTypeFloat64"
	case DataTypeString:
		return "DataTypeString"
	case DataTypeFixString:
		return "DataTypeFixString"
	case DataTypeText64K:
		return "DataTypeText64K"
	case DataTypeText16M:
		return "DataTypeText16M"
	case DataTypeText4G:
		return "DataTypeText4G"
	case DataTypeBlob64K:
		return "DataTypeBlob64K"
	case DataTypeBlob16M:
		return "DataTypeBlob16M"
	case DataTypeBlob4G:
		return "DataTypeBlob4G"
	case DataTypeBytes:
		return "DataTypeBytes"
	case DataTypeJson:
		return "DataTypeJson"
	case DataTypeTimestamp:
		return "DataTypeTimestamp"
	case DataTypeDatetime:
		return "DataTypeDatetime"
	case DataTypeSmallSerial:
		return "DataTypeSmallSerial"
	case DataTypeSerial:
		return "DataTypeSerial"
	case DataTypeBigSerial:
		return "DataTypeBigSerial"
	case DataTypeMacaddr:
		return "DataTypeMacaddr"
	case DataTypeArrayInt8:
		return "DataTypeArrayInt8"
	case DataTypeArrayString:
		return "DataTypeArrayString"
	case DataTypeArrayMacaddr:
		return "DataTypeArrayMacaddr"
	case DataTypeBps:
		return "DataTypeBps"
	case DataTypeMaskedString:
		return "DataTypeMaskedString"
	default:
		panic(fmt.Errorf("未定義のデータタイプ: %d", p))
	}
}
