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
