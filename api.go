package main

import (
	"fmt"
	dateparser "task/dateParser"
	"task/datesGenerator"
	"task/logic"
)

func main() {
	dates := datesGenerator.GenerateDates()
	allPeriodReports := dateparser.SendAndParse(dates)
	result := logic.GetMaxMinAVGValute(allPeriodReports)

	fmt.Println("Max valute name is: ", result.MaxValuteName, " cource ", result.MaxValuteValue, " date: ", result.MaxDate)
	fmt.Println("Min valute name is: ", result.MinValuteName, " cource ", result.MinValuteValue, " date: ", result.MinDate)
	fmt.Println("Avg ruble cource is ", result.Avg)
}
