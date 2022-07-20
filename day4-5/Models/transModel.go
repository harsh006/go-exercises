package Models

import "time"

type Transaction struct {
	TransactionId   int    `json:"id"`
	CustomerName    string `json:"name"`
	ProductId       int    `json:"prod_id"`
	ProductQuantity int    `json:"quantity_purchased"`
	//AmountPaid      int    `json:"amount"`
	Status string `json:"status"`
	TransTime time.Time `json:"trans_time"`
}
func (b *Transaction) TableName() string {
	return "Transaction"
}

