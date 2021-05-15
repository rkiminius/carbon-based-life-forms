package rabbit

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type SimpleMessage struct {
	Message string `json:"message"`
}
