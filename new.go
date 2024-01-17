package catwalk

import (
	"fmt"

	"github.com/h-nosaka/catwalk/base"
)

func NewSchema(name string, mode SchemaMode, tables ...Table) *Schema {
	return &Schema{
		Name:   name,
		Mode:   mode,
		Tables: tables,
	}
}

func NewTable(schema string, name string, code JsonCase, comment string) *Table {
	return &Table{
		Schema:   schema,
		Name:     name,
		JsonCase: code,
		Comment:  base.String(comment),
	}
}

func (p *Table) SetColumns(src ...Column) *Table {
	p.Columns = src
	return p
}

func (p *Table) SetDefaultColumns(primary DataType, src ...Column) *Table {
	columns := []Column{}
	mode := GetSchemaMode()
	switch mode {
	case SchemaModeMySQL:
		if primary == DataTypeUUID {
			columns = append(columns, NewColumn("id", DataTypeUUID, 0, false, "ID").SetDefault("uuid()").Done())
		} else {
			columns = append(columns, NewColumn("id", DataTypeUint64, 0, false, "ID").SetExtra("auto_increment").Done())
		}
	case SchemaModePostgres:
		columns = append(columns, NewColumn("id", DataTypeUUID, 0, false, "ID").SetDefault("uuid_generate_v4()").Done())
	}
	columns = append(columns, src...)
	switch mode {
	case SchemaModeMySQL:
		columns = append(columns,
			NewColumn("created_at", DataTypeTimestamp, 0, false, "作成日").SetDefault("current_timestamp()").Done(),
			NewColumn("updated_at", DataTypeTimestamp, 0, false, "更新日").SetDefault("current_timestamp()").Done(),
		)
	case SchemaModePostgres:
		columns = append(columns,
			NewColumn("created_at", DataTypeTimestamp, 0, false, "作成日").SetDefault("(now())::timestamp(0) without time zone").Done(),
			NewColumn("updated_at", DataTypeTimestamp, 0, false, "更新日").SetDefault("(now())::timestamp(0) without time zone").Done(),
		)
	}
	return p.SetColumns(columns...)
}

func (p *Table) SetIndexes(src ...Index) *Table {
	p.Indexes = src
	return p
}

func (p *Table) SetDefaultIndexes(src ...Index) *Table {
	indexes := []Index{}
	switch GetSchemaMode() {
	case SchemaModeMySQL:
		indexes = append(indexes,
			NewIndex("PRIMARY", IndexTypePrimary, "id"),
		)
	case SchemaModePostgres:
		indexes = append(indexes,
			NewIndex(fmt.Sprintf("%s_primary_idx", p.Name), IndexTypePrimary, "id"),
		)
	}
	indexes = append(indexes, src...)
	return p.SetIndexes(indexes...)
}

func (p *Table) SetRelations(relations ...Relation) *Table {
	p.Relations = relations
	return p
}

func (p *Table) SetEnums(enums ...Enum) *Table {
	p.Enums = enums
	return p
}

func (p *Table) SetPartitions(partition *Partition) *Table {
	p.Partitions = partition
	return p
}

func (p *Table) Done() Table {
	return *p
}

func NewColumn(name string, data DataType, count int, nullable bool, comment string) *Column {
	return &Column{
		Name:     name,
		DataType: data,
		Count:    count,
		Null:     nullable,
		Comment:  base.String(comment),
	}
}

func (p *Column) SetDefault(src string) *Column {
	p.Default = base.String(src)
	return p
}

func (p *Column) SetExtra(src string) *Column {
	p.Extra = base.String(src)
	return p
}

func (p *Column) SetRename(src string) *Column {
	p.Rename = base.String(src)
	return p
}

func (p *Column) Done() Column {
	return *p
}

func NewIndex(name string, data IndexType, columns ...string) Index {
	return Index{
		Name:    name,
		Type:    string(data),
		Columns: columns,
	}
}

func NewRelation(name string, column string, refTable string, refColumn string, one bool, isFK bool) Relation {
	return Relation{
		Name:         name,
		Column:       column,
		RefTable:     refTable,
		RefColumn:    refColumn,
		HasOne:       one,
		HasAny:       !one,
		IsForeignKey: isFK,
	}
}

func NewEnum(column string, data EnumType, values ...EnumValue) Enum {
	return Enum{
		Column: column,
		Type:   data,
		Values: values,
	}
}
