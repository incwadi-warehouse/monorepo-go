package main

import (
	"log"

	"github.com/incwadi-warehouse/monorepo-go/blog/router"
	"github.com/incwadi-warehouse/monorepo-go/framework/config"
	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("CORS_ALLOW_ORIGIN", "http://127.0.0.1")

    config.LoadAppConfig()

	r := router.Routes()
	log.Fatal(r.Run(":8080"))
}
