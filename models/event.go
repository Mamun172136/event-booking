package models

import (
	"fmt"
	"time"

	"example.com/db"
)

type Event struct {
    ID          int64     `json:"id"`
    Name        string    `json:"name" binding:"required"`
    Description string    `json:"description" binding:"required"`
    Location    string    `json:"location" binding:"required"`
    DateTime    time.Time `json:"dateTime" binding:"required" time_format:"2006-01-02T15:04:05Z07:00"`
    UserID      int64     `json:"userId"`
}

var events = []Event{}

func (e *Event) Save() error {
    query := `INSERT INTO events(name, description, location, dateTime, user_id)
              VALUES(?, ?, ?, ?, ?)`
    
    result, err := db.DB.Exec(query, 
        e.Name,
        e.Description,
        e.Location,
        e.DateTime, // SQL driver will handle time conversion
        e.UserID,
    )
    if err != nil {
        return fmt.Errorf("database error: %w", err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return fmt.Errorf("failed to get last insert ID: %w", err)
    }
    
    e.ID = id
    return nil
}

func GetAllEvents() ([]Event,error){
	query := "SELECT * FROM events"
	rows,err:=db.DB.Query(query)
	if err != nil{
		return nil,err
	}
	defer rows.Close()

	var events []Event

	for rows.Next(){
		var event Event
		err:=rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err !=nil{
			return nil,err
		}
		events = append(events, event)
	}
	

	return events, nil
}

func GetEventById(id int64) (*Event, error){
	query := "SELECT * FROM events WHERE id=?"
	row:=db.DB.QueryRow(query, id)

	var event Event
	err:=row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil{
		return nil, err
	}

	return &event, nil
}

// func (event Event)Update()error{
// 	query :=`
// 	UPDATE events
// 	SET name=?, description=?, location=?, dateTime=?, 
// 	WHERE id=?
// 	`
// 	stmt, err:= db.DB.Prepare(query)

// 	if err != nil{
// 		return err
// 	}

// 	defer stmt.Close()
// 	_,err =stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
// 	return err
// }

func (event *Event) Update() error {
    query := `
    UPDATE events 
    SET name = ?, description = ?, location = ?, dateTime = ?
    WHERE id = ?
    `
    
    result, err := db.DB.Exec(query,
        event.Name,
        event.Description,
        event.Location,
        event.DateTime,
        event.ID,
    )
    
    if err != nil {
        return fmt.Errorf("database error: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("failed to check rows affected: %w", err)
    }

    if rowsAffected == 0 {
        return fmt.Errorf("no rows were updated")
    }

    return nil
}

// func (event Event)Delete()error{
// 	query := "DELETE FROM events WHERE ID =?"
// 	stmt , err :=db.DB.Prepare(query)

// 	if err != nil{
// 		return err
// 	}

// 	defer stmt.Close()

// 	_,err = stmt.Exec(event.ID)
// 	return err
// }

func (event *Event) Delete() error {
    result, err := db.DB.Exec("DELETE FROM events WHERE id = ?", event.ID)
    if err != nil {
        return err
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return fmt.Errorf("no rows deleted")
    }

    return nil
}

// func (e Event) Register(userId int64)error{
// 	query := "INSERT INTO  registrations(event_id, user_id) VALUES (?,?)"
// 	stmt, err := db.DB.Prepare(query)

// 	if err != nil {
// 		return err
// 	}

// 	defer stmt.Close()

// 	_, err = stmt.Exec(e.ID, userId)

// 	return err
// }

func (e *Event) Register(userId int64) error {
    _, err := db.DB.Exec(
        "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)", 
        e.ID, userId,
    )
    return err
}