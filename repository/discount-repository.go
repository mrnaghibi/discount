package repository

import (
	"errors"
	"github.com/mrnaghibi/discount/entity"
)

type DiscountRepository struct{}

var (
	arvanCoupon = make(map[string]string)
	discount    = entity.Discount{
		Code: "arvan",
		Num:  1000,
	}
)

func DiscountRepositoryProvider() DiscountRepository {
	return DiscountRepository{}
}

func (r *DiscountRepository) DecreaseDiscountNum(mobile string) (bool, error) {
	if _, ok := arvanCoupon[mobile]; ok {
		return false, errors.New("used before")
	}
	if discount.Num > 0 {
		discount.Num -= 1
		arvanCoupon[mobile] = mobile
		return true, nil
	}
	return false, nil
}

func (*DiscountRepository) ReportArvanCoupon() entity.Statistic {

	var st entity.Statistic
	for _, statistic := range arvanCoupon {
		st.Mobile = append(st.Mobile, statistic)
	}
	st.Count = len(arvanCoupon)
	return st
}
