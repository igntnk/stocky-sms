package models

type Product struct {
	Uuid        string  `bson:"uuid" json:"uuid" mapstructure:"uuid"`
	ProductCode string  `bson:"product_code" json:"product_code" mapstructure:"product_code"`
	StoreCost   float64 `bson:"store_cost" json:"store_cost" mapstructure:"store_cost"`
	StoreAmount float64 `bson:"store_amount" json:"store_amount" mapstructure:"store_amount"`
}

type RemoveProductRequest struct {
	Uuid   string  `bson:"uuid" json:"uuid" mapstructure:"uuid"`
	Amount float64 `bson:"amount" json:"amount" mapstructure:"amount"`
}
