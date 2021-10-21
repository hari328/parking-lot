package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShould_CreateParkingLot_OfSizeN(t *testing.T) {
	expectedType := "service.ParkingLot"

	parkingLot := NewParkingLot(4)
	got := fmt.Sprintf("%T", parkingLot)

	if got != expectedType {
		t.Errorf("expected type %q but got %q", expectedType, got)
	}
}

func TestShouldCreateParkingLot_And_CheckInCar(t *testing.T) {
	parkingLot := NewParkingLot(4)
	got, err := parkingLot.CheckIn("KA-04-HN-2001", "white")

	expected := 1

	if got != uint(expected) {
		t.Errorf("expected first check in to have ticket number %q but got %q", expected, got)
	}

	assert.Equal(t, err, nil)
}

func TestShouldNot_CheckInIfParkingIsFull(t *testing.T) {
	parkingLot := NewParkingLot(1)

	ticket1, err := parkingLot.CheckIn("KA-04-HN-2001", "white")
	expectedTicket1 := 1
	if ticket1 != uint(expectedTicket1) {
		t.Errorf("expected first check in to have ticket number %q but got ticket1 %q", expectedTicket1, ticket1)
	}

	assert.Equal(t, nil, err)

	_, err2 := parkingLot.CheckIn("KA-04-HN-2021", "white")
	assert.EqualError(t, err2, "parking is full")
}

func TestShould_CreateParkingLotAnd_CheckInCars(t *testing.T) {
	parkingLot := NewParkingLot(4)

	ticket1, err := parkingLot.CheckIn("KA-04-HN-2001", "white")
	expectedTicket1 := 1
	if ticket1 != uint(expectedTicket1) && err == nil {
		t.Errorf("expected first check in to have ticket number %q but got ticket1 %q", expectedTicket1, ticket1)
	}

	ticket2, err := parkingLot.CheckIn("KA-04-HN-2021", "white")
	expectedTicket2 := 2
	if ticket2 != uint(expectedTicket2) && err == nil {
		t.Errorf("expected first check in to have ticket number %q but got ticket2 %q", expectedTicket2, ticket2)
	}
}

func TestShould_CheckOutCarFor_GivenTicket(t *testing.T) {
	parkingLot := NewParkingLot(4)

	regNumber := "KA-04-HN-2001"
	colour := "Green"
	ticket1, _ := parkingLot.CheckIn(regNumber, colour)

	car := parkingLot.CheckOut(ticket1)
	assert.NotNil(t, car)

	assert.Equal(t, regNumber, car.RegNumber)
	assert.Equal(t, colour, car.Colour)
}

func TestShouldNot_Checkout_WhenGivenTicketIsInvalid(t *testing.T) {
	parkingLot := NewParkingLot(4)
	dummyTicket := uint(22)
	car := parkingLot.CheckOut(dummyTicket)
	assert.Nil(t, car)
}

func TestShouldCheckInCar_PostCheckOut_AfterFullParking(t *testing.T) {
	parkingLot := NewParkingLot(1)
	regNumber := "KA-04-HN-2001"
	colour := "Green"
	ticket1, _ := parkingLot.CheckIn(regNumber, colour)

	car := parkingLot.CheckOut(ticket1)
	assert.NotNil(t, car)

	assert.Equal(t, regNumber, car.RegNumber)
	assert.Equal(t, colour, car.Colour)

	ticket2, err2 := parkingLot.CheckIn(regNumber, colour)

	assert.Equal(t, ticket2, uint(1))
	assert.Nil(t, err2)
}
