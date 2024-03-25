package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AccountRegistrationHandler(ctx *gin.Context) {

	// get session
	session := sessions.Default(ctx)

	// initialise template data
	data := TemplateData{
		Title:   "Account Registration",
		Profile: session.Get("profile"),
	}

	// render template and pass in data
	ctx.HTML(http.StatusOK, "sign-up", gin.H{
		"data": data,
	})
}
