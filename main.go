package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Event struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type allEvents []Event

var events = allEvents{
	{
		Id:          1,
		Title:       "APIs",
		Description: "Trocando conhecimento com a turma da pós-graduação.",
	},
}

func main() {
	log.Println("Starting API")
	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/events", GetAllEvents).Methods("GET")

	http.ListenAndServe(":8086", router)
}

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessamdo o endpoint get all events")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessamdo health-check")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Aplicação em execução!\n")
}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessamdo home")
	fmt.Fprintf(w, "HOME!\n")
}
