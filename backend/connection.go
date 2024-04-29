package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"

	"github.com/spf13/viper"
)

var db *sql.DB

type envVariables struct {
	Username     string
	Password     string
	DataBaseName string
}

var envVarStruct *envVariables

func envLoadFunction() (config *envVariables) {
	viper.AddConfigPath("../")
	viper.SetConfigName("db_connection")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env variables into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}

func main() {
	envVarStruct = envLoadFunction()
	config := mysql.Config{
		User:                 envVarStruct.Username,
		Passwd:               envVarStruct.Password,
		Net:                  "tcp",
		Addr:                 "127.0.1.1:3306",
		DBName:               envVarStruct.DataBaseName,
		AllowNativePasswords: true,
	}
	// config := mysql.Config{
	// 	User:                 "sreeram",
	// 	Passwd:               "sreeram",
	// 	Net:                  "tcp",
	// 	Addr:                 "127.0.1.1:3306",
	// 	DBName:               "sample_db",
	// 	AllowNativePasswords: true,
	// }
	var err error
	db, err = sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected to Database Succesfully")

}
