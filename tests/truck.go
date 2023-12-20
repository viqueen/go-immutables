//go:generate go run ../cmd/builders.go
package tests

type Truck struct {
	Id                 string `immutables:"required"`
	Brand              string `immutables:"required"`
	Model              string `immutables:"required"`
	RegistrationNumber string `immutables:"required"`
	SerialNumber       string
	vin                string
}
