package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func connectionString(dbName, host, port, usrName, usrPass string) (string, error) {
	if usrName == "" {
		return "", fmt.Errorf("userName should not be empty")
	}

	if usrPass == "" {
		return "", fmt.Errorf("userPass should not be empty")
	}

	if host == "" {
		return "", fmt.Errorf("host should not be empty")
	}

	if port == "" {
		port = "3306"
	}

	if dbName == "" {
		return "", fmt.Errorf("dbName should not be empty")
	}

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true", usrName, usrPass, host, port, dbName), nil
}

func loadEnv() (string, string, string, string, string) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	return os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_USRNAME"), os.Getenv("DB_USRPASS")
}

func NewMySQLConn() (*sql.DB, error) {
	dbName, host, port, usrName, usrPass := loadEnv()
	DSN, err := connectionString(dbName, host, port, usrName, usrPass)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", DSN)
	if err != nil {
		return nil, err
	}

	return db, nil
}
