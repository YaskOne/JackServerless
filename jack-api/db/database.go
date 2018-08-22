package db

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// table names constants
const (
	BusinessTable = "BusinessTable"
	UserTable = "UserTable"
	OrderTable = "OrderTable"
)

type Model struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type IDResponse struct {
	ID uint `json:"id"`
}

var database *gorm.DB = nil
//var database *dynamodb.DynamoDB = nil

func DB() *gorm.DB {
	if database == nil {
		initializeDB()
	}
	return database
}

func initializeDB() {

	user := "JackAdmin"
	password := "ArgosBubble3"
	_db := "JackDB"
	//endpoint := "172.31.28.36:3308"
	endpoint := "jackdbmysql.crqo2vw40anm.eu-west-2.rds.amazonaws.com:3308"

	db, err := gorm.Open("mysql", user+":"+password+"@tcp("+endpoint+")/"+_db+"?parseTime=true")

	if err != nil {
		fmt.Println("ERROR: Failed opening aws session")
		fmt.Println(err.Error())
		os.Exit(0)
	}

	database = db

	database.AutoMigrate(&Business{})
	database.AutoMigrate(&Category{})
	database.AutoMigrate(&LatLng{})
	database.AutoMigrate(&Product{})
	database.AutoMigrate(&Order{})
	database.AutoMigrate(&OrderProduct{})
	database.AutoMigrate(&Business{})
}
