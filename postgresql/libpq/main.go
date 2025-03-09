package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	pgHost     = "localhost"
	pgUser     = "pguser"
	pgPassword = "daskfljdas"
	pgDB       = "playground"
	pgPort     = 5432
)

// User represents a table user
type User struct {
	ID        int
	FirstName string
	LastName  string
}

func testReadPg(db *sql.DB) {
	selectQuery := `SELECT * FROM users`

	rows, err := db.Query(selectQuery)
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var user User

		err = rows.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("ID: %d, First Name: %s, Last Name: %s\n", user.ID, user.FirstName, user.LastName)
	}
}

func testReadSpecificPg(db *sql.DB, id int) (User, error) {
	selectQuery := `SELECT * FROM users WHERE id = $1`

	var user User
	err := db.QueryRow(selectQuery, id).Scan(&user.ID, &user.FirstName, &user.LastName)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func testInsertPg(db *sql.DB, firstName, lastName string) {
	//var user User

	// 1st approach return an ID
	//query := `
	//INSERT INTO users (first_name, last_name)
	//VALUES ($1, $2) RETURNING id`
	//
	//err := db.QueryRow(query, firstName, lastName).Scan(&user.ID)

	// 2nd approach without returning an ID
	query := `
	INSERT INTO users (first_name, last_name) 
	VALUES ($1, $2)
	`
	execResult, err := db.Exec(query, firstName, lastName)
	if err != nil {
		fmt.Println("Failed to insert data:", err)
		return
	}

	numRow, err := execResult.RowsAffected()
	if err != nil {
		fmt.Println("Failed to get RowsAffected:", err)
		return
	}

	//fmt.Println("Created user:", user.ID)
	fmt.Println("Number of affected row:", numRow)
}

func testDeletePg(db *sql.DB, id int) {
	delQuery := `DELETE FROM users WHERE id = $1`

	execResult, err := db.Exec(delQuery, id)
	if err != nil {
		fmt.Println("Failed to delete data:", err)
		return
	}

	numRow, err := execResult.RowsAffected()
	if err != nil {
		fmt.Println("Failed to get RowsAffected:", err)
		return
	}

	fmt.Println("Number of affected row:", numRow)
}

func testUpdatePg(db *sql.DB, id int, firstName, lastName string) {
	updateQuery := `UPDATE users 
SET first_name = $2, last_name = $3 
WHERE id = $1`

	execResult, err := db.Exec(updateQuery, id, firstName, lastName)
	if err != nil {
		fmt.Printf("Failed to update data at id: %d with error: %s", id, err)
		return
	}

	numRow, err := execResult.RowsAffected()
	if err != nil {
		fmt.Println("Failed to get RowsAffected:", err)
		return
	}

	fmt.Println("Number of affected row:", numRow)
}

func main() {
	// 1st approach: using full url of postgres.
	//db, err := sql.Open("postgres", "postgres://pguser:aafsav@@localhost:5432/playground?sslmode=disable")

	// 2nd approach: only define the postgres field seperated by white space
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", pgHost, pgUser, pgPassword, pgDB, pgPort))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			return
		}
	}(db)

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Successfully connected to PostgreSQL!")
}
