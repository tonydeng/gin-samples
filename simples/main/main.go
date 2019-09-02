package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//engine := gin.Default()
	//engine.Any("/", WebRoot)
	//
	//engine.Run(":9205")
	router()
}

func router() {

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World!")
	})

	router.GET("/user/:name", func(context *gin.Context) {
		name := context.Param("name")
		context.String(http.StatusOK, "Hello %s!", name)
	})

	router.GET("/user/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")

		message := name + " is " + action

		context.String(http.StatusOK, message)
	})

	router.Run(":9205")
}
