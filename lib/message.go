package satms

// Message represent a basic message exchange between clients
type Message struct {
	from    int
	to      int
	content []byte
}
