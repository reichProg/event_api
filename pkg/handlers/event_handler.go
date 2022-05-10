package handlers

import (
	"encoding/json"
	"event_api/model"
	"event_api/model/repositories"
	"net/http"
)

func NewEventHandler(eventRepository repositories.EventRepositoryI) *EventHandler {
	return &EventHandler{
		eventRepository: eventRepository,
	}
}

type EventHandlerI interface {
	GetListEvent(w http.ResponseWriter, r *http.Request)
	GetItemEvent(w http.ResponseWriter, r *http.Request)
}

type EventHandler struct {
	eventRepository repositories.EventRepositoryI
}

func (e EventHandler) GetListEvent(w http.ResponseWriter, r *http.Request) {
	event, err := e.eventRepository.GetListEvent()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}
	jEvent, err := json.Marshal(*event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(jEvent)
}

func (e EventHandler) GetItemEvent(w http.ResponseWriter, r *http.Request) {

	var eventID model.EventID
	err := json.NewDecoder(r.Body).Decode(&eventID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	eventS, err := e.eventRepository.GetItemEvent(eventID.Id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return
	}

	jEvent, err := json.Marshal(*eventS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jEvent)
}
