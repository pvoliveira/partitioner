package partitioner

import (
	"context"
	"testing"
)

func TestAddClient(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p, err := NewPartitioner(ctx)

	if err != nil {
		t.Error(err)
		return
	}

	cb := func(m Message) error {
		t.Logf("Message received: ID %s", m.ID())
		return nil
	}

	err = p.AddClient(ctx, cb)
	if err != nil {
		t.Error(err)
		return
	}

	m, err := NewMessage("1", nil, []byte("test"))

	err = p.IncomeMessage(m)
	if err != nil {
		t.Error(err)
		return
	}
}
