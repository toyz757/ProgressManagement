package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func PostTicket(ctx *gin.Context) {
	println("post ticket")
	title := ctx.PostForm("title")
	explanation := ctx.PostForm("explanation")
	responsible := ctx.PostForm("responsible")
	deadline := ctx.PostForm("deadline")
	println("tilet : "+title)
	println("explanation : "+explanation)
	println("responsible : "+responsible)
	println("deadline : "+deadline)
	ctx.Redirect(http.StatusSeeOther, "/ticket")
}

func PostNewTicket(ctx *gin.Context){
	println("post new ticket")
	ctx.Redirect(http.StatusSeeOther, "/new_ticket")
}
