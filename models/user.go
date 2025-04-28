package models

import (
	"errors"
	"fmt"
	"strings"

	"example.com/db"
	"example.com/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error{
	query := `insert into users (email, password)values (?,?)`
	stmt, err := db.DB.Prepare(query)

	if err != nil{
		return err
	}

	defer stmt.Close()

	hashedPassword,err :=utils.HashPassword(u.Password)

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

// func (u *User) ValidateCredentials()error{
// 	query := "SELECT  id,password, FROM users WHERE email =?"
// 	row := db.DB.QueryRow(query, u.Email)

// 	var retrievedPassword string
// 	err :=row.Scan(&u.ID,&retrievedPassword)

// 	if err != nil{
// 		return err
// 	}

// 	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

// 	if !passwordIsValid{
// 		return errors.New("credentials invalid")
// 	}
// 	return nil
// }

func (u *User) ValidateCredentials() error {
    u.Email = strings.TrimSpace(u.Email)
    fmt.Printf("Attempting login for: %s\n", u.Email)

    query := "SELECT id, password FROM users WHERE email = ?"
    row := db.DB.QueryRow(query, u.Email)

    var retrievedPassword string
    err := row.Scan(&u.ID, &retrievedPassword)
    
    if err != nil {
        fmt.Printf("Database error: %v\n", err)
        return errors.New("invalid credentials")
    }

    fmt.Printf("Stored hash: %s\nInput pass: %s\n", retrievedPassword, u.Password)
    fmt.Printf("Hash length: %d\n", len(retrievedPassword))

    if err := bcrypt.CompareHashAndPassword([]byte(retrievedPassword), []byte(u.Password)); err != nil {
        fmt.Printf("Password mismatch: %v\n", err)
        return errors.New("invalid credentials")
    }

    return nil
}