package db

import (
	"database/sql"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var (
	conn *sql.DB
	once sync.Once
)

func GetConnection() (*sql.DB, error) {
	var err error
	once.Do(func() {
		connStr := "host=" + os.Getenv("DB_HOST") +
			" user=" + os.Getenv("DB_USER") +
			" password=" + os.Getenv("DB_PASSWORD") +
			" dbname=" + os.Getenv("DB_NAME") +
			" sslmode=require"
		conn, err = sql.Open("postgres", connStr)
		if err == nil {
			conn.SetMaxOpenConns(1) // mesma lógica: 1 conexão por container
		}
	})
	return conn, err
}
