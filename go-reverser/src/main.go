package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/reverse/{input}", reverseHandler).Methods("GET")

	http.Handle("/", r)

	fmt.Println("Server is running on :8001")
	http.ListenAndServe(":8001", nil)
}

func reverseHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	input := vars["input"]
	reversed := reverseString(input)
	fmt.Fprintf(w, "%s", reversed)
}

func reverseString(input string) string {
	chars := []rune(input)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}
