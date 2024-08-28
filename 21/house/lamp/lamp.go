package lamp

import "fmt"

var lampId int = 0

type Lamp interface {
	GetID() int
	GetStatus()

	ScrewIn()
	ScrewOut()
	TurnOn()
	TurnOff()
}

type lamp struct {
	id             int
	state          bool
	hasElectricity bool
}

func NewLamp(hasElectricity bool) Lamp {
	lampId++
	return &lamp{
		id:             lampId,
		state:          false,
		hasElectricity: hasElectricity,
	}
}

func (l *lamp) GetID() int {
	return l.id
}

func (l *lamp) ScrewIn() {
	l.hasElectricity = true
}

func (l *lamp) ScrewOut() {
	l.hasElectricity = false
}

func (l *lamp) TurnOn() {
	l.state = true
}

func (l *lamp) TurnOff() {
	l.state = false
}

func (l *lamp) GetStatus() {
	fmt.Printf("Lamp: %d Light %v\n", l.id, l.state && l.hasElectricity)
}
