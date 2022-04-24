package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Event model info
// @Description Event information about the event
type Event struct {
	Id          uuid.UUID `json:"uuid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type allEvents []Event

var events = allEvents{}

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
// @BasePath /

func main() {
	log.Println("Starting API")
	router := mux.NewRouter()
	router.HandleFunc("/", home)
	router.HandleFunc("/health-check", healthCheck).Methods("GET")
	router.HandleFunc("/event/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/event", getAllEvents).Methods("GET")
	router.HandleFunc("/event", createNewEvent).Methods("POST")
	router.HandleFunc("/event/{id}", updateEvent).Methods("PUT")
	router.HandleFunc("/event/{id}", deleteEvent).Methods("DELETE")

	port := os.Getenv("PORT")
	http.ListenAndServe(":"+port, router)
	// http.ListenAndServe(":8086", router)
}

// @Summary      Show all events
// @Description  List all events
// @Tags         events
// @Accept       json
// @Produce      json
// @Param   request	body	Event	false  "Event"
// @Success      200  {object}  Event

// @Router       /events [get]
// GetAllEvents
func getAllEvents(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessando o endpoint get all event")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

// HealthCheck
func healthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessando health-check")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Aplicação em execução!\n")
}

// Home
func home(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessando home")
	fmt.Fprintf(w, "HOME!\n")
}

// @Summary      Create new event
// @Description  Create new event
// @Tags         events
// @Accept       json
// @Produce      json
// @Param   request	body	Event	false  "Event"
// @Success      201  {object}  Event

// @Router       /events [post]
// createNewEvent
func createNewEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessando create new event")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var newEvent Event

	newEvent.Id = uuid.New()

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(reqBody, &newEvent)

	events = append(events, newEvent)

	json.NewEncoder(w).Encode(events)
}

// @Summary Get a Event by ID
// @ID get-one-event
// @Produce json
// @Param id path string true "event ID"
// @Success 200 {object} todo
// @Failure 404 {object} message
// @Router /evetnt/{id} [get]
// getOneEvent
func getOneEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("Get one event")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	eventID := mux.Vars(r)["id"]

	for _, singleEvent := range events {
		if singleEvent.Id.String() == eventID {
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

// updateEvent
func updateEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("Atualizando event")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	eventID := mux.Vars(r)["id"]
	var updatedEvent Event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range events {
		if singleEvent.Id.String() == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}

func deleteEvent(w http.ResponseWriter, r *http.Request) {
	log.Println("Deletando event")
	w.Header().Set("Content-Type", "application/json")

	eventID := mux.Vars(r)["id"]

	index, searchEvent := events.SearchEventById(eventID)

	if searchEvent {
		events = append(events[:index], events[index+1:]...)
		log.Printf("The event with ID %v has been deleted successfully", eventID)
		w.WriteHeader(http.StatusOK)

	} else {
		log.Printf("The event with ID %v not found", eventID)
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (e allEvents) SearchEventById(id string) (int, bool) {

	for i, singleEvent := range e {
		if singleEvent.Id.String() == id {
			return i, true
		}
	}
	return 0, false
}
