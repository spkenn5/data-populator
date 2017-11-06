package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"

	"country-generator/src/dish"
	"country-generator/src/user"
)

var filePath string

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	filePath = "files/emak/dish.csv"
}

func main() {
	lastID := dishHandler()
	fmt.Println("LastID: ", lastID)
}

func dishHandler() int64 {
	dishList := dish.PopulateDish(filePath)
	lastID := dish.InsertToDB(dishList)

	return lastID
}

func userHandler() int64 {
	userList := user.PopulateUsers()
	lastID := user.InsertUsersToDB(userList)

	return lastID
}
