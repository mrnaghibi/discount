package entity

type Statistic struct{
	ID int64 `json:"id"`
	Mobile string `json:"mobile"`
	Discount Discount `json:"discount"`
}