package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zorrokid/film-db-rest-api/data/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	connection *gorm.DB
	logger     *log.Logger
}

func NewDatabase(logger *log.Logger) *Database {
	//	conn, err := connect()
	//if err != nil {
	//logger.Fatal(err)
	//}
	return &Database{logger: logger, connection: nil}
}

func (db *Database) GetConnection() *gorm.DB {
	if db.connection == nil {
		db.logger.Println("Connecting db")
		conn, err := db.connect()
		if err != nil {
			db.logger.Fatal(err)
		}
		db.connection = conn
		db.initDB()
	}
	return db.connection
}

func (db *Database) initDB() {
	db.connection.AutoMigrate(&models.Movie{})
}

func (db *Database) connect() (*gorm.DB, error) {

	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlUser := os.Getenv("MYSQL_USER")
	mysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDb := os.Getenv("MYSQL_DB")

	if mysqlHost == "" {
		return nil, fmt.Errorf("No MYSQL_HOST provided")
	}
	if mysqlUser == "" {
		return nil, fmt.Errorf("No MYSQL_USER provided")
	}
	if mysqlPassword == "" {
		return nil, fmt.Errorf("No MYSQL_PASSWORD provided")
	}
	if mysqlDb == "" {
		return nil, fmt.Errorf("No MYSQL_DB provided")
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", mysqlUser, mysqlPassword, mysqlHost, mysqlDb)

	db.logger.Println(connStr)

	dbconn, err := sql.Open("mysql", connStr)

	if err != nil {
		return nil, err
	}

	dbconn.SetConnMaxLifetime(time.Minute * 3)
	dbconn.SetMaxOpenConns(10)
	dbconn.SetMaxIdleConns(10)

	gdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn: dbconn,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return gdb, nil
}
