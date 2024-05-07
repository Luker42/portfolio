package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Cute picture of Tom coming soon!\n")
}
func getPortfolio(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /portfolio request\n")
	io.WriteString(w, "Hello, MAX!\n")
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/portfolio", getPortfolio)

	err := http.ListenAndServe(":3001", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
