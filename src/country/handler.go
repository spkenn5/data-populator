package country

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	util "country-generator/src/util"

	log "github.com/sirupsen/logrus"
)

func populateCountries() {
	log.Info("Populating countries...")
	url := "https://restcountries.eu/rest/v2/all"

	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}

	log.Info("Requesting ", url)
	response, err := netClient.Get(url)
	util.CheckErr(err)

	bytes, err := ioutil.ReadAll(response.Body)
	util.CheckErr(err)

	var countries []Country
	err = json.Unmarshal(bytes, &countries)
	util.CheckErr(err)

	log.Info("Succesfully inserted : ", insertCountries(countries))
}

func insertCountries(data []Country) int64 {
	db, err := sql.Open("mysql", "kenjidb:maingames@/grubber")
	util.CheckErr(err)

	query := `
		INSERT grubber_location 
		SET 
		name=?,
		calling_code=?,
		web_domain=?,
		alpha2_code=?,
		alpha3_code=?,
		capital=?,
		region=?,
		latitude=?,
		longitude=?,
		currency_code=?,
		currency_name=?,
		currency_symbol=?,
		language=?,
		language_native=?,
		japanese_translations=?,
		picture_url=?,
		created_time=?,
		updated_time=?
	`
	stmt, err := db.Prepare(query)
	util.CheckErr(err)

	log.Info("Query : ", query)
	var res sql.Result
	for _, country := range data {
		latitude := 0.0
		longitude := 0.0

		if len(country.LatLong) > 0 {
			latitude = country.LatLong[0]
			longitude = country.LatLong[1]
		}

		countryCode := ""
		countryName := ""
		countrySymbol := ""
		if len(country.Currencies) > 0 {
			countryCode = country.Currencies[0].Code
			countryName = country.Currencies[0].Name
			countrySymbol = country.Currencies[0].Symbol
		}

		languageName := ""
		languageNative := ""
		if len(country.Languages) > 0 {
			languageName = country.Languages[0].Name
			languageNative = country.Languages[0].NativeName
		}

		res, err = stmt.Exec(
			country.Name,
			fmt.Sprintf("+%s", country.CallingCodes[0]),
			country.TopLevelDomain[0],
			country.Alpha2Code,
			country.Alpha3Code,
			country.Capital,
			country.Region,
			latitude,
			longitude,
			countryCode,
			countryName,
			countrySymbol,
			languageName,
			languageNative,
			country.Translations.JA,
			country.Flag,
			time.Now(),
			time.Now(),
		)
		util.CheckErr(err)
	}

	id, err := res.LastInsertId()
	util.CheckErr(err)

	fmt.Println(id)
	return id
}
