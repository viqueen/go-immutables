//go:generate go run ../cmd/builders.go
package assets

type Truck struct {
	Id                 string
	Brand              string
	Model              string
	RegistrationNumber string
	SerialNumber       string
	Vin                string
}
