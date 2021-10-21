package service

import "errors"


type ParkingLot struct {
	currentCapacity uint
	size int
	parking map[uint]*Car
}


func NewParkingLot(size int) ParkingLot {
	return ParkingLot{
		currentCapacity: 0,
		parking: make(map[uint]*Car),
		size: size,
	}
}

func (lot *ParkingLot) CheckIn(regNumber string, colour string) (uint, error) {
	slot := lot.currentCapacity+1

	if slot > uint(lot.size) {
		return uint(1), errors.New("parking is full")
	}
	lot.parking[slot] = &Car{RegNumber: regNumber, Colour: colour}
	lot.currentCapacity++
	return slot, nil
}

func (lot *ParkingLot) CheckOut(ticket uint) *Car {
	car := lot.parking[ticket]
	delete(lot.parking, uint(ticket))
	lot.currentCapacity--
	return car
}

type Car struct {
	RegNumber string
	Colour string
}