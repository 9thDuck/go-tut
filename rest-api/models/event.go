package models

import (
	"fmt"
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"date_time"`
	UserID      int       `json:"user_id"`
}

func (e *Event) Save() error {
	query := `INSERT INTO events
			  (name, description, location, date_time, user_id)
	 		  VALUES(?, ?, ?, ?, ?)`

	statement, err := db.Db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	res, err := statement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	if err != nil {
		return err
	}

	ID, err := res.LastInsertId()

	if err != nil {
		return err
	}

	e.ID = ID

	return nil
}

func GetAllEvents() (*[]Event, error) {
	query := `SELECT id, name, description, location, date_time, user_id FROM events`
	rows, err := db.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return &events, nil
}

func GetEventById(id int64) (*Event, error) {

	query := `SELECT id, name, description, location, date_time, user_id from events WHERE id = ?`

	row := db.Db.QueryRow(query, id)

	// if err != nil {
	// 	return Event{}, err
	// }

	// defer row.Close()

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) Update() error {
	query := `UPDATE events SET name=?, description=?, location=?, date_time=?, user_id=? where id=?`

	res, err := db.Db.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.UserID, event.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("error: Couldn't update the event with id:%v", event.ID)
	}
	return nil
}

func DeleteEventById(eventId int64) error {
	query := `DELETE FROM events WHERE id=?`

	res, err := db.Db.Exec(query, eventId)

	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		return fmt.Errorf("error: Couldn't update the event with id:%v", eventId)
	}
	return nil
}
