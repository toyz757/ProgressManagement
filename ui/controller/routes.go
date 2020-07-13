package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context){
	ctx.HTML(http.StatusOK, "home.html", gin.H{})
}

func Ticket(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ticket.html", gin.H{})
}

func NoRoute(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
}
