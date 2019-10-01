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
func newClient(id int, callback func(m Message) error) (client, error) {
	var c client
	if callback == nil {
		return c, errors.New("Callback must be passed")
	}

	c.id = id
	c.callback = callback

	return c, nil
}
