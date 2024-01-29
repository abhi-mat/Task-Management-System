package db

import (
	"database/sql"
	"fmt"
	"log"

	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

const (
	maxOpenConnection    = 10
	maxConnectionTimeout = time.Hour * 1
	maxIdleConnection    = 5
)

var Client *sql.DB

func InitMySQL() error {
	ConnClient, err := getConnection()
	if err != nil {
		return err
	}
	Client = ConnClient
	return nil
}

func getConnection() (*sql.DB, error) {
	Client, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", viper.GetString("database.userName"), viper.GetString("database.password"), viper.GetString("database.host"), viper.GetString("database.port"), viper.GetString("database.dbName")))
	Client.SetMaxOpenConns(maxOpenConnection)
	Client.SetMaxIdleConns(maxIdleConnection)
	Client.SetConnMaxLifetime(maxConnectionTimeout)
	if err != nil {
		return nil, err
	}
	log.Println("MySQL Connection established")
	return Client, nil
}
