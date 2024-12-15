package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	// Importing `pq` to register the PostgreSQL driver.
	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

func ConnectDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(Red + "Error loading .env file. Please ensure it exists and is configured correctly." + Reset)
		panic(err)
	}

	requiredVars := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME"}
	for _, key := range requiredVars {
		if os.Getenv(key) == "" {
			fmt.Printf(Red+"Missing required environment variable: %s"+Reset+"\n", key)
			panic(fmt.Sprintf("Environment variable %s not set", key))
		}
	}

	host := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Println(Red + "Error converting DB_PORT to integer: " + err.Error() + Reset)
		panic(err)
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, pass)

	db, errSQL := sql.Open("postgres", psqlSetup)
	if errSQL != nil {
		fmt.Println(Red + "Error while opening the database connection: " + errSQL.Error() + Reset)
		panic(errSQL)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(Red + "Database connection failed: " + err.Error() + Reset)
		panic(err)
	}

	DB = db

	fmt.Println(Green + "Successfully connected to the database!" + Reset)
}
