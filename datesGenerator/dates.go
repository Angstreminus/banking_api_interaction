package datesGenerator

import "strconv"

func GenerateDates() []string {
	var dates []string
	var day string = ""
	for i := 1; i <= 31; i++ {
		if i < 10 {
			day = "0" + string(strconv.FormatInt(int64(i), 10)+"/03/"+"2023")
		} else {
			day = string(strconv.FormatInt(int64(i), 10) + "/03/" + "2023")
		}

		dates = append(dates, day)
	}

	for i := 1; i <= 30; i++ {
		if i < 10 {
			day = "0" + string(strconv.FormatInt(int64(i), 10)+"/04/"+"2023")
		} else {
			day = string(strconv.FormatInt(int64(i), 10) + "/04/" + "2023")
		}

		dates = append(dates, day)
	}

	for i := 1; i <= 31; i++ {
		if i < 10 {
			day = "0" + string(strconv.FormatInt(int64(i), 10)+"/05/"+"2023")
		} else {
			day = string(strconv.FormatInt(int64(i), 10) + "/05/" + "2023")
		}

		dates = append(dates, day)
	}
	return dates
}
