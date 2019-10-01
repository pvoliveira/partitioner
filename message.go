package partitioner

import (
	"bytes"
	"errors"
	"io"
	"strings"
)

// Message is the content which Partitioner receive as input
// and redirects to correspondent client
type Message struct {
	id      string
	headers map[string]string
	body    []byte
}

// ID returns the id of Message (used for partitioning)
func (m Message) ID() string {
	return m.id
}

// Body returns a io.Reader to the content body's Message
func (m Message) Body() io.Reader {
	return bytes.NewReader(m.body)
}

// NewMessage builds a Message
func NewMessage(id string, headers map[string]string, body []byte) (Message, error) {
	if strings.TrimSpace(id) == "" {
		return Message{}, errors.New("id is required")
	}

	m := Message{
		id:      id,
		headers: headers,
		body:    body,
	}

	return m, nil
}
