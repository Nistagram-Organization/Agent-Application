package products

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	OnStock     uint    `json:"on_stock"`
	Image       string  `json:"image"`
}
