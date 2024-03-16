package impl

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/main-shopping-ms/service/purchase_service"
	"github.com/prakash-p-3121/ordermgtmodel"
	"github.com/prakash-p-3121/restlib"
)

type PurchaseControllerImpl struct {
	PurchaseService purchase_service.PurchaseService
}

func (controller *PurchaseControllerImpl) CreatePurchase(restCtx restlib.RestContext) {
	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerError := errorlib.NewInternalServerError("expected-gin-controller")
		internalServerError.SendRestResponse(ginRestCtx.CtxGet())
		return
	}

	ctx := ginRestCtx.CtxGet()
	var req ordermgtmodel.OrderCreateReq
	err := ctx.BindJSON(&req)
	if err != nil {
		badReqErr := errorlib.NewBadReqError("payload-serialization=" + err.Error())
		badReqErr.SendRestResponse(ctx)
		return
	}

	idGenResp, appErr := controller.PurchaseService.CreatePurchase(&req)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}

	restlib.OkResponse(ctx, idGenResp)
}
