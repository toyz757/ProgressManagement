package main

import (
	"./ui/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("ui/views/*.html")

//	user := router.Group("/user")
//	{
//		user.POST("/signup", routes.UserSignUp)
//		user.POST("/login", routes.UserLogIn)
//		user.POST("/logout", routes.UserLogOut)
//	}

	router.GET("/", routes.Home)
	router.GET("/ticket", routes.Ticket)
	router.NoRoute(routes.NoRoute)

	router.Run(":8080")
}
