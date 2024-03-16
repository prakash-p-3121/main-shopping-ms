package impl

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/main-shopping-ms/service/signup_service"
	"github.com/prakash-p-3121/restlib"
	"github.com/prakash-p-3121/usermodel"
	"log"
)

type SignupControllerImpl struct {
	Service signup_service.SignupService
}

func (controller *SignupControllerImpl) Signup(restCtx restlib.RestContext) {
	ginRestCtx, ok := restCtx.(*restlib.GinRestContext)
	if !ok {
		internalServerError := errorlib.NewInternalServerError("expected-gin-controller")
		internalServerError.SendRestResponse(ginRestCtx.CtxGet())
		return
	}

	ctx := ginRestCtx.CtxGet()
	var req usermodel.UserCreateReq
	err := ctx.BindJSON(&req)
	if err != nil {
		badReqErr := errorlib.NewInternalServerError("payload-serialization=" + err.Error())
		badReqErr.SendRestResponse(ctx)
		return
	}

	idGenResp, appErr := controller.Service.Signup(&req)
	if appErr != nil {
		appErr.SendRestResponse(ctx)
		return
	}
	log.Println("idGenResp=", idGenResp)
	restlib.OkResponse(ctx, idGenResp)
}
