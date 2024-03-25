package handlers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ProfileSettingsHandler(ctx *gin.Context) {

	// get session
	session := sessions.Default(ctx)

	// initialise template data
	data := TemplateData{
		Title:   "Profile Settings",
		Profile: session.Get("profile"),
	}

	// render template and pass in data
	ctx.HTML(http.StatusOK, "profile-settings", gin.H{
		"data": data,
	})

}
