package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // Not sure why we have to do this
)

var (
	db gorm.DB
)

// Conn return global db connection if exists
func Conn() gorm.DB {
	return db
}

// Init opens the connection and init tables
func Init() gorm.DB {
	fmt.Printf("DB init\n")

	newDbConn, err := openConnection()
	checkErr(err, "DB open failed")
	db = newDbConn

	db.AutoMigrate(&Project{}, &Lib{}, &ProjectLib{})

	return db
}

// Truncate deletes everything
func Truncate() {
	db.Where("ID > 0").Delete(Project{})
	db.Where("ID > 0").Delete(ProjectLib{})
	db.Where("ID > 0").Delete(Lib{})
}

func openConnection() (gorm.DB, error) {
	return gorm.Open("postgres", "user=root dbname=gomdm sslmode=disable")
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
