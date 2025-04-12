package main

import (
	"giron/config"
	"giron/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    // setup
    config.LoadEnv()
    config.InitMongo(config.GetEnv("MONGO_URI", "mongodb://localhost:27017"), config.GetEnv("MONGO_DB", "giron"))

    // setup gin routes
    r := gin.Default()
    routes.SetupRoutes(r)

    // run gin
    r.Run(":" + config.GetEnv("PORT", "9000"))
}
