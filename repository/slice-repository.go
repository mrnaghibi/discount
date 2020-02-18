package repository

import (
	"errors"
	"math/rand"

	"github.com/mrnaghibi/discount/entity"
)

var (
	discounts  []entity.Discount
	statistics []entity.Statistic
)

type sliceRepository struct{}

func NewSliceRepository() Discount {
	return &sliceRepository{}
}

func (*sliceRepository) Save(dsc *entity.Discount) (*entity.Discount, error) {
	for _, discount := range discounts {
		if discount.Name == dsc.Name {
			return nil, errors.New("Discount Code Is Exist")
		}
	}
	dsc.ID = rand.Int63()
	discounts = append(discounts, *dsc)
	return dsc, nil
}
func (*sliceRepository) GetDiscount(discountName string) (*entity.Discount, error) {
	for _, discount := range discounts {
		if discount.Name == discountName {
			return &discount, nil
		}
	}
	return nil, errors.New("Discount Doesn't Exist")
}
func (*sliceRepository) Consume(discountName string, mobileNumber string) (*entity.Discount, error) {
	for index, discount := range discounts {
		if discount.Name == discountName {
			if discount.Number == 0 {
				return nil, errors.New("Discount Is InCorrect")
			}

			for _,statistic := range statistics {
				if statistic.Discount.Name == discountName && statistic.Mobile == mobileNumber {
					return nil, errors.New("Discount Used Before")
				}
			}

			statistic := entity.Statistic{
				ID:       rand.Int63(),
				Discount: discount,
				Mobile:   mobileNumber,
			}
			statistics = append(statistics, statistic)
			discounts[index].Number -= 1
			return &discounts[index], nil
		}
	}
	return nil, errors.New("Discount Is InCorrect")
}
func (*sliceRepository) Report(discountName string) ([]entity.Statistic, error) {

	var selectedStatistics []entity.Statistic
	for _,statistic := range statistics {
		if statistic.Discount.Name == discountName {
			selectedStatistics = append(selectedStatistics , statistic)
		}
	}
	return selectedStatistics, nil
}
