package mineral

import "errors"

const MINERAL_STATE_SOLID = "SOLID" // Mineral that is in its regular state and does not posses fractures
const MINERAL_STATE_LIQUID = "LIQUID" // Mineral that has been melted. Such Mineral can't posses fractures
const MINERAL_STATE_FRACTURED = "FRACTURED" // Mineral that is in Solid state, but possesses fractures

type Mineral struct {
	Name 		string
	State 		string
	Fractures 	int
}

// this action would split the Mineral in half, doubling its current amount of fractures
func (m *Mineral) Fracture () error {
	return nil
}

// this action would attempt to melt a Mineral and turn it to Liquid state
func (m *Mineral) Melt () error {
	if m.State == MINERAL_STATE_LIQUID {
		return errors.New("Mineral state already in liquid stage")
	}

	m.State = MINERAL_STATE_LIQUID
	return nil
}

// this action would attempt to solidify a Mineral and turn it to Solid state
func (m *Mineral) Condense () error {
	return nil
}
