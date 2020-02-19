package main

import router "github.com/mrnaghibi/discount/http"


var httpRouter = router.NewMuxRouter()

func handleRequest() {
	discountController := initWalletController()
	httpRouter.POST("/api/discounts/consume",discountController.ConsumeDiscount)
	httpRouter.GET("/api/discounts/statistics",discountController.ReportDiscount)
	//httpRouter.SERVE(os.Getenv("PORT"))
	httpRouter.SERVE(":8000")

}
func main() {
	handleRequest()
}
