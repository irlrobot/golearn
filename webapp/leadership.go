package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, randomPhrase())
}

func randomPhrase() string {
	// The phrases() function is in phrases.go.
	phrases := phrases()

	// Seed/initialize the psuedo random number generator.
	rand.Seed(time.Now().Unix())
	// Return a random phrase.
	return phrases[rand.Intn(len(phrases))]
}
