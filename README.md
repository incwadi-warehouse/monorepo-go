# monorepo-go

incwadi is a book database to manage your books.

## Getting Started

[Installing](https://github.com/incwadi-warehouse/docu)

## conf

```go
import "github.com/incwadi-warehouse/monorepo-go/conf/settings"

// Load config
data, err := settings.LoadFromString(schema, file)
if err != nil {
    log.Fatal(err)
}

// Manipulation
v := data.Get("app.key")
data.Add("app.key2", '1')
data.Rm("app.key2")
```

## settings

Create a `.env` file to define some settings.

```env
// .env

ENV=prod
FILE_PATH=./
CORS_ALLOW_ORIGIN=http://localhost:8080
```

- ENV - Set to `prod`, `dev` or `test`
- FILE_PATH - Path where to store the files
- CORS_ALLOW_ORIGIN - Allowed origins
