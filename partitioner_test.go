package partitioner

import (
	"context"
	"strconv"
	"testing"
	"time"
)

func TestIncomeMessage(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	p, err := NewPartitioner(ctx)

	if err != nil {
		t.Error(err)
		return
	}

	clientReceived := make(chan bool)
	defer close(clientReceived)

	cb := func(m Message) error {
		t.Logf("Message received: ID %s", m.ID())
		clientReceived <- true
		return nil
	}

	err = p.AddClient(ctx, cb)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Clients: %v\n", p.clients)

	m, err := NewMessage("1", nil, []byte("test"))

	err = p.IncomeMessage(m)
	if err != nil {
		t.Error(err)
		return
	}

	select {
	case <-clientReceived:
	case <-ctx.Done():
		t.Error("Client does not received the message")
	}
}

func TestIncomeMessageMultipleClients(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	p, err := NewPartitioner(ctx)

	if err != nil {
		t.Error(err)
		return
	}

	cb1 := func(m Message) error {
		t.Logf("Message received client1: ID %s\n", m.ID())
		return nil
	}

	cb2 := func(m Message) error {
		t.Logf("Message received client2: ID %s\n", m.ID())
		return nil
	}

	err = p.AddClient(ctx, cb1)
	if err != nil {
		t.Error(err)
		return
	}

	err = p.AddClient(ctx, cb2)
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Clients: %v\n", p.clients)

	for i := 0; i < 100; i++ {
		m, err := NewMessage(strconv.Itoa(i), nil, []byte("test"))
		if err != nil {
			t.Error(err)
			return
		}

		err = p.IncomeMessage(m)
		if err != nil {
			t.Error(err)
			return
		}
	}

	<-time.After(time.Millisecond * 2000)
}
