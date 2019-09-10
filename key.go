package partitioner

// Key is the association of a Client with a partition key
type Key struct {
	id     string
	client *Client
}