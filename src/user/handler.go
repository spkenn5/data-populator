package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	util "country-generator/src/util"

	log "github.com/sirupsen/logrus"
)

func InsertUsersToDB(data Results) int64 {
	db, err := sql.Open("mysql", "testapp:maingames@tcp(35.201.131.140:3306)/emak_masak")
	util.CheckErr(err)

	query := `
		INSERT em_user_details 
		SET 
		email=?,
		username=?,
		first_name=?,
		last_name=?,
		address=?,
		city=?,		
		state=?,
		postal_code=?,
		dob=?,
		phone_number=?,
		gender=?,
		rating=?,
		picture_url=?,
		created_time=?,
		updated_time=?
	`
	stmt, err := db.Prepare(query)
	util.CheckErr(err)

	log.Info("Query : ", query)
	var res sql.Result

	for index, user := range data.Results {
		if index < 1000 {
			res, err = stmt.Exec(
				user.Email,
				fmt.Sprintf("%s%s", user.Name.First, user.Nationality),
				user.Name.First,
				user.Name.Last,
				user.Location.Street,
				user.Location.City,
				user.Location.State,
				fmt.Sprintf("%d", rand.Intn(99999-10000)+10000),
				user.DOB,
				user.Phone,
				user.Gender,
				rand.Float64(),
				user.PictureURL.Large,
				time.Now(),
				time.Now(),
			)
			util.CheckErr(err)
		}
	}

	id, err := res.LastInsertId()
	util.CheckErr(err)

	return id
}

func PopulateUsers() Results {
	log.Info("Populating users...")
	url := "https://randomuser.me/api/?results=1000"

	var netClient = &http.Client{
		Timeout: time.Second * 20,
	}

	log.Info("Requesting ", url)
	response, err := netClient.Get(url)
	util.CheckErr(err)

	bytes, err := ioutil.ReadAll(response.Body)
	util.CheckErr(err)

	var users Results
	err = json.Unmarshal(bytes, &users)
	util.CheckErr(err)

	return users
}
