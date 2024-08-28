package house

import (
	"21/house/lamp"
)

type House interface {
	ShowLamps()

	ElectricityOn()
	ElectricityOff()
	TurnOnLampByID(id int)
}

type house struct {
	hasElectricity bool
	lamps          []lamp.Lamp
}

func InitHouse(hasElectricity bool, lampNumber int) House {
	lamps := make([]lamp.Lamp, lampNumber)
	for i := range lampNumber {
		lamps[i] = lamp.NewLamp(hasElectricity)
	}
	return &house{
		hasElectricity: hasElectricity,
		lamps:          lamps,
	}
}

func (h *house) ElectricityOn() {
	h.hasElectricity = true
	for _, l := range h.lamps {
		l.ScrewIn()
	}
}

func (h *house) ElectricityOff() {
	h.hasElectricity = false
	for _, l := range h.lamps {
		l.ScrewOut()
	}
}

func (h *house) TurnOnLampByID(id int) {
	for _, l := range h.lamps {
		if l.GetID() == id {
			l.TurnOn()
			return
		}
	}
}

func (h *house) ShowLamps() {
	for _, l := range h.lamps {
		l.GetStatus()
	}
}
