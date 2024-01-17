# Catwalk

Catwalk is a tool to manage DB schema.  
supported mysql  
supported postgresql  
later elasticsearch

### Instration

```shell
go get github.com/h-nosaka/catwalk
```

### Usage

```go
catwalk.NewSchema(
  "app", catwalk.SchemaModeMySQL,
  catwalk.NewTable("app", "accounts", catwalk.JsonCaseSnake, "アカウントマスタ").SetDefaultColumns(
		catwalk.DataTypeUUID,
		catwalk.NewColumn("email", catwalk.DataTypeString, 256, false, "メールアドレス").Done(),
	).SetDefaultIndexes().Done(),
)
```

### examples

- [schema](./examples/examples.go)
- [test](./examples/examples_test.go)
