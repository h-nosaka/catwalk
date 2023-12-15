package catwalk

import (
	"regexp"
	"strconv"

	"github.com/h-nosaka/catwalk/base"
)

func (p *Table) GetColumn() {
	data := []map[string]interface{}{}
	switch GetSchemaMode() {
	case SchemaModeMySQL:
		data = GetData("Columns", p.Schema, p.Name)
	case SchemaModePostgres:
		data = GetData("Columns", p.Name, p.Name)
	}
	for _, record := range data {
		// fmt.Println(base.ToPrettyJson(record))
		ctype := record["COLUMN_TYPE"].(string)
		count := 0
		counts := regexp.MustCompile(`\((\d+)\)`).FindStringSubmatch(ctype)
		if len(counts) == 2 {
			count, _ = strconv.Atoi(counts[1])
		}
		if GetSchemaMode() == SchemaModePostgres {
			if record["character_maximum_length"] != nil && record["character_maximum_length"].(int32) > 0 {
				count = int(record["character_maximum_length"].(int32))
			}
			if record["numeric_precision"] != nil && record["numeric_precision"].(int32) > 0 {
				count = int(record["numeric_precision"].(int32))
			}
		}
		dtype := FindDataType(ctype, count)
		var def *string
		if record["COLUMN_DEFAULT"] != nil && record["COLUMN_DEFAULT"].(string) != "" {
			def = base.String(record["COLUMN_DEFAULT"].(string))
		}
		var extra *string
		if record["EXTRA"] != nil && record["EXTRA"].(string) != "" {
			extra = base.String(record["EXTRA"].(string))
		}
		var comment *string
		if record["COLUMN_COMMENT"] != nil && record["COLUMN_COMMENT"].(string) != "" {
			comment = base.String(record["COLUMN_COMMENT"].(string))
		}
		col := Column{
			Name:     record["COLUMN_NAME"].(string),
			DataType: dtype,
			Count:    count,
			Extra:    extra,
			Default:  def,
			Null:     record["IS_NULLABLE"].(string) == "YES",
			Comment:  comment,
		}
		p.Columns = append(p.Columns, col)
	}
}

func (p *Table) GetIndexes() {
	data := []map[string]interface{}{}
	switch GetSchemaMode() {
	case SchemaModeMySQL:
		data = GetData("Indexes", p.Schema, p.Name)
	case SchemaModePostgres:
		data = GetData("Indexes", p.Name)
	}
	for _, record := range data {
		// fmt.Println(base.ToPrettyJson(record))
		ctype := ""
		if record["NON_UNIQUE"] != nil && record["NON_UNIQUE"].(int64) == 0 {
			ctype = "UNIQUE"
		}
		if record["INDEX_NAME"].(string) == "PRIMARY" {
			ctype = "PRIMARY KEY"
		}
		if record["TYPE"] != nil && record["TYPE"].(string) == "PRIMARY KEY" {
			ctype = "PRIMARY KEY"
		}
		index := Index{
			Name: record["INDEX_NAME"].(string),
			Type: ctype,
		}
		index.GetColumn(p)
		p.Indexes = append(p.Indexes, index)
	}
}

func (p *Index) GetColumn(table *Table) {
	p.Columns = []string{}
	data := []map[string]interface{}{}
	switch GetSchemaMode() {
	case SchemaModeMySQL:
		data = GetData("IndexColumn", table.Schema, table.Name, p.Name)
	case SchemaModePostgres:
		data = GetData("IndexColumn", p.Name)
	}
	for _, record := range data {
		// fmt.Println(base.ToPrettyJson(record))
		p.Columns = append(p.Columns, record["COLUMN_NAME"].(string))
	}
}

func (p *Table) GetForeignkeys() {
	data := []map[string]interface{}{}
	switch GetSchemaMode() {
	case SchemaModeMySQL:
		data = GetData("Foreignkeys", p.Schema, p.Name)
	case SchemaModePostgres:
		data = GetData("Foreignkeys", p.Name)
	}
	for _, record := range data {
		// fmt.Println(base.ToPrettyJson(record))
		relation := Relation{
			Name:         record["CONSTRAINT_NAME"].(string),
			Column:       record["COLUMN_NAME"].(string),
			RefTable:     record["REFERENCED_TABLE_NAME"].(string),
			RefColumn:    record["REFERENCED_COLUMN_NAME"].(string),
			IsForeignKey: true,
		}
		p.Relations = append(p.Relations, relation)
	}
}

func (p *Table) GetPartitions() {
	data := []map[string]interface{}{}
	switch GetSchemaMode() {
	case SchemaModeMySQL:
		data = GetData("Partitions", p.Schema, p.Name)
	}
	if len(data) > 0 && data[0]["PARTITION_METHOD"] != nil {
		p.Partitions = &Partition{
			Type:   data[0]["PARTITION_METHOD"].(string),
			Column: data[0]["PARTITION_EXPRESSION"].(string),
			Keys:   []PartitionKey{},
		}
		for _, item := range data {
			p.Partitions.Keys = append(p.Partitions.Keys, PartitionKey{Key: item["PARTITION_NAME"].(string), Value: item["PARTITION_DESCRIPTION"].(string)})
		}
	}
}
