package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// seed the random number generator with the current time
func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	http.HandleFunc("/mockme", mockme)
	http.ListenAndServe(":8080", nil)
}

// This gives us a go struct to interact with the incoming mockme request
type mockmeMessage struct {
	Message string `json:"message"`
}

// This receives the request, modifies the text, and sends it back
func mockme(w http.ResponseWriter, r *http.Request) {
	// Check to make sure the correct content type was submitted
	headerContentType := r.Header.Get("Content-Type")
	if headerContentType != "application/json" {
		log.Fatal("Wrong content-type used")
	}

	// initialize the needed struct in the function
	var m mockmeMessage

	// decode the json body so we can interact with it and validate proper json
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&m)
	if err != nil {
		log.Fatal(err)
	}
	// This line will return the "message" _value_ from the body sent
	// fmt.Fprintln(w, m.Message)

	// Get just the text we want to modify and return out of the json object
	usersText := m.Message

	// initialize the variable for the outgoing text
	var outputTxt string

	for i := 0; i < len(usersText); i++ {
		if (usersText[i] >= 'a' && usersText[i] <= 'z') || (usersText[i] >= 'A' && usersText[i] <= 'Z') {
			v := rand.Intn(10)
			if v > 5 {
				outputTxt += strings.ToUpper(string(usersText[i]))
			} else {
				outputTxt += strings.ToLower(string(usersText[i]))
			}
		} else {
			outputTxt += string(usersText[i])
		}
	}
	fmt.Fprintln(w, outputTxt)
}
