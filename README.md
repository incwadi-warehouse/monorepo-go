# monorepo-go

incwadi is a book database to manage your books.

## Getting Started

[Installing](https://github.com/incwadi-warehouse/docu)

## Example for module conf

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
go run . ./settings/example.schema.json ./settings/example.json add app.key2 value
go run . ./settings/example.schema.json ./settings/example.json get app.key2
go run . ./settings/example.schema.json ./settings/example.json rem app.key2
```
