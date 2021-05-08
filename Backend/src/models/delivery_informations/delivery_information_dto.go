package delivery_informations

type DeliveryInformationDto struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	City    string `json:"city"`
	ZipCode uint   `json:"zip_code"`
}
