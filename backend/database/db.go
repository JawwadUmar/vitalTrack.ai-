package database

import (
	"vita-track-ai/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dbtype string = "postgres"

// const username string = "appuser"
const username string = "postgres"
const password string = "0000"
const dbhost string = "localhost"
const port string = "5433"
const dbname string = "vitadb"
const security string = "sslmode=disable"

// const DSN string = dbtype + "://" + username + ":" + password + "@" + dbhost + ":" + port + "/" + dbname + "?" + security
const dsn string = "host=" + dbhost + " user=" + username + " password=" + password + " dbname=" + dbname + " port=" + port + " " + security

var DB *gorm.DB
var err error

func Init() {
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	createTables()
}

func createTables() {
	createUserTable()
	createFileTable()
}

func createUserTable() {

	err := DB.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to migrate User table")
	}
}

func createFileTable() {
	err := DB.AutoMigrate(&models.File{})
	if err != nil {
		panic("failed to migrate File table")
	}
}
