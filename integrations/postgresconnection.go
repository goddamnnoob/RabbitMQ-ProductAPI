package integrations

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetNewPostgresConnection() (connection *sql.DB, erro error) {
	host := os.Getenv("POSTGRES_HOST")
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	user := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DATABASE")

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	connection, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(" Error trying to connect to DB ")
		return nil, err
	}
	return connection, nil
}
