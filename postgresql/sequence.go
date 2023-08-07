package postgresql

import (
	"bytes"
	"fmt"
)

type ISequence struct {
	Sequencename string
	StartValue   int
	MinValue     int
	MaxValue     int
	IncrementBy  int
}

func NewSeq(name string, start int, min int, max int, inc int) ISequence {
	return ISequence{
		Sequencename: name,
		StartValue:   start,
		MinValue:     min,
		MaxValue:     max,
		IncrementBy:  inc,
	}
}

func (p *ISequence) GetMin() string {
	if p.MinValue == 1 {
		return "NO MINVALUE"
	}
	return fmt.Sprintf("MINVALUE %d", p.MinValue)
}

func (p *ISequence) GetMax() string {
	if p.MaxValue == 9223372036854775807 {
		return "NO MAXVALUE"
	}
	return fmt.Sprintf("MAXVALUE %d", p.MaxValue)
}

func (p *ISequence) Create(t *ITable) string {
	return fmt.Sprintf(
		"CREATE SEQUENCE %s%s START WITH %d INCREMENT BY %d %s %s CACHE 1;\n\n",
		t.SchemaName(),
		p.Sequencename,
		p.StartValue,
		p.IncrementBy,
		p.GetMin(),
		p.GetMax(),
	)
}

func (p *ISequence) Drop(t *ITable) string {
	return fmt.Sprintf(
		"DROP SEQUENCE %s%s;\n",
		t.SchemaName(),
		p.Sequencename,
	)
}

func (p *ISequence) Diff(t *ITable, src *[]ISequence) string {
	buf := bytes.NewBuffer([]byte{})
	dest := ISequence{}
	for _, item := range *src {
		if item.Sequencename == p.Sequencename {
			dest = item
		}
	}
	if dest.Sequencename == "" {
		buf.WriteString(p.Create(t))
	} else if p.Create(t) != dest.Create(t) {
		buf.WriteString(dest.Drop(t))
		buf.WriteString(p.Create(t))
	}
	return buf.String()
}
