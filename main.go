package main

import (
	"github.com/gin-gonic/gin"
	"github.com/skumaran141989/game-api/controller"
	"github.com/skumaran141989/game-api/utilities"
)

func main() {
	router := gin.Default()
	var creds *utilities.DBCreds
	controller.Analytics(router, creds)
	router.Run("localhost:8080")
}
