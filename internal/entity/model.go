package entity

type AddressField struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func NewAddressField(name, address, phone string) (field *AddressField) {
	field = new(AddressField)

	field.Name = name
	field.Address = address
	field.Phone = phone

	return
}
