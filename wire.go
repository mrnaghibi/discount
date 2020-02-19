//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/mrnaghibi/discount/controller"
	"github.com/mrnaghibi/discount/repository"
	"github.com/mrnaghibi/discount/service"
)

func initWalletController() controller.DiscountController {
	wire.Build(repository.DiscountRepositoryProvider,service.DiscountServiceProvider,controller.DiscountControllerProvider)
	return controller.DiscountController{}
}
