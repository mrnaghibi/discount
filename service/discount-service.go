package service

import (
	"github.com/mrnaghibi/discount/entity"
	"github.com/mrnaghibi/discount/repository"
)

type DiscountService struct {
	repo repository.DiscountRepository
}

func DiscountServiceProvider(repository repository.DiscountRepository) DiscountService {
	return DiscountService{repo: repository}
}

func (s *DiscountService) DecreaseDiscountNum(discountName string) (bool, error) {
	return s.repo.DecreaseDiscountNum(discountName)
}
func (s *DiscountService) ReportArvanCoupon() entity.Statistic {
	return s.repo.ReportArvanCoupon()
}
