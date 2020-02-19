package main

import (
	router "github.com/mrnaghibi/discount/http"
	"os"
)


var httpRouter = router.NewMuxRouter()

func handleRequest() {
	discountController := initWalletController()
	httpRouter.POST("/api/discounts/consume",discountController.ConsumeDiscount)
	httpRouter.GET("/api/discounts/statistics",discountController.ReportDiscount)
	httpRouter.HTML("/html/")
	httpRouter.SERVE(os.Getenv("PORT"))
}
func main() {
	handleRequest()
}
