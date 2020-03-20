package unifiedrecharge

import (
	"context"
	ap "google.golang.org/api/androidpublisher/v3"
	"google.golang.org/api/option"
	"sync"
)

//call androidpushliser v2
var (
	once            sync.Once
	purchaseService *ap.PurchasesProductsService
	lastError       error
)

func getService(configFile string) *ap.PurchasesProductsService {
	once.Do(func() {
		ctx := context.Background()
		service, err := ap.NewService(ctx, option.WithCredentialsFile(configFile))
		if err != nil {
			lastError = err
			return
		}

		purchaseService = ap.NewPurchasesProductsService(service)
	})

	return purchaseService
}

func CheckGoogleAndroidPurchase(configFile, packageName, productId, token string) (productPurchase *ap.ProductPurchase, err error) {
	service := getService(configFile)
	if service == nil {
		return nil, lastError
	}

	return service.Get(packageName, productId, token).Do()
}
