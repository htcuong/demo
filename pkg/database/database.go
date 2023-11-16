package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/htcuong/demo/config"
)

func Connection(config config.Database) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn(config))
	if err != nil {
		log.Printf("error %s during the open db\n", err)
		return nil, err
	}
	return db, nil
}

func dsn(config config.Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", config.Username, config.Password, config.Hostname, config.DBname)
}
