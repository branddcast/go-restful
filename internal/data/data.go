package data

import (
	"database/sql"
	"sync"
)

var (
	data *Data
	once sync.Once
)

// Data manages the connection to the database.
type Data struct {
	DB *sql.DB
}
