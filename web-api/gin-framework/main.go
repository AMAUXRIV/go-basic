package main

import (
	
	"github.com/gin-gonic/gin"
)

func main () {
	app := gin.Default()

	route := app
	route.GET("/", func(ctx *gin.Context) {
		isValidated := false

		if (!isValidated) {

			ctx.AbortWithStatusJSON(400, gin.H{
				"message": "bad request coeg",
			} )
			
		}
		ctx.JSON(200, gin.H{
			"hello": "world",
		} )
		
	})

	app.Run(":8000")
}