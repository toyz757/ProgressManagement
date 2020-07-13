package main

import (
	"./ui/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("ui/views/*.html")

	router.GET("/", routes.Home)
	router.GET("/ticket", routes.Ticket)
	router.GET("/new_ticket", routes.NewTicket)
	router.POST("/ticket", routes.PostTicket)
	router.POST("/new_ticket", routes.PostNewTicket)
	router.NoRoute(routes.NoRoute)

	router.Run(":8080")
}
