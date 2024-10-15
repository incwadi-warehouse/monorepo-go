# monorepo-go

monorepo-go is a database to manage your warehouse.

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

The routes needs the API-Key to contain the `articles` permission.

## Framework

```go
package main

import (
 "log"

 "github.com/incwadi-warehouse/monorepo-go/blog/router"
 "github.com/incwadi-warehouse/monorepo-go/framework/config"
)

func main() {
    config.LoadAppConfig()

    r := router.Routes()
    log.Fatal(r.Run(":8080"))
}
```

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

viper.SetDefault("CORS_ALLOW_ORIGIN", "*")

corsConfig := cors.NewCors()
corsConfig.AllowOrigins = []string{viper.GetString("CORS_ALLOW_ORIGIN"), "http://127.0.0.1"}
corsConfig.SetCorsHeaders()

r := gin.Default()
r.Use(corsConfig.SetCorsHeaders())
```

## Settings

|Var                    |Description                                |Used by
|-----------------------|-------------------------------------------|--------------------------------
|ENV                    |Set to `prod`, `dev` or `test`             |conf-api
|DATA_DIR               |Path where to store the files              |conf-api
|CORS_ALLOW_ORIGIN      |Allowed origins                            |conf-api, gateway, blog
|BASE_PATH              |Configure the base path                    |conf-api
|AUTH_API_ME            |API endpoint for the user object           |conf-api
|API_CORE               |API endpoint for the core                  |gateway
|project_dir            |Path to docker compose                     |admincli
|database               |Database name to dump                      |admincli

admincli will read a config file from following paths:

- /etc/admincli/admincli.yaml

Example

```yaml
// admincli.yaml
project_dir: .
database: db-1
```
