# monorepo-go

incwadi is a book database to manage your books.

## Getting Started

[Installing](https://github.com/incwadi-warehouse/docu)

## Limitations

- The config goes only two levels deep.

## Example

Run in conf dir.

```shell
go run . ./settings/example.schema.json ./settings/example.json add app.key2 value
go run . ./settings/example.schema.json ./settings/example.json get app.key2
go run . ./settings/example.schema.json ./settings/example.json rem app.key2
```
