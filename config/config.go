package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	user := "root"
	pass := ""
	host := ""
	port := "3306"
	dbname := "p2ngc2"
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, pass, host, port, dbname)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}
