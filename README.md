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
DATA_DIR=./data/
CORS_ALLOW_ORIGIN=http://localhost:80
API_KEY=
```

- ENV - Set to `prod`, `dev` or `test`
- DATA_DIR - Path where to store the files
- CORS_ALLOW_ORIGIN - Allowed origins
- API_KEY - A secure token to authenticate access
