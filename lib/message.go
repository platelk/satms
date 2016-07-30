package satms

// Message represent a basic message exchange between clients
type Message struct {
	From int    `json:"from"`
	To   int    `json:"to"`
	Body string `json:"body"`
}
