package purchase_controller

import (
	"github.com/prakash-p-3121/main-shopping-ms/controller/purchase_controller/impl"
	"github.com/prakash-p-3121/main-shopping-ms/service/purchase_service"
)

func NewPurchaseController() PurchaseController {
	return &impl.PurchaseControllerImpl{PurchaseService: purchase_service.NewPurchaseService()}
}
