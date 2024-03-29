package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AccountRecoveryHandler(ctx *gin.Context) {

	// get session
	session := sessions.Default(ctx)

	// initialise template data
	data := TemplateData{
		Title:   "Account Recovery",
		Profile: session.Get("profile"),
	}

	// render template and pass in data
	ctx.HTML(http.StatusOK, "recovery", gin.H{
		"data": data,
	})
}
