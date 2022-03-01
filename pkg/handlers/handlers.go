package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func Start(http.ResponseWriter, *http.Request) {
	fmt.Println("starting")
	for {
		for i := 0; i < 100000000; i++ {
		}
	}
}

func Stop(http.ResponseWriter, *http.Request) {
	fmt.Println("stopping")
	os.Exit(200)
}
