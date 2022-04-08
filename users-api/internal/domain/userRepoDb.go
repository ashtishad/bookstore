package domain

import (
	"database/sql"
	"github.com/ashtishad/bookstore/lib"
	"log"
)

const (
	sqlGetUser         = `SELECT id, name, gender,date_of_birth, email,city, date_created, status FROM users WHERE id=$1;`
	sqlInsertUser      = `INSERT INTO users(name, gender, date_of_birth, email, city) VALUES($1, $2, $3, $4, $5) RETURNING id;`
	sqlUpdateUser      = `UPDATE users SET name=$1, email=$2, city=$3 ,status=$4 WHERE id=$5;`
	sqlCheckUserExists = `SELECT id FROM users WHERE id=$1;`
	sqlDeleteUser      = `DELETE FROM users WHERE id=$1;`
	sqlSearchUser      = `SELECT id, name, email, date_created, status FROM users WHERE name=$1;`
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
	switch err {
	case sql.ErrNoRows:
		log.Printf("Error while scanning users by id : %s", err.Error())
		return nil, lib.NewNotFoundError("user not found in database")
	case nil:
		return &u, nil
	default:
		log.Printf("Error while scanning users by id : %s", err.Error())
		return nil, lib.NewInternalServerError("error while scanning user", err)
	}
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

// Update updatable fields of a user
func (d UserRepoDb) Update(u User) (*User, lib.RestErr) {
	// check user exists
	if check, err := d.checkUserExists(u.Id); check == false {
		return nil, lib.NewInternalServerError("user not found in database.", err)
	}

	_, err := d.db.Exec(sqlUpdateUser, u.Name, u.Email, u.City, u.Status, u.Id)
	if err != nil {
		return nil, lib.NewInternalServerError("Could not update user : ", err)
	}
	// retrieve updated user
	return d.FindById(u.Id)
}

// Delete deletes a user from the database
func (d UserRepoDb) Delete(id int64) lib.RestErr {
	// check user exists
	if check, err := d.checkUserExists(id); check == false {
		return lib.NewInternalServerError("user not found in database,can't delete.", err)
	}

	_, err := d.db.Exec(sqlDeleteUser, id)
	if err != nil {
		return lib.NewInternalServerError("Could not delete user : ", err)
	}
	return nil
}

// FindByName searches by user's full name, multiple records can be found
func (d UserRepoDb) FindByName(name string) (*[]User, lib.RestErr) {
	rows, err := d.db.Query(sqlSearchUser, name)
	if err != nil {
		return nil, lib.NewInternalServerError("Could not search user : ", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Printf("Error while closing rows : %s", err.Error())
		}
	}(rows)

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Id, &u.Name, &u.Email, &u.DateCreated, &u.Status); err != nil {
			return nil, lib.NewInternalServerError("Could not scan user : ", err)
		}
		users = append(users, u)
	}
	if len(users) == 0 {
		return nil, lib.NewNotFoundError("no user found")
	}
	return &users, nil
}

// CheckUserExists checks if a user exists in the database
func (d UserRepoDb) checkUserExists(id int64) (bool, error) {
	row := d.db.QueryRow(sqlCheckUserExists, id)

	var u User
	err := row.Scan(&u.Id)
	if err == sql.ErrNoRows {
		return false, err
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
