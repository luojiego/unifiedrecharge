package unifiedrecharge

import (
	"context"
	"fmt"
	ap "google.golang.org/api/androidpublisher/v2"
	"google.golang.org/api/option"
	"sync"
)

//直接调用androidpushliser v2
var (
	once sync.Once
	purchaseService *ap.PurchasesProductsService
)

func getService(configFile string) *ap.PurchasesProductsService{
	once.Do(func() {
		ctx := context.Background()
		service, err := ap.NewService(ctx, option.WithCredentialsFile(configFile))
		if err != nil {
			fmt.Println("new service err: ", err)
		}

		purchaseService = ap.NewPurchasesProductsService(service)
	})

	return purchaseService
}

func CheckGoogleAndroidPurchase(configFile, packageName, productId, token string) (productPurchase *ap.ProductPurchase, err error) {
	return getService(configFile).Get(packageName, productId, token).Do()
}