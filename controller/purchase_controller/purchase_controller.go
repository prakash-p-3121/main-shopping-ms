package purchase_controller

import (
	"github.com/prakash-p-3121/restlib"
)

type PurchaseController interface {
	CreatePurchase(restCtx restlib.RestContext)
}
