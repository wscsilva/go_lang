package util

import (
	"fmt"
	"strconv"
	"time"
)

// ConvertDateToString converts a time.Time object to a string in the specified format.
func ConvertDateToString(date time.Time, format string) string {
	return date.Format(format)
}

// ConvertStringToDate converts a string to a time.Time object using the specified format.
func ConvertStringToDate(dateStr string, format string) (time.Time, error) {
	return time.Parse(format, dateStr)
}

// ConvertToCurrency converts a float64 to a string representing a currency value.
func ConvertToCurrency(value float64) string {

	return fmt.Sprintf("%.2f", value)
}

// ConvertCurrencyToFloat converts a currency string to a float64.
func ConvertCurrencyToFloat(currencyStr string) (float64, error) {
	return strconv.ParseFloat(currencyStr, 64)
}

func testeData() {
	dataString := "20240608"
	t, err := time.Parse("20060102", dataString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(t)
	dataFormatada := t.Format("02/01/2006")
	fmt.Println(dataFormatada)
}
