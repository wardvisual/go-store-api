package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/wardvisual/go-store-api/cmd/api"
	"github.com/wardvisual/go-store-api/configs"
	"github.com/wardvisual/go-store-api/database"
)

func main() {
	cfg := mysql.Config{
		User:                 configs.Envs.User,
		Passwd:               configs.Envs.Password,
		Addr:                 configs.Envs.Address,
		DBName:               configs.Envs.Name,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := database.NewMySQLStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB Status: Connected")
}
