package partitioner

import (
	"errors"
)

// client represents a client's binding
type client struct {
	id       int
	callback func(Message) error
}

// newClient returns a new instance of a Client
func newClient(id int, callback func(m Message) error) (*client, error) {
	c := new(client)
	if callback == nil {
		return nil, errors.New("Callback must be passed")
	}

	c.id = id
	c.callback = callback

	return c, nil
}
