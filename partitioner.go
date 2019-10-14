package partitioner

import (
	"context"
	"errors"
	"math/rand"
	"strings"
	"sync"
)

// Partitioner represents an instance which controls
// the distribution of massages to clients
type Partitioner struct {
	clients map[int]*client
	input   chan Message
	keys    map[string]int
	mtx     sync.Mutex
}

// AddClient adds a new Client to the instance
func (p *Partitioner) AddClient(ctx context.Context, callback func(Message) error) error {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	newID := len(p.clients)
	c, err := newClient(newID, callback)
	if err != nil {
		return err
	}
	p.clients[c.id] = c

	return nil
}

// IncomeMessage sends to Partitioner's stream a new message to redistribute
func (p *Partitioner) IncomeMessage(message Message) error {
	if strings.TrimSpace(message.id) == "" {
		return errors.New("Id is required")
	}
	p.input <- message
	return nil
}

func (p *Partitioner) routeMessage(message Message) {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	if message.ID() == "" {
		return
	}

	clientID, ok := p.keys[message.ID()]
	if ok {
		p.clients[clientID].callback(message)
		return
	}

	var id int

	if t := len(p.clients); t > 1 {
		id = rand.Intn(t)
	}

	p.keys[message.ID()] = id

	p.clients[id].callback(message)
}

// NewPartitioner returns a new instance of Partitioner
func NewPartitioner(ctx context.Context) (*Partitioner, error) {
	var partitioner Partitioner
	partitioner.clients = make(map[int]*client)
	partitioner.input = make(chan Message, 1)
	partitioner.keys = make(map[string]int)

	go func(innerCtx context.Context, p *Partitioner) {
		defer close(p.input)

		for {
			select {
			case m := <-p.input:
				p.routeMessage(m)
			case <-innerCtx.Done():
				return
			}
		}
	}(ctx, &partitioner)

	return &partitioner, nil
}
