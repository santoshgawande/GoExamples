package main

import (
	"log"
	"net/http"
)

func apiResponse(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"Hello World"}`))

}

func main(){
	http.HandleFunc("/",apiResponse)
	log.Fatal(http.ListenAndServe(":8080", nil))
}