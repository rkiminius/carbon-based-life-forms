package mineral

// Mineral actions
const MINERAL_ACTION_TYPE_FRACTURE = "FRACTURE"
const MINERAL_ACTION_TYPE_MELT = "MELT"
const MINERAL_ACTION_TYPE_CONDENSE = "CONDENSE"

type ActionType string

func (a ActionType) IsValid() bool {
	if a == MINERAL_ACTION_TYPE_CONDENSE ||
		a == MINERAL_ACTION_TYPE_FRACTURE ||
		a == MINERAL_ACTION_TYPE_MELT {
		return true
	}
	return false
}
