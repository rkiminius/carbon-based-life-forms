package mineral

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// State of Mineral
const MINERAL_STATE_SOLID = "SOLID"         // Mineral that is in its regular state and does not posses fractures
const MINERAL_STATE_LIQUID = "LIQUID"       // Mineral that has been melted. Such Mineral can't posses fractures
const MINERAL_STATE_FRACTURED = "FRACTURED" // Mineral that is in Solid state, but possesses fractures

type Mineral struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UUID      string             `json:"uuid" bson:"uuid"`
	Name      string             `json:"name" bson:"name"`
	State     string             `json:"state" bson:"state"`
	Fractures int                `json:"fractures" bson:"fractures"`
}
