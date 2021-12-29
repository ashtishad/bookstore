package domain

import (
	"database/sql"
	"github.com/ashtishad/bookstore/lib"
	"log"
)

const (
	sqlGetUser    = `SELECT id, name, gender,date_of_birth, email,city, date_created, status FROM users WHERE id=$1;`
	sqlInsertUser = `INSERT INTO users(name, gender, date_of_birth, email, city) VALUES($1, $2, $3, $4, $5) RETURNING id;`
)

type UserRepoDb struct {
	db *sql.DB
}

// NewUserRepoDb creates a new customer repository
func NewUserRepoDb(client *sql.DB) UserRepoDb {
	return UserRepoDb{client}
}

// FindById returns a customer by id
func (d UserRepoDb) FindById(id int64) (*User, lib.RestErr) {
	// Note: Select * would supply data on db table order, order would mismatch with struct fields
	row := d.db.QueryRow(sqlGetUser, id)

	var u User
	err := row.Scan(&u.Id, &u.Name, &u.Gender, &u.DateOfBirth, &u.Email, &u.City, &u.DateCreated, &u.Status)
	if err == sql.ErrNoRows {
		log.Printf("Error while scanning users by id : %s", err.Error())
		return nil, lib.NewNotFoundError("user not found in database")
	}
	// catch other errors that might occur
	if err != nil {
		log.Printf("Error while scanning users by id : %s", err.Error())
		return nil, lib.NewUnexpectedError("Unexpected database error")
	}

	return &u, nil
}

// Save inserts a user into the database
func (d UserRepoDb) Save(u User) (*User, lib.RestErr) {
	log.Println("Inserting user :", u)

	var id int64
	err := d.db.QueryRow(sqlInsertUser, u.Name, u.Gender, u.DateOfBirth, u.Email, u.City).Scan(&id)
	if err != nil {
		return nil, lib.NewInternalServerError("Could not insert id : ", err)
	}
	log.Println("New user id is:", id)

	u.Id = id
	return &u, nil
}
