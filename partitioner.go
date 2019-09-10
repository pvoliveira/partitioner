package partitioner

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// Partitioner represents an instance which controls
// the distribution of massages to clients
type Partitioner struct {
	mtx     *sync.Mutex
	clients map[int]*Client
	Input   chan Message
	server  *http.Server
}

func (p Partitioner) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "connect" && r.Method == http.MethodPost {
		b, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {

		}
		p.addClient(context.Background(), string(b))
	}
}

// addClient adds a new Client to the instance
func (p *Partitioner) addClient(ctx context.Context, webhookURL string) error {
	p.mtx.Lock()
	defer p.mtx.Unlock()

	// for now this will be setted here,
	// but the consumer of Partitioner should say whats he want to do
	cb := func(m Message) error {
		r, errCb := http.NewRequest("POST", webhookURL, m.Body())
		if errCb != nil {
			return errCb
		}

		r.Header.Set("x-partitioner-id", string(m.ID()))

		for _, k := range m.headers {
			r.Header.Set(fmt.Sprintf("x-partitioner-%s", k), m.headers[k])
		}

		ctxCb, cancel := context.WithTimeout(r.Context(), time.Millisecond*50)
		defer cancel()

		r = r.WithContext(ctxCb)

		_, errCb = http.DefaultClient.Do(r)
		if errCb != nil {
			return errCb
		}

		return nil
	}

	newID := len(p.clients) + 1
	c, err := NewClient(ctx, newID, p.Input, cb)
	if err != nil {
		return err
	}
	p.clients[c.id] = &c

	return nil
}

func NewPartitioner(addr string, inputCallback func(chan<- Message) error) (*Partitioner, error) {
	var partitioner Partitioner
	partitioner.Input = make(chan Message, 1)

	partitioner.server = &http.Server{
		Addr:           addr,
		Handler:        partitioner,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := partitioner.server.ListenAndServe()
	if err != nil {

	}

	return &partitioner, nil
}
