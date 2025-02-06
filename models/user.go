package models

import "example.com/db"

type User struct {
	ID       int64
	Email    string `binding:"required"`
	password string `binding:"required"`
}

func (u User) Save() error{
	query := `insert into users (email, password)values (?,?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil{
		return err
	}

	defer stmt.Close()

	result,err:=stmt.Exec(u.Email, u.password)

	if err != nil{
		return err
	}

	userId, err:= result.LastInsertId()

	u.ID = userId
	return err
}