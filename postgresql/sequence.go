package postgresql

import (
	"bytes"
	"fmt"
)

type PGSequence struct {
	Schemaname    string
	Sequencename  string
	Sequenceowner string
	StartValue    int
	MinValue      int
	MaxValue      int
	IncrementBy   int
}

func (p *PGSequence) GetMin() string {
	if p.MinValue == 1 {
		return "NO MINVALUE"
	}
	return fmt.Sprintf("MINVALUE %d", p.MinValue)
}

func (p *PGSequence) GetMax() string {
	if p.MaxValue == 9223372036854775807 {
		return "NO MAXVALUE"
	}
	return fmt.Sprintf("MAXVALUE %d", p.MaxValue)
}

func (p *PGSequence) Create() string {
	return fmt.Sprintf(
		"CREATE SEQUENCE %s.%s START WITH %d INCREMENT BY %d %s %s CACHE 1;\n\n",
		p.Schemaname,
		p.Sequencename,
		p.StartValue,
		p.IncrementBy,
		p.GetMin(),
		p.GetMax(),
	)
}

func (p *PGSequence) Drop() string {
	return fmt.Sprintf(
		"DROP SEQUENCE %s.%s;\n",
		p.Schemaname,
		p.Sequencename,
	)
}

func (p PGSequence) Diff(src *[]PGSequence) string {
	buf := bytes.NewBuffer([]byte{})
	dest := PGSequence{}
	for _, item := range *src {
		if item.Sequencename == p.Sequencename {
			dest = item
		}
	}
	if dest.Sequencename == "" {
		buf.WriteString(p.Create())
	} else if p.Create() != dest.Create() {
		buf.WriteString(dest.Drop())
		buf.WriteString(p.Create())
	}
	return buf.String()
}
