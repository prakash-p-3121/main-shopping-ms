package purchase_service

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
	"github.com/prakash-p-3121/ordermgtmodel"
)

type PurchaseService interface {
	CreatePurchase(req *ordermgtmodel.OrderCreateReq) (*idgenmodel.IDGenResp, errorlib.AppError)
}
