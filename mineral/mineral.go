package mineral

import "errors"

// State of Mineral
const MINERAL_STATE_SOLID = "SOLID"         // Mineral that is in its regular state and does not posses fractures
const MINERAL_STATE_LIQUID = "LIQUID"       // Mineral that has been melted. Such Mineral can't posses fractures
const MINERAL_STATE_FRACTURED = "FRACTURED" // Mineral that is in Solid state, but possesses fractures

type Mineral struct {
	ID        int
	Name      string
	State     string
	Fractures int
}

var minerals = []Mineral{
	{
		1,
		"topaz",
		MINERAL_STATE_LIQUID,
		10,
	},
	{
		2,
		"diamond",
		MINERAL_STATE_SOLID,
		100,
	},
}

func GetMinerals() []Mineral {
	return minerals
}

func FindMineralById(mineralId int) (*Mineral, error) {
	for _, value := range minerals {
		if mineralId == value.ID {
			return &value, nil
		}
	}
	return nil, errors.New("mineral not available")
}
