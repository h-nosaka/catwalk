package postgresql

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type PGSchema struct {
	Sequences   []PGSequence
	Tables      []PGTable
	Indexes     []PGIndex
	Foreignkeys []PGForeignkey
}

func (p PGSchema) Load(filename string) *PGSchema {
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(fp)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(data, &p); err != nil {
		panic(err)
	}
	return &p
}

func (p *PGSchema) Dump() string {
	buf := bytes.NewBuffer([]byte{})

	for _, item := range p.Sequences {
		buf.WriteString(item.Create())
	}
	for _, item := range p.Tables {
		buf.WriteString(item.Create())
	}
	for _, item := range p.Indexes {
		buf.WriteString(item.Create())
	}
	for _, item := range p.Foreignkeys {
		buf.WriteString(item.Create())
	}

	return buf.String()
}

func (p *PGSchema) Sql(filename string) {
	fp, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	if _, err := fp.WriteString(p.Dump()); err != nil {
		fmt.Printf("WriteString Error: %s\n", err.Error())
	}
}

func (p *PGSchema) Yaml(filename string) {
	fp, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	data, err := yaml.Marshal(&p)
	if err != nil {
		panic(err)
	}
	if _, err := fp.Write(data); err != nil {
		fmt.Printf("WriteString Error: %s\n", err.Error())
	}
}

func (p *PGSchema) Diff(src *PGSchema) string {
	buf := bytes.NewBuffer([]byte{})
	for _, item := range p.Sequences {
		buf.WriteString(item.Diff(&src.Sequences))
	}
	for _, item := range p.Tables {
		buf.WriteString(item.Diff(&src.Tables))
	}
	for _, item := range p.Indexes {
		buf.WriteString(item.Diff(&src.Indexes))
	}
	for _, item := range p.Foreignkeys {
		buf.WriteString(item.Diff(&src.Foreignkeys))
	}
	return buf.String()
}
