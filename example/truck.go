//go:generate go run ../builders/main.go
package example

type Truck struct {
	id                 string `immutables:"required"`
	brand              string `immutables:"required"`
	model              string `immutables:"required"`
	registrationNumber string `immutables:"required"`
	serialNumber       string
	vin                string
}
