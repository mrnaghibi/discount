// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/mrnaghibi/discount/controller"
	"github.com/mrnaghibi/discount/repository"
	"github.com/mrnaghibi/discount/service"
)

// Injectors from wire.go:

func initWalletController() controller.DiscountController {
	discountRepository := repository.DiscountRepositoryProvider()
	discountService := service.DiscountServiceProvider(discountRepository)
	discountController := controller.DiscountControllerProvider(discountService)
	return discountController
}
