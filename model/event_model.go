package model

type Event struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type EventID struct {
	Id int `json:"id"`
}

type eventMap struct {
	x, y int
}
