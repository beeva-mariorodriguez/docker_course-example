package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const version = "0.4"

func main() {
	log.Println("starting hellogo ...")
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/health", health).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println(r.UserAgent())
	w.WriteHeader(http.StatusOK)
	h, _ := os.Hostname()
	fmt.Fprintf(w, "%s > hello GO %s\n", h, version)
}

func health(w http.ResponseWriter, r *http.Request) {
	log.Println(r.UserAgent())
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}
