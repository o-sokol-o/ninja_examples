package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type User struct {
	ID           int64  `json:"id"`
	Name         string `json:"name,omitempty"`
	Email        string
	Password     string
	RegisteredAt time.Time
}

func main() {

	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres sslmode=disable password=qwerty123")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var u User
	err = db.QueryRow("select * from users where id = $1", 2).
		Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("no rows")
			return
		}
		log.Fatal(err)
	}

	fmt.Println(u)

	// rows, err := db.Query("select * from users")
	// if err != nil {
	// log.Fatal(err)
	// }
	// defer rows.Close()

	// users := make([]User, 0)
	// for rows.Next()
	// u := User{}
	// err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)
	// if err != nil {
	// log.Fatal(err)
	// }

	// users = append(users, u)
	// }

	// err = rows.Err()
	// if err != nil {

	// err = insertUser(db, User{
	// Name:     "Petya",
	// Email:    "petya@gmail.com",
	// Password: "wjsdggjkgjhff3g2f23727gtds78gds7g8*^&56t6dsa",
	// })

	// users, err := getUsers(db)
	// if err != nil {
	// log.Fatal(err)
	// }

	// fmt.Println(users)

	err = insertUser(db, User{
		Name:  "Гриша",
		Email: "grisha@ninja.go",
	})

	users, err := getUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)

	// fmt.Println(users)
}

func getUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.Name, &u.Email, &u.Password, &u.RegisteredAt)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func getUserBylD(db *sql.DB, id int) (User, error) {
	var u User
	err := db.QueryRow("select * from users where id = $1", 1).
		Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisteredAt)

	return u, err
}

func insertUser(db *sql.DB, u User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	_, err = tx.Exec("insert into users (name, email, password) values ($1, $2, $3)",
		u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}

	_, err = tx.Exec("insert into logs (entity, action) values ($1, $2)",
		"user", "created")
	if err != nil {
		return err
	}

	return tx.Commit()
}

func deleteUser(db *sql.DB, id int) error {
	_, err := db.Exec("delete from users where id = $1", id)

	return err
}

func updateUser(db *sql.DB, id int, newUser User) error {
	_, err := db.Exec("update users set name=$1, email=$2 where id = $3",
		newUser.Name, newUser.Email, id)

	return err
}
