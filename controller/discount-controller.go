package controller

import (
	"net/http"

	"encoding/json"

	"github.com/mrnaghibi/discount/entity"
	"github.com/mrnaghibi/discount/errors"
	"github.com/mrnaghibi/discount/service"
)

type DiscountController interface {
	SaveDiscount(response http.ResponseWriter, request *http.Request)
	GetDiscount(response http.ResponseWriter, request *http.Request)
	ConsumeDiscount(response http.ResponseWriter, request *http.Request)
	ReportDiscount(response http.ResponseWriter, request *http.Request)
}
type Body struct {
	Mobile   string `json:"mobile"`
	Discount string `json:"discount"`
}

var (
	discountService service.DiscountService
)

type controller struct{}

func NewDiscountController(discountSRV service.DiscountService) DiscountController {
	discountService = discountSRV
	return &controller{}
}

func (*controller) SaveDiscount(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var discount entity.Discount
	err := json.NewDecoder(request.Body).Decode(&discount)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: err.Error()})
		return
	}
	result, err1 := discountService.Create(&discount)

	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: err1.Error()})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
func (*controller) ConsumeDiscount(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var body Body
	err := json.NewDecoder(request.Body).Decode(&body)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: "Unmarshalling Data"})
		return
	}
	result, err1 := discountService.Consume(body.Discount, body.Mobile)
	if err1 != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: err1.Error()})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)

}
func (*controller) ReportDiscount(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params, ok := request.URL.Query()["discount"]
	if !ok || len(params[0]) < 1 {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: "discount param required"})
		return
	}

	statistics, err := discountService.Report(params[0])
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: err.Error()})
		return
	}
	if len(statistics) == 0 {
		response.WriteHeader(http.StatusOK)
		response.Write([]byte("[]"))
	} else {
		response.WriteHeader(http.StatusOK)
		json.NewEncoder(response).Encode(statistics)
	}

}

func (*controller) GetDiscount(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	params, ok := request.URL.Query()["discount"]
	if !ok || len(params[0]) < 1 {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: "discount param required"})
		return
	}
	result, err := discountService.GetDiscount(params[0])
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiecError{Message: err.Error()})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
