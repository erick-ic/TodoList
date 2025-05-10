package main

import (
	"server/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("../client/*")
	routers.Routers(router)
	router.Run()
}
