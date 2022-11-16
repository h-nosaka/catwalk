# Catwalk

Catwalk is a tool to manage DB schema.  
supported mysql  
later postgresql

### Instration

```shell
go install github.com/h-nosaka/catwalk@latest
```

```shell
go get github.com/h-nosaka/catwalk
```

### Usage

```shell
catwalk diff -c ./schema/schema.yaml
```

```shell
catwalk run -c ./shcema/shcema.yaml
```

```shell
catwalk help
```

### Usage on mage

```shell
# exsample magefile
mage -compile ./catwalk
./catwalk
```

```shell
./catwalk createDatabase
./catwalk migrate
```

### mysql examples

- [schema](./examples/schema.go)
- [yaml](./examples/schema.yaml)
- [test](./examples/schema_test.go)
