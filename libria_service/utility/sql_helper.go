package utility

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDbClient() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		"localhost", 3306, "postgres", "postgres", "libria")
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(fmt.Sprintf("Verbindung mit Datenbank kann nicht hergestellt werden: %v", err))
	}
	conn.SetMaxIdleConns(2)
	conn.SetMaxOpenConns(20)
	return conn
}

func GetStringValue(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}
