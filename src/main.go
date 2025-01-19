package main

import (
	"net/http"
	"sync/atomic"
	"time"
)

var ready int32 = 0 // Atomic readiness state

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("healthy"))
	})

	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&ready) == 1 {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ready"))
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
			w.Write([]byte("not ready"))
		}
	})

	go func() {
		// Simulate a long startup time
		<-time.After(10 * time.Second)
		atomic.StoreInt32(&ready, 1)
	}()

	http.ListenAndServe(":8080", nil)

}
