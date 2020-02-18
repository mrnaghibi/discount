package service

import (
	"github.com/mrnaghibi/discount/entity"
	"github.com/mrnaghibi/discount/repository"
)

type DiscountService interface {
	Create(*entity.Discount) (*entity.Discount, error)
	GetDiscount(string) (*entity.Discount, error)
	Consume(string, string) (*entity.Discount,error)
	Report(string) ([]entity.Statistic,error)
}

var (
	repo repository.Discount
)

type service struct{}

func NewDiscountService(repository repository.Discount) DiscountService {
	repo = repository
	return &service{}
}

func (*service) Create(discount *entity.Discount) (*entity.Discount, error) {
	return repo.Save(discount)
}
func (*service) GetDiscount(discount string) (*entity.Discount,error) {
	return repo.GetDiscount(discount)
}
func (*service) Consume(discountName string, mobileNumber string) (*entity.Discount,error) {
	return repo.Consume(discountName,mobileNumber)
}
func (*service) Report(discountName string) ([]entity.Statistic,error) {
	return repo.Report(discountName)
}
