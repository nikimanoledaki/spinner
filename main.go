package main

import (
	"net/http"
	"os"
	"sync/atomic"
)

var spinning int32

func main() {
	http.HandleFunc("/spin", func(w http.ResponseWriter, r *http.Request) {
		myVal := atomic.AddInt32(&spinning, 1)
		go func() {
			for {
				isSpinning := atomic.LoadInt32(&spinning)
				if isSpinning < myVal {
					break
				}
				// spin!
			}
		}()
	})

	http.HandleFunc("/unspin", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&spinning, -1)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, nil)
}
