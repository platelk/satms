package satms

// Message represent a basic message exchange between clients
type Message struct {
	// From is the id of the sender of the message
	From int `json:"from"`
	// To is a ID of the receiver
	To int `json:"to"`
	// Body is the content of the message
	Body string `json:"body"`
}
