package main

import (
	"fmt"
	"sync"
)

type ParkingSystem struct {
	big     int
	medium  int
	small   int
	mutex   *sync.Mutex
	engaged map[int]int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	res := ParkingSystem{big, medium, small, &sync.Mutex{}, make(map[int]int)}
	return res
}

func (p *ParkingSystem) occupy(carType int) (bool, error) {
	if carType != 1 && carType != 2 && carType != 3 {
		return false, fmt.Errorf("error car type of %d", carType)
	}

	used := p.engaged[carType]
	if carType == 1 && used == p.big {
		return false, nil
	} else if carType == 2 && used == p.medium {
		return false, nil
	} else if carType == 3 && used == p.small {
		return false, nil
	}
	p.engaged[carType] = used + 1
	return true, nil
}

func (p *ParkingSystem) AddCar(carType int) bool {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	ok, _ := p.occupy(carType)
	return ok
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
func main() {

}
