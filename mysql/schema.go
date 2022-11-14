package mysql

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type ISchema struct {
	Tables []ITable
}

func (p ISchema) Load(filename string) *ISchema {
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

func (p *ISchema) Dump() string {
	buf := bytes.NewBuffer([]byte{})

	for _, item := range p.Tables {
		buf.WriteString(item.Create())
	}

	return buf.String()
}

func (p *ISchema) Sql(filename string) {
	fp, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	if _, err := fp.WriteString(p.Dump()); err != nil {
		fmt.Printf("WriteString Error: %s\n", err.Error())
	}
}

func (p *ISchema) Yaml(filename string) {
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

func (p *ISchema) Diff(src *ISchema) string {
	buf := bytes.NewBuffer([]byte{})
	for _, item := range p.Tables {
		buf.WriteString(item.Diff(&src.Tables))
	}
	return buf.String()
}
