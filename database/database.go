package database

import (
	"os"
	"database/sql"
	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	db, err := sql.Open("postgres", "user=" + os.Getenv("DATABASE_USERNAME") +
		" password='" + os.Getenv("DATABASE_PASSWORD") +
		"' host=" + os.Getenv("DATABASE_HOST") +
		" port=5432 dbname=" + os.Getenv("DATABASE_NAME")  + " sslmode=disable")
	if err != nil {
		panic(err)
	}

	return db
}
