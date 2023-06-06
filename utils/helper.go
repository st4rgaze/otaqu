package utils

import (
	"strconv"
	"strings"
)

// convert string price from scrape to uint
func ConvertPriceToUint(priceStr string) uint {
	priceStr = strings.ReplaceAll(priceStr, "IDR ", "")
	priceStr = strings.ReplaceAll(priceStr, ".", "")

	price, err := strconv.ParseUint(priceStr, 10, 32)
	if err != nil {
		return 0
	}

	return uint(price)
}

// convert uint price from table to formatted string
func FormatPrice(price uint) string {
	currencySymbol := "IDR"
	priceStr := strconv.Itoa(int(price))

	formattedPrice := ""
	for i := len(priceStr) - 1; i >= 0; i-- {
		formattedPrice = string(priceStr[i]) + formattedPrice
		if (len(priceStr)-i)%3 == 0 && i != 0 {
			formattedPrice = "." + formattedPrice
		}
	}

	formattedPrice = currencySymbol + " " + formattedPrice

	return formattedPrice
}
