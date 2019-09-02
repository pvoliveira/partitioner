package partitioner

import (
	"bytes"
	"io"
	"fmt"
)

// Message is the content which Partitioner receive as input
// and redirects to correspondent client
type Message struct {
	id string
	headers map[string]string
	body []byte
}

// ID returns the id of Message (used for partitioning)
func(m *Message) ID() (id string) {
	return m.id
}

// Body returns a io.Reader to the content body's Message
func(m *Message) Body() (body io.Reader) {
	return bytes.NewReader(m.body)
}

func NewMessage(id string, headers map[string]string, body []byte) (message Message) {
	message = Message{
		id: id, 
		headers: headers, 
		body: body,
	}

	return message
}

// Client represents a client's binding
type Client struct {
	id int
	buffer chan<- Message
	callback func(Message) error
}

// Connect connects a stream input (channel) to the buffer of client
func(c *Client) Connect(input <-chan Message) (err error) {
	if input == nil {
		return fmt.Errorf("Input channel can not be nil")
	}
	
	if c.buffer == nil {
		return fmt.Errorf("Client not configured")
	}

	go func() {
		for m := range input {
			c.buffer <- m
		}
	}()

	return nil
}

// Key is the association of a Client with a partition key
type Key struct { 
	id string
	client *Client
}

// Partitioner represents an instance which controls 
// the distribution of massages to clients
type Partitioner struct {
	clients map[int]*Client
	keys map[string]*Key
}