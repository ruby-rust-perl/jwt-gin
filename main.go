package main

import (
	"jwt-gin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)

	r.Run(":8080")

}
