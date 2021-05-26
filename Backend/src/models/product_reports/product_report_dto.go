package product_reports

type ProductReport struct {
	Name string  `json:"product_name"`
	Sold        uint    `json:"sold"`
	Income      float32 `json:"income"`
}
