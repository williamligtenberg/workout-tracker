package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/williamligtenberg/workout-tracker/config"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init(cfg config.Config) {
	mysqlConfig := mysql.Config{
		User:   cfg.DBUser,
		Passwd: cfg.DBPass,
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", cfg.DBHost, cfg.DBPort),
		DBName: cfg.DBName,
	}

	var err error
	DB, err = sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		log.Fatal("[ERROR] Error opening database: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("[ERROR] Error pinging the database: ", err)
	}

	log.Println("[INFO] Database connected")
}
