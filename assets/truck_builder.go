
package assets

type TruckBuilder struct {
	target Truck
}

func NewTruckBuilder() *TruckBuilder {
	return &TruckBuilder{ Truck{} }
}

func (b *TruckBuilder) WithId(Id string) *TruckBuilder {
	b.target.Id = Id
	return b
}

func (b *TruckBuilder) WithBrand(Brand string) *TruckBuilder {
	b.target.Brand = Brand
	return b
}

func (b *TruckBuilder) WithModel(Model string) *TruckBuilder {
	b.target.Model = Model
	return b
}

func (b *TruckBuilder) WithRegistrationNumber(RegistrationNumber string) *TruckBuilder {
	b.target.RegistrationNumber = RegistrationNumber
	return b
}

func (b *TruckBuilder) WithSerialNumber(SerialNumber string) *TruckBuilder {
	b.target.SerialNumber = SerialNumber
	return b
}

func (b *TruckBuilder) WithVin(Vin string) *TruckBuilder {
	b.target.Vin = Vin
	return b
}

func (b *TruckBuilder) Build() Truck {
	return b.target
}
