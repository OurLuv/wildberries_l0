package model

type Payment struct {
	Id           int
	Transaction  string  `json:"transaction"`
	RequestID    *string `json:"request_id"`
	Currency     string  `json:"currency"`
	Provider     string  `json:"provider"`
	Amount       int     `json:"amount"`
	PaymentDate  int64   `json:"payment_dt"`
	Bank         string  `json:"bank"`
	DeliveryCost int     `json:"delivery_cost"`
	GoodsTotal   int     `json:"goods_total"`
	CustomFee    int     `json:"custom_fee"`
}
