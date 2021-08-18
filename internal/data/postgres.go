package data

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	// registering database driver
	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {
	host, port, database := os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_NAME")
	user, password, sslmode := os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_SSLMODE")
	uri := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + database + "?sslmode=" + sslmode

	return sql.Open("postgres", uri)
}

func initDB() {
	db, err := getConnection()
	if err != nil {
		log.Panic(err)
	}

	err = MakeMigration(db)
	if err != nil {
		log.Panic(err)
	}

	data = &Data{
		DB: db,
	}
}

func New() *Data {
	once.Do(initDB)
	return data
}

func Close() error {
	if data == nil {
		return nil
	}

	return data.DB.Close()
}

func MakeMigration(db *sql.DB) error {
	file := os.Getenv("DATABASE_MIGRATION_FILE")
	b, err := ioutil.ReadFile("../../database/" + file)

	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))

	if err != nil {
		return err
	}

	return rows.Close()
}
