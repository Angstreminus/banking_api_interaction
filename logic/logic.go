package logic

import (
	"strconv"
	"strings"
	dateparser "task/dateParser"
)

type Result struct {
	MaxValuteValue float64
	MaxValuteName  string
	MinValuteName  string
	MinValuteValue float64
	Avg            float64
	MaxDate        string
	MinDate        string
}

func GetMaxMinAVGValute(Reports []dateparser.ValCurs) Result {
	var (
		res           Result
		MaxValuteName string
		MinValuteName string
		MinDate       string
		MaxDate       string
		min                   = float64(999999999.)
		max                   = float64(0.)
		avg           float64 = float64(0.)
	)
	for _, dayReport := range Reports {
		for _, valuteReport := range dayReport.Valute {
			//array of Valutes
			currentValuteValue := valuteReport.Value
			if strings.Count(currentValuteValue, ",") != 0 {
				currentValuteValue = strings.ReplaceAll(currentValuteValue, ",", ".")
			}
			floatcurrentValuteValue, _ := strconv.ParseFloat(currentValuteValue, 64)
			// getting Min
			if floatcurrentValuteValue < min {
				min = floatcurrentValuteValue
				MinValuteName = valuteReport.Name
				MinDate = dayReport.Date
			}
			//getting Max
			if floatcurrentValuteValue > max {
				max = floatcurrentValuteValue
				MaxValuteName = valuteReport.Name
				MaxDate = dayReport.Date
			}
			// counting AVG
			avg += floatcurrentValuteValue
		}
	}
	res.MaxValuteValue = max
	res.MinValuteValue = min
	res.MaxValuteName = MaxValuteName
	res.MinValuteName = MinValuteName
	res.MaxDate = MaxDate
	res.MinDate = MinDate
	res.Avg = (avg / 90) / 34
	return res
}
