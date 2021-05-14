package client

type Client struct {
	Name		string
	Minerals 	[]Mineral
}


// Clients must be able to ask Manager about their Minerals
func (c Client) AskMinerals () ([]Mineral, error) {

	var man Manager
	minerals, err := man.GetAvailableMinerals()
	if err != nil {
		return nil, err
	}

	return minerals, nil
}

// Clients must be able to request Manager to perform Actions on selected Minerals
func (c Client) PerformActionsOnMinerals ([]Mineral) (error) {
	var manager Manager
	minerals, err := manager.GetAvailableMinerals()
	if err != nil {
		return err
	}

	err = manager.PerformActions(minerals)
	if err != nil {
		return err
	}

	return nil
}