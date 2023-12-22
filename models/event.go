package models

import "time"

type Event struct {
	ID int 
	Name string
	Description string
	Location string
	DateTime time.Time
	UserID int
}

var events = []Event{}

func (e Event) Save() {
	// later: add ot to database
	events = append(events, e)
}

func GetAllEvents() []Event {

}