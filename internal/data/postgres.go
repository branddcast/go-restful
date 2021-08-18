package data

import (
	"database/sql"
	"io/ioutil"
	"log"

	// registering database driver
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	envs map[string]string
	err  error
)

func getConnection() (*sql.DB, error) {
	host, port, database := envs["DATABASE_HOST"], envs["DATABASE_PORT"], envs["DATABASE_NAME"]
	user, password, sslmode := envs["DATABASE_USER"], envs["DATABASE_PASSWORD"], envs["DATABASE_SSLMODE"]
	uri := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + database + "?sslmode=" + sslmode

	return sql.Open("postgres", uri)
}

func initDB() {
	envs, err = godotenv.Read(".env")

	log.Println(envs)

	if err != nil {
		log.Panic(err)
		log.Fatal("Error loading .env file")
	}

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
	file := envs["DATABASE_MIGRATION_FILE"]
	b, err := ioutil.ReadFile("./database/" + file)

	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))

	if err != nil {
		return err
	}

	return rows.Close()
}
