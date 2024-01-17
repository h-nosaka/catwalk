package catwalk

import (
	"bytes"
	"fmt"
	"strings"
)

func (p *Partition) Create(t *Table) string {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(GetParts("PartitionCreate", t.SchemaTable(), p.Type, p.Column))
	parts := []string{}
	for _, item := range p.Keys {
		switch p.Type {
		case "RANGE":
			parts = append(parts, GetParts("PartitionCreateRange", item.Key, item.Value))
		case "LIST":
			parts = append(parts, GetParts("PartitionCreateList", item.Key, item.Value))
		}
	}
	buf.WriteString(fmt.Sprintf("%s\n);\n\n", strings.Join(parts, ",\n")))
	return buf.String()
}

func (p *PartitionKey) Drop(t *Table) string {
	return GetParts("PartitionDrop", t.SchemaTable(), p.Key)
}

func (p *Partition) Diff(t *Table, src *Partition) string {
	buf := bytes.NewBuffer([]byte{})
	// 新規パーティションテーブルの場合
	if p != nil && src == nil {
		buf.WriteString(p.Create(t))
		return buf.String()
	}
	// パーティション全体削除の場合
	if p == nil && src != nil {
		buf.WriteString(fmt.Sprintf("ALTER TABLE %s REMOVE PARTITIONING;\n\n", t.SchemaTable()))
		return buf.String()
	}
	// パーティション追加
	for _, item := range p.Keys {
		ok := false
		for _, dest := range src.Keys {
			if item.Key == dest.Key {
				ok = true
				break
			}
		}
		if !ok {
			switch p.Type {
			case "RANGE":
				buf.WriteString(fmt.Sprintf("ALTER TABLE %s ADD PARTITION (PARTITION %s VALUES LESS THAN (%s));\n", t.SchemaTable(), item.Key, item.Value))
			case "LIST":
				buf.WriteString(fmt.Sprintf("ALTER TABLE %s ADD PARTITION (PARTITION %s VALUES IN (%s));\n", t.SchemaTable(), item.Key, item.Value))
			}
		}
	}
	// パーティション削除
	for _, dest := range src.Keys {
		ok := false
		for _, item := range p.Keys {
			if dest.Key == item.Key {
				ok = true
				break
			}
		}
		if !ok {
			buf.WriteString(fmt.Sprintf("ALTER TABLE %s TRUNCATE PARTITION %s;\n", t.SchemaTable(), dest.Key)) // パーティションの中身をクリアする
			buf.WriteString(fmt.Sprintf("ALTER TABLE %s DROP PARTITION %s;\n", t.SchemaTable(), dest.Key))
		}
	}
	buf.WriteString("\n")
	return buf.String()
}
