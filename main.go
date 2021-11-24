package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func findIPO(ticker string) {

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
		// return stock data matching ticker
		// Using ToUpper to capitalize user input
		if record[0] == strings.ToUpper(ticker) {
			fmt.Println("Stock:", record[0], "IPO'd on", record[6], record[7], record[5], "at $", record[12])
			break
		}
	}
}

func main() {
	var stock string
	fmt.Println("Enter a Stock Symbol: ")
	fmt.Scanln(&stock)
	findIPO(stock)
}
