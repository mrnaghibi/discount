package main

import (
	"os"
	"github.com/mrnaghibi/discount/controller"
	router "github.com/mrnaghibi/discount/http"
	"github.com/mrnaghibi/discount/repository"
	"github.com/mrnaghibi/discount/service"
)

var (
	discountRepository repository.Discount           = repository.NewSliceRepository()
	discountService    service.DiscountService       = service.NewDiscountService(discountRepository)
	discountController controller.DiscountController = controller.NewDiscountController(discountService)
	httpRouter         router.Router                 = router.NewMuxRouter()
)

func handleRequest() {
	
	httpRouter.POST("/api/discounts", discountController.SaveDiscount)
	httpRouter.GET("/api/discounts", discountController.GetDiscount)
	httpRouter.POST("/api/discounts/consume",discountController.ConsumeDiscount)
	httpRouter.GET("/api/discounts/statistics",discountController.ReportDiscount)
	httpRouter.SERVE(os.Getenv("PORT"))
}

func main() {
	handleRequest()
}
