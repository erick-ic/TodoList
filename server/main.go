package main

import (
	"server/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers.Routers(router)
	router.Run()
}
