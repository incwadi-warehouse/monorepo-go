# monorepo-go

monorepo-go is a database to manage your warhouse.

## Getting Started

## Conf

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

## Blog

Mount auth volume to `/usr/src/app/data/auth/` and data volume to `/usr/src/app/data/content/`.

## Framework

### Config

```go
import "github.com/incwadi-warehouse/monorepo-go/framework/config"

config.LoadAppConfig(config.WithName("myconfig"), config.WithFormat("json"), config.WithPaths("./config", "."))

viper.SetDefault("CORS_ALLOW_ORIGIN", "http://127.0.0.1")
```

### ApiKey

```go
import "github.com/incwadi-warehouse/monorepo-go/framework/apikey"

apikey.IsValidAPIKey("key")
apikey.HasPermission("key", "permission")
```

### Cors

```go
import "github.com/incwadi-warehouse/monorepo-go/framework/cors"

cors.Headers("*")
cors.Headers("http://127.0.0.1")
```

## Settings

|Var                    |Used by                        |Description
|-----------------------|-------------------------------|--------------------------------
|ENV                    |conf-api, search-api           |Set to `prod`, `dev` or `test`
|DATA_DIR               |conf-api                       |Path where to store the files
|CORS_ALLOW_ORIGIN      |conf-api, search-api, gateway  |Allowed origins
|BASE_PATH              |conf-api, search-api           |Configure the base path
|AUTH_API_ME            |conf-api, search-api           |API endpoint for the user object
|ADMINCLI_PROJECT_DIR   |admincli                       |Where to execute the commands
|MEILI                  |search-api                     |URL to meilisearch instance
|MEILI_TOKEN            |search-api                     |API key for meilisearch
|BRANCHES               |search-api                     |List of allowed branches, comma-separated
|INDEXES                |search-api                     |List of allowed indexes, comma-separated
|API_CORE               |gateway                        |API endpoint for the core

admincli will read a config file from following paths:

- /etc/admincli/admincli.yaml
- $HOME/.admincli/admincli.yaml
- ./admincli.yaml

Example

```yaml
// admincli.yaml
project_dir: .
```
