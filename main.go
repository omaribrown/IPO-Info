package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func findIPO(ticker string) string {

	// Declare a variable to store negative value
	result := "Nothing Found."

	// Open the file
	csvfile, err := os.Open("IPODataFull.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		// break at end of file
		if err == io.EOF {
			break
		}
		// log an error
		if err != nil {
			log.Fatal(err)
		}
		// Using ToUpper to capitalize user input
		if record[0] == strings.ToUpper(ticker) {

			// Assign returned date values to a string
			ipoDay := record[6]
			ipoMonth := record[7]
			ipoYear := record[5]
			ipoPrice := record[11]
			ipoDate := ipoDay + "/" + ipoMonth + "/" + ipoYear + " at $" + ipoPrice
			result := ipoDate

			// Return the result
			fmt.Println(result)
			return result
		}
	}
	return result
}

func main() {
	var stock string
	fmt.Println("Enter a Stock Symbol: ")
	fmt.Scanln(&stock)
	// Run function passing any stock symbol
	findIPO(stock)
}
