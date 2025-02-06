package models

import (
	"example.com/db"
	"example.com/utils"
)

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

	hashedPassword,err :=utils.HashPassword(u.password)

	if err != nil{
		return err
	}
	result,err:=stmt.Exec(u.Email, hashedPassword)

	if err != nil{
		return err
	}

	userId, err:= result.LastInsertId()

	u.ID = userId
	return err
}