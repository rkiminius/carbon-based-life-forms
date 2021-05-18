package mineral

import (
	"github.com/rkiminius/carbon-based-life-forms/config"
	"github.com/rkiminius/carbon-based-life-forms/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func Test_getByMineralTypeId(t *testing.T) {
	config.GetConfig("../test.yaml")
	db.MongoConnect()

	mineralTypesIds := []string{
		"5bd0e70b9db9ea0011519bd1",
		"5bd0e70b9db9ea0011519bd0",
	}

	for _, id := range mineralTypesIds {
		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			t.Error(err)
			return
		}

		mineralType, err := getByMineralTypeId(objId)
		if err != nil {
			t.Error(err)
			return
		}

		if mineralType == nil {
			t.Error("Expected mineralType, got nil")
			return
		}

		if mineralType.ID != objId {
			t.Errorf("retrieved mineralType.ID (%v) != expected %v", objId, mineralType)
			return
		}
	}
}
