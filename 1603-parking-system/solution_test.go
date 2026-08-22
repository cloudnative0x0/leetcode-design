package _1603_parking_system

import "testing"

func TestParkingSystem_AddCarBig(t *testing.T) {
	ps := Constructor(1, 0, 0)

	if !ps.AddCar(1) {
		t.Errorf("expected true for big car, got false")
	}
	if ps.AddCar(1) {
		t.Errorf("expected false for second big car, got true")
	}
}

func TestParkingSystem_AddCarMedium(t *testing.T) {
	ps := Constructor(0, 2, 0)

	if !ps.AddCar(2) {
		t.Errorf("expected true for first medium car, got false")
	}
	if !ps.AddCar(2) {
		t.Errorf("expected true for second medium car, got false")
	}
	if ps.AddCar(2) {
		t.Errorf("expected false for third medium car, got true")
	}
}

func TestParkingSystem_AddCarSmall(t *testing.T) {
	ps := Constructor(0, 0, 3)

	for i := 0; i < 3; i++ {
		if !ps.AddCar(3) {
			t.Errorf("expected true for small car %d, got false", i+1)
		}
	}
	if ps.AddCar(3) {
		t.Errorf("expected false for fourth small car, got true")
	}
}

func TestParkingSystem_AddCarMixed(t *testing.T) {
	ps := Constructor(2, 1, 0)

	if !ps.AddCar(1) {
		t.Errorf("expected true for big car 1")
	}
	if !ps.AddCar(2) {
		t.Errorf("expected true for medium car")
	}
	if !ps.AddCar(1) {
		t.Errorf("expected true for big car 2")
	}
	if ps.AddCar(1) {
		t.Errorf("expected false for third big car")
	}
	if ps.AddCar(2) {
		t.Errorf("expected false for second medium car")
	}
}

func TestParkingSystem_AddCarInvalidType(t *testing.T) {
	ps := Constructor(1, 1, 1)

	if ps.AddCar(0) {
		t.Errorf("expected false for invalid type 0")
	}
	if ps.AddCar(4) {
		t.Errorf("expected false for invalid type 4")
	}
}

func TestParkingSystem_AddCarZeroSpots(t *testing.T) {
	ps := Constructor(0, 0, 0)

	if ps.AddCar(1) {
		t.Errorf("expected false for big car with zero spots")
	}
	if ps.AddCar(2) {
		t.Errorf("expected false for medium car with zero spots")
	}
	if ps.AddCar(3) {
		t.Errorf("expected false for small car with zero spots")
	}
}

func TestParkingSystem_AddCarFullAfterPartial(t *testing.T) {
	ps := Constructor(1, 1, 1)

	if !ps.AddCar(1) {
		t.Errorf("expected true for big")
	}
	if !ps.AddCar(2) {
		t.Errorf("expected true for medium")
	}
	if !ps.AddCar(3) {
		t.Errorf("expected true for small")
	}
	if ps.AddCar(1) {
		t.Errorf("expected false after full")
	}
}
