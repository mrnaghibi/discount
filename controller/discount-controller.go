package controller

import (
	"encoding/json"
	"github.com/mrnaghibi/discount/entity"
	router "github.com/mrnaghibi/discount/http"
	"net/http"

	"github.com/mrnaghibi/discount/errors"
	"github.com/mrnaghibi/discount/service"
)



type DiscountController struct {
	service service.DiscountService
}

func DiscountControllerProvider(DiscountService service.DiscountService) DiscountController {
	return DiscountController{service: DiscountService}
}

var (
	walletCharger  router.SendHttpRequest
	discountCH  = make(chan string, 1000)
	duplicateCH = make(chan struct{})
	doneCH      = make(chan struct{})
)

func (discountController *DiscountController) ConsumeDiscount(response http.ResponseWriter, request *http.Request) {
	var requestModel entity.DiscountRequestModel
	err := json.NewDecoder(request.Body).Decode(&requestModel)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: "Bad Request!"})
		return
	}
	go makeDecrease(requestModel, discountController)
	select {
		case mobile := <-discountCH:

			err := walletCharger.Send(mobile)
			if err != nil {
				response.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(response).Encode(errors.ServiecError{Message: "something went wrong!"})
			} else {
				response.WriteHeader(http.StatusOK)
				json.NewEncoder(response).Encode(errors.ServiecError{Message: "Wallet Charged Successfully"})
				//Send Statistics to Pusher
				statistic := discountController.service.ReportArvanCoupon()

			}

		case _ = <-doneCH:
			response.WriteHeader(http.StatusForbidden)
			json.NewEncoder(response).Encode(errors.ServiecError{Message: "Coupon not allowed!"})
		case _ = <-duplicateCH:
			response.WriteHeader(http.StatusForbidden)
			json.NewEncoder(response).Encode(errors.ServiecError{Message: "Coupon Used Before!"})
	}
}

func (discountController *DiscountController) ReportDiscount(response http.ResponseWriter, request *http.Request) {
	statistic := discountController.service.ReportArvanCoupon()
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(statistic)
}

func makeDecrease(model entity.DiscountRequestModel, discountController *DiscountController) {
	result, err := discountController.service.DecreaseDiscountNum(model.Mobile)
	if err == nil {
		if result {
			discountCH <- model.Mobile
		} else {
			doneCH <- struct{}{}
		}
	} else {
		duplicateCH <- struct{}{}
	}
}
