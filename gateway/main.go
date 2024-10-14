package main

import (
	"log"

	"github.com/incwadi-warehouse/monorepo-go/framework/config"
	"github.com/incwadi-warehouse/monorepo-go/framework/cors"
	"github.com/incwadi-warehouse/monorepo-go/gateway/router"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func main() {
    godotenv.Load()

    config.LoadAppConfig()

    viper.SetDefault("CORS_ALLOW_ORIGIN", "*")

    corsConfig := cors.NewCors()
    corsConfig.AllowOrigins = []string{viper.GetString("CORS_ALLOW_ORIGIN")}
    corsConfig.SetCorsHeaders()

    r := router.Routes()
	r.Use(corsConfig.SetCorsHeaders())

    log.Fatal(r.Run(":8080"))
}
