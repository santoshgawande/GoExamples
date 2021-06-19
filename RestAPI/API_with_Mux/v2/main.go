package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// API using net/http
// type server struct{}

// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	switch r.Method {
// 	case "GET":
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"message":"get called`))
// 	case "POST":
// 		w.WriteHeader(http.StatusCreated)
// 		w.Write([]byte(`{"message":"post called"`))
// 	case "PUT":
// 		w.WriteHeader((http.StatusAccepted))
// 		w.Write([]byte(`{"message":"put called`))
// 	case "DELETE":
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"message":"delete called"`))
// 	default:
// 		w.WriteHeader(http.StatusNotFound)
// 		w.Write([]byte(`{"message":"not found"}`))
// 	}
// }

// func main() {
// 	// Add Mux to pass parameters
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", home)
// 	log.Fatal(http.ListenAndServe(":8080", r))
// }

// Using MUX
// GET
func get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"get called"}`))

}

// POST
func post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"post called"}`))
}

// PUT
func put(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message":"put called"}`))

}

//DELETE
func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"delete called"}`))
}

// Not Found
func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message":"not found"}`))
}

func main() {
	// SubRouter
	r := mux.NewRouter()
	r.HandleFunc("/", get).Methods(http.MethodGet)
	r.HandleFunc("/", post).Methods(http.MethodPost)
	r.HandleFunc("/", put).Methods(http.MethodPut)
	r.HandleFunc("/", delete).Methods(http.MethodDelete)
	r.HandleFunc("/", notFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}
