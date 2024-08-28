package main

import (
	"21/house"
)

func main() {
	h := house.InitHouse(true, 5)
	h.TurnOnLampByID(2)
	h.TurnOnLampByID(5)
	h.ShowLamps()
}
