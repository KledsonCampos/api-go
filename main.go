package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Event model info
// @Description Event information about the event
type Event struct {
	Id          uuid.UUID `json:"uuid" swaggertype:"uuid"`
	Title       string    `json:"title" swaggertype:"string"`
	Description string    `json:"description" swaggertype:"string"`
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

func main() {
	log.Println("Starting API")
	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/events", GetAllEvents).Methods("GET")
	router.HandleFunc("/events", CreateNewEvent).Methods("POST")
	// port := os.Getenv("PORT")
	// http.ListenAndServe(":"+port, router)
	http.ListenAndServe(":8086", router)
}

// ShowAllEvents godoc
// @Summary      Show all events
// @Description  List all events
// @Tags         events
// @Accept       json
// @Produce      json
// @Param   request	body	Event	false  "Event"
// @Success      200  {object}  Event

// @Router       /events [get]
func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	log.Println("Acesando o endpoint get all events")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessando health-check")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Aplicação em execução!\n")
}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessando home")
	fmt.Fprintf(w, "HOME!\n")
}

// CreateEvent godoc
// @Summary      Create new event
// @Description  Create new event
// @Tags         events
// @Accept       json
// @Produce      json
// @Param   request	body	Event	false  "Event"
// @Success      201  {object}  Event

// @Router       /events [post]
func CreateNewEvent(w http.ResponseWriter, r *http.Request) {
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
