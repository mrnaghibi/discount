package repository

import (
	"github.com/mrnaghibi/discount/entity"
)
type Discount interface{
	Save(*entity.Discount) (*entity.Discount,error)
	GetDiscount(string) (*entity.Discount,error)
	Consume(string,string) (*entity.Discount,error)
	Report(string) ([]entity.Statistic,error)

}