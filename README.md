# monorepo-go

incwadi is a book database to manage your books.

## Getting Started

[Installing](https://github.com/incwadi-warehouse/docu)

## conf

Usage: conf [action]

Actions

get [key] [value] [schema-url] [file-url] - Get the value of an entry
add [key] [schema-url] [file-url] - Add or update entry
rm [key] [schema-url] [file-url] - Remove an entry

```go
import "github.com/incwadi-warehouse/monorepo-go/conf/settings"

// LoadFromUrl() OR LoadFromString()
data, err : = settings.LoadFromUrl("./example.schema.json", "./example.json")
if err != nil {
    log.Fatal(err)
}
data, err : = settings.LoadFromString(schema, file)
if err != nil {
    log.Fatal(err)
}

// Manipulation
v := data.Get("app.key")
fmt.PrintLn(v)
data.Add("app.key2", '1')
data.Rm("app.key2")

// Works only in conjunction with LoadFromUrl()
data.Write()
```

Run in conf dir.

```shell
go run . add app.key2 value ./settings/example.schema.json ./settings/example.json
go run . get app.key2 ./settings/example.schema.json ./settings/example.json
go run . rm app.key2 ./settings/example.schema.json ./settings/example.json
```
