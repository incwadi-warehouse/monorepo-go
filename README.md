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
CORS_ALLOW_ORIGIN=http://localhost:8080
BASE_PATH=/
AUTH_API_ME=http://localhost:8000/api/me
```

- ENV - Set to `prod`, `dev` or `test`
- DATA_DIR - Path where to store the files
- CORS_ALLOW_ORIGIN - Allowed origins
- BASE_PATH - Configure the base path
- AUTH_API_ME - APT endpoint for the user object

## Testing

```shell
curl -X GET http://localhost:8080/api/branch/1/app.key -H "Content-Type: application/json" -H "Authorization: Bearer token"
```

```shell
curl -X POST http://localhost:8080/api/branch/1/app.key2 -H "Content-Type: application/json" -H "Authorization: Bearer token" -d '{"value": 1 }'
```

```shell
curl -X DELETE http://localhost:8080/api/branch/1/app.key2 -H "Content-Type: application/json" -H "Authorization: Bearer token"
```
