package signup_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/restlib"
)

func Signup(ctx *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(ctx)
	controller := NewSignupController()
	controller.Signup(ginRestCtx)
}
