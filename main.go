package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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

// @title           Swagger event-API
// @version         1.0
// @description     This is a document event api.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API-event Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      https://api-go-events.herokuapp.com/

func main() {
	log.Println("Starting API")
	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/events", GetAllEvents).Methods("GET")
	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)
}

// ShowAllEvents godoc
// @Summary      Show all events
// @Description  List all events
// @Tags         events
// @Accept       json
// @Produce      json
// @Success      200  {object}  Event

// @Router       /events [get]
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
