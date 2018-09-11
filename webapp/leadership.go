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
	phrases := []string{
		"Needs more Bias for Action",
		"Doesn't sound very Customer Obsessed",
		"This is a Disagree and Commit moment for me",
	}

	// seed/initialize the psuedo random num generator
	rand.Seed(time.Now().Unix())
	// grab a random index from phrases
	return phrases[rand.Intn(len(phrases))]
}
