package main

import (
	"database/sql"
	"fmt"
)

//type UserRepository interface {
//	GetByID(id int) (*User, error)
//}

const (
	pgHost     = "localhost"
	pgUser     = "pguser"
	pgPassword = "daskfljdas"
	pgDB       = "playground"
	pgPort     = 5432
)

func main() {
	connectionCfg := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", pgHost, pgUser, pgPassword, pgDB, pgPort)
	db, err := sql.Open("postgres", connectionCfg)

	if err != nil {
		panic(err)
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	//// service not implement yet, but it presents the idea how repository pattern work
	//userRepository := NewPostgresUserRepository(db)
	//userService := newService(userRepository)
	//
	//user, err := userService.GetUserByID(1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("user: %+v\n", user)
}
