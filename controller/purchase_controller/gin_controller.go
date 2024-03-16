package purchase_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prakash-p-3121/restlib"
)

func CreatePurchase(ctx *gin.Context) {
	ginRestCtx := restlib.NewGinRestContext(ctx)
	controller := NewPurchaseController()
	controller.CreatePurchase(ginRestCtx)
}
