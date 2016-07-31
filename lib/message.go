package satms

// Message represent a basic message exchange between clients
type Message struct {
	// Topic of the message, use to send/receive a message only for specific group of message
	Topic string `json:"topic"`
	// From is the id of the sender of the message
	From int `json:"from"`
	// To is a ID of the receiver
	To int `json:"to"`
	// Body is the content of the message
	Body string `json:"body"`
}
