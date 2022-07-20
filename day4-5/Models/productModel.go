package Models

type Product struct {
	ProdId     int    `json:"id"`
	ProdName   string `json:"name"`
	Price      int    `json:"price"`
	Quantity   int    `json:"quantity"`
}

func (b *Product) TableName() string {
	return "Product"
}
