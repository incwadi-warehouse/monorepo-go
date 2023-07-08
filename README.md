# monorepo-go

incwadi is a book database to manage your books.

## Getting Started

[Installing](https://github.com/incwadi-warehouse/docu)

## conf

```go
import "github.com/incwadi-warehouse/monorepo-go/conf/manager"

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
PROJECT_DIR=
```

|Var                    |Used by                |Description
|-----------------------|-----------------------|--------------------------------
|ENV                    |conf-api, search-api   |Set to `prod`, `dev` or `test`
|DATA_DIR               |conf-api               |Path where to store the files
|CORS_ALLOW_ORIGIN      |conf-api, search-api   |Allowed origins
|BASE_PATH              |conf-api, search-api   |Configure the base path
|AUTH_API_ME            |conf-api, search-api   |API endpoint for the user object
|ADMINCLI_PROJECT_DIR   |admincli               |Where to execute the commands
|MEILI                  |search-api             |URL to meilisearch instance
|MEILI_TOKEN            |search-api             |API key for meilisearch
|BRANCHES               |search-api             |List of allowed branches, comma-separated
|INDEXES                |search-api             |List of allowed indexes, comma-separated

admincli will read a config file from following paths:

- /etc/admincli/admincli.yaml
- $HOME/.admincli/admincli.yaml
- ./admincli.yaml

Example

```yaml
// admincli.yaml
project_dir: .
```

## Testing

```shell
curl -X GET http://localhost:8080/api/user/1/snow -H "Content-Type: application/json" -H "Authorization: Bearer token"
```

```shell
curl -X POST http://localhost:8080/api/user/1/snow2 -H "Content-Type: application/json" -H "Authorization: Bearer token" -d '{"value": 1 }'
```

```shell
curl -X DELETE http://localhost:8080/api/user/1/snow2 -H "Content-Type: application/json" -H "Authorization: Bearer token"
```
