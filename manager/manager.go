package manager

type Manager struct {
	Name 	string
}

func (m Manager) GetAvailableMinerals () ([]Mineral, error) {

	minerals := []Mineral{
		{
			"topaz",
			MINERAL_STATE_LIQUID,
			10,
		},
		{
			"diamond",
			MINERAL_STATE_SOLID,
			100,
		},
	}

	return minerals, nil
}

//func (m Manager) ReceiveAction (mineral Mineral, Action) error {
//
//	return nil
//}

func (m Manager) PerformActions ([]Mineral) error {

	return nil
}

// Manager must be able to send a task request to the Factory
func (m Manager) SendTaskToFactory () error {
	var factory Factory

	return nil
}
