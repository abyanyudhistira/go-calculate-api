package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Response structure
type Response struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

// addHandler calculates
func parseAB(w http.ResponseWriter, r *http.Request) (a, b float64, ok bool) {
	w.Header().Set("Content-Type", "application/json")
	astring := r.URL.Query().Get("a")
	bstring := r.URL.Query().Get("b")
	var errA, errB error
	a, errA = strconv.ParseFloat(astring, 64)
	b, errB = strconv.ParseFloat(bstring, 64)
	if errA != nil || errB != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "Invalid numbers"})
		return 0, 0, false
	}
	return a, b, true
}

// addHandler: /add?a=10&b=5
func addHandler(w http.ResponseWriter, r *http.Request) {
	a, b, ok := parseAB(w, r)
	if !ok {
		return
	}
	json.NewEncoder(w).Encode(Response{Result: a + b})
}

// subHandler: /sub?a=10&b=5
func subHandler(w http.ResponseWriter, r *http.Request) {
	a, b, ok := parseAB(w, r)
	if !ok {
		return
	}
	json.NewEncoder(w).Encode(Response{Result: a - b})
}

// mulHandler: /mul?a=10&b=5
func mulHandler(w http.ResponseWriter, r *http.Request) {
	a, b, ok := parseAB(w, r)
	if !ok {
		return
	}
	json.NewEncoder(w).Encode(Response{Result: a * b})
}

// divHandler: /div?a=10&b=5
func divHandler(w http.ResponseWriter, r *http.Request) {
	a, b, ok := parseAB(w, r)
	if !ok {
		return
	}
	if b == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "Division by zero"})
		return
	}
	json.NewEncoder(w).Encode(Response{Result: a / b})
}

func main() {
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/sub", subHandler)
	http.HandleFunc("/mul", mulHandler)
	http.HandleFunc("/div", divHandler)
	fmt.Println("Server starting on :8081...")
	http.ListenAndServe(":8081", nil) // [10]
}
