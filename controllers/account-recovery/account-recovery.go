package accountrecovery

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type TemplateData struct {
	Title   string
	Profile interface{}
}

func Handler(ctx *gin.Context) {

	// get session
	session := sessions.Default(ctx)

	// initialise template data
	data := TemplateData{
		Title:   "Error",
		Profile: session.Get("profile"),
	}

	// render template and pass in data
	ctx.HTML(http.StatusOK, "account-recovery", gin.H{
		"data": data,
	})
}
