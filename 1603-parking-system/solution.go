package _1603_parking_system

type ParkingSystem struct {
	slotBig    int
	slotMedium int
	slotSmall  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		slotBig:    big,
		slotMedium: medium,
		slotSmall:  small,
	}
}

func (ps *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if ps.slotBig > 0 {
			ps.slotBig--
			return true
		}
	case 2:
		if ps.slotMedium > 0 {
			ps.slotMedium--
			return true
		}
	case 3:
		if ps.slotSmall > 0 {
			ps.slotSmall--
			return true
		}
	}

	return false
}
