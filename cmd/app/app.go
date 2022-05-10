package app

import (
	"event_api/conf"
	"event_api/model/repositories"
	"event_api/pkg/handlers"
	"log"
	"net/http"
)

func Start() {
	db, err := conf.GetDB()
	if err != nil {
		log.Fatalf(err.Error())
	}

	eventRepo := repositories.NewEventRepository(db)
	eventHandler := handlers.NewEventHandler(eventRepo)

	http.HandleFunc("/getlistevent", eventHandler.GetListEvent)
	http.HandleFunc("/getitemevent", eventHandler.GetItemEvent)

	//
	log.Fatal(http.ListenAndServe(":8080", nil))

}
