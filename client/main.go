package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("transfer-encoding", "chunked") // simulate a stream
		for i := 0; i < 3; i++ {
			w.Write([]byte("chk\n"))
			w.(http.Flusher).Flush()
			<-time.Tick(time.Second * 1)
		}
	}))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	b := make([]byte, 256)
	for {
		s, err := resp.Body.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Printf("chunck: %v\n.\n", string(b[:s]))
	}

	if err != nil {
		log.Fatal(err)
	}
}
