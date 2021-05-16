package mineral

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MineralType struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Name          string             `json:"name" bson:"name"`
	Hardness      int                `json:"hardness" bson:"hardness"`
	MeltingPoint  float64            `json:"meltingPoint" bson:"meltingPoint"`
	FractureLimit int                `json:"fractureLimit" bson:"fractureLimit"`
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
