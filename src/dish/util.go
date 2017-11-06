package dish

import (
	"country-generator/src/util"
	"strconv"
)

func StringToNumber(number string) int {
	num, err := strconv.Atoi(number)
	util.CheckErr(err)

	return num
}

func StringToFloat(number string) float64 {
	num, err := strconv.ParseFloat(number, 64)
	util.CheckErr(err)

	return num
}

func NullCheckFloat(number string) string {
	if number != "" {
		return number
	} else {
		return "0.0"
	}
}

func NullCheckNumber(number string) string {
	if number != "" {
		return number
	} else {
		return "0"
	}
}

func NullCheckString(str string) string {
	if str != "" {
		return str
	} else {
		return "N/A"
	}
}
