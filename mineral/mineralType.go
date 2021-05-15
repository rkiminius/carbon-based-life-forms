package mineral

import "errors"

type MineralType struct {
	Name          string
	Hardness      int
	MeltingPoint  float64
	FractureLimit int
}

var mineralTypes = []MineralType{
	{
		Name:          "topaz",
		Hardness:      200,
		MeltingPoint:  1000,
		FractureLimit: 32,
	},
	{
		Name:          "diamond",
		Hardness:      1500,
		MeltingPoint:  5000,
		FractureLimit: 8,
	},
}

func FindMineralTypeByName(name string) (*MineralType, error) {
	for _, value := range mineralTypes {
		if name == value.Name {
			return &value, nil
		}
	}
	return nil, errors.New("mineral type not available")
}
