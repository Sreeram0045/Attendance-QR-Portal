package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// db is a pointer to a struct named DB from the sql module
var db *sql.DB

// envVariables is a user-defined struct to hold the data from the .env file
type envVariables struct {
	Username string
	Password string
	DBName   string
}

// envVarStruct is a pointer variable that points to the user-defined struct
var envVarStruct *envVariables

func checkingPath() string {

	comparison := "/Attendance-QR-Portal"
	var dir string = "Attendance-QR-Portal"
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return ""
	}

	slices := strings.SplitAfter(rootDir, "/")

	if len(slices) > 0 {
		if dir == slices[len(slices)-1] {
			return rootDir
		}
	}

	if strings.Contains(rootDir, comparison) {
		slice := strings.SplitAfter(rootDir, comparison)
		fmt.Printf("%s,%s\n", slice[0], slice[1])
		return slice[0]
	}
	return ""
}

// envLoadFunction reads the .env file and returns its contents to the pointer
func envLoadFunction() (config *envVariables) {
	path := checkingPath()
	viper.AddConfigPath(path)
	// Specify the path of the .env file
	// viper.AddConfigPath("../../")
	// Specify the name of the .env file
	viper.SetConfigName("db_connection")
	// Specify the type of the configuration file (env for environment variables)
	viper.SetConfigType("env")
	// Attempt to read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Unmarshal the loaded environment variables into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}

func ConnectToDB() {
	// Load the environment variables from the .env file
	envVarStruct = envLoadFunction()

	// Configure the MySQL connection using the loaded environment variables
	config := mysql.Config{
		User:                 envVarStruct.Username,
		Passwd:               envVarStruct.Password,
		Net:                  "tcp",
		Addr:                 "127.0.1.1:3306",
		DBName:               envVarStruct.DBName,
		AllowNativePasswords: true,
	}

	// Attempt to open a connection to the MySQL database
	var err error
	db, err = sql.Open("mysql", config.FormatDSN())

	// Check for errors when opening the database connection
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database to ensure the connection is alive
	pingErr := db.Ping()

	// Check for errors when pinging the database
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	// Print a success message if the database connection is successful
	fmt.Println("Connected to Database Successfully")
}
