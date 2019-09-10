package partitioner

import (
	"context"
	"errors"
)

// Client represents a client's binding
type Client struct {
	id       int
	callback func(Message) error
}

// NewClient returns a new instance of a Client
func NewClient(ctx context.Context, id int, input <-chan Message, callback func(m Message) error) (Client, error) {
	var client Client
	if callback == nil {
		return client, errors.New("Callback must be passed")
	}

	client.id = id
	client.callback = callback

	go func(cl *Client) {
		for {
			select {
			case m := <-input:
				cl.callback(m)
			case <-ctx.Done():
				return
			}
		}
	}(&client)

	return client, nil
}
