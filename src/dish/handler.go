package dish

import (
	"bufio"
	"country-generator/src/util"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

func InsertToDB(data []Dish) int64 {
	db, err := sql.Open("mysql", "testapp:maingames@tcp(35.201.131.140:3306)/emak_masak")
	util.CheckErr(err)

	query := `
		INSERT em_dish_details 
		SET 
		name=?,
		description=?,
		category=?,
		price=?,
		status=?,
		picture_url=?,
		created_by=?,
		created_time=?,
		updated_time=?
	`
	stmt, err := db.Prepare(query)
	util.CheckErr(err)

	log.Info("Query : ", query)
	var res sql.Result

	for index, dish := range data {
		if index < 1000 {
			res, err = stmt.Exec(
				dish.Name,
				dish.Description,
				rand.Intn(4-1)+1,
				((dish.LowPrice * 9350) + (dish.HighPrice*9350)/2),
				rand.Intn(2-1)+1,
				"https://media.foody.id/res/g6/54700/prof/s640x400/foody-mobile-nasi-goreng-2-jpg-905-635997833615016836.jpg",
				rand.Intn(1000-1)+1,
				time.Now(),
				time.Now(),
			)
		}
	}

	id, err := res.LastInsertId()
	util.CheckErr(err)

	return id
}

func PopulateDish(filePath string) []Dish {
	// Load a TXT file.
	here, err := filepath.Abs(".")
	util.CheckErr(err)

	filePath = filepath.Join(here, filePath)
	util.CheckErr(err)

	f, err := os.Open(filePath)
	util.CheckErr(err)

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))

	var totalDish []Dish

	for {
		record, err := r.Read()
		if err != io.EOF {
			util.CheckErr(err)
		} else {
			break
		}

		dish := Dish{
			ID:            StringToNumber(NullCheckNumber(record[0])),
			Name:          NullCheckString(record[1]),
			Description:   NullCheckString(record[2]),
			MenuAppeared:  StringToNumber(NullCheckNumber(record[3])),
			TimesAppeared: StringToNumber(NullCheckNumber(record[4])),
			FirstAppeared: NullCheckString(record[5]),
			LastAppeared:  NullCheckString(record[6]),
			LowPrice:      StringToFloat(NullCheckFloat(record[7])),
			HighPrice:     StringToFloat(NullCheckFloat(record[8])),
		}
		totalDish = append(totalDish, dish)
	}

	fmt.Println("Total data : ", len(totalDish))

	return totalDish
}
