package strategy

import (
	"testing"
)

func TestTravel(t *testing.T) {
	var (
		car      Car
		air      Airplane
		traveler Traveler
	)
	traveler.SetTraveler(car)
	traveler.Travel()

	traveler.SetTraveler(air)
	traveler.Travel()
}

func TestNewPayment(t *testing.T) {
	payment := NewPayment("tlf", "123456", 10, &Bank{})
	payment.Pay()

	payment = NewPayment("pxy", "666666", 100, &CreditCard{})
	payment.Pay()
}
