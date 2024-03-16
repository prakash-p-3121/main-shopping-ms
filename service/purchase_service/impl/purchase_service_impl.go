package impl

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/main-shopping-ms/cfg"
	"github.com/prakash-p-3121/ordermgtclient"
	"github.com/prakash-p-3121/ordermgtmodel"
)

type PurchaseServiceImpl struct {
}

func (service *PurchaseServiceImpl) CreatePurchase(req *ordermgtmodel.OrderCreateReq) (*idgenmodel.IDGenResp, errorlib.AppError) {
	if req == nil {
		return nil, errorlib.NewBadReqError("req-nil")
	}

	userMgtMsCfg, err := cfg.GetMsConnectionCfg("ordermgtms")
	if err != nil {
		return nil, errorlib.NewInternalServerError(err.Error())
	}

	orderMgtClient := ordermgtclient.NewOrderMgtClient(userMgtMsCfg.Host, uint(userMgtMsCfg.Port))
	idResp, appErr := orderMgtClient.OrderCreate(req)
	if appErr != nil {
		return nil, appErr
	}

	return idResp, nil
}
