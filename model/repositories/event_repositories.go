package repositories

import (
	"database/sql"
	"event_api/model"
)

type EventRepositoryI interface {
	GetListEvent() (*[]model.Event, error)
	GetItemEvent(id int) (*model.Event, error)
	//	GetMapList(long, lat int) *[]model.Event
}

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (e EventRepository) GetListEvent() (*[]model.Event, error) {
	rows, err := e.db.Query("select * from event")
	if err != nil {
		return nil, err
	}
	var event []model.Event
	for rows.Next() {
		var eventItems model.Event
		rows.Scan(&eventItems.Id, &eventItems.Title, &eventItems.Description, &eventItems.Image)
		event = append(event, eventItems)
	}
	return &event, nil
}

func (e EventRepository) GetItemEvent(id int) (*model.Event, error) {
	rows, err := e.db.Query("select * from event where id = ?", id)
	if err != nil {
		return &model.Event{}, err
	}
	var event model.Event
	for rows.Next() {
		rows.Scan(&event.Id, &event.Title, &event.Description, &event.Image)
	}
	return &event, nil
}

//func (e EventRepository) GetMapList(long, lat int) *[]model.Event {
//	//TODO implement me
//	panic("implement me")
//}
