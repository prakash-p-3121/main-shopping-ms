package purchase_service

import "github.com/prakash-p-3121/main-shopping-ms/service/purchase_service/impl"

func NewPurchaseService() PurchaseService {
	return &impl.PurchaseServiceImpl{}
}
