package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Define a struct that matches the JSON structure. (This is an example; modify according to your JSON structure.)
type Data struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
	// Add more fields as necessary.
}

func main() {
	// Step 1: Fetch the JSON data
	url := "https://api.example.com/data"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	// Decode the JSON data into a slice of Data structs.
	var records []Data
	if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Step 2: Create the CSV file
	file, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the headers (customize according to your struct)
	headers := []string{"Field1", "Field2"} // Add more headers as necessary.
	if err := writer.Write(headers); err != nil {
		fmt.Println("Error writing headers:", err)
		return
	}

	// Step 3: Write data to the CSV file
	for _, record := range records {
		row := []string{record.Field1, record.Field2} // Add more fields as necessary.
		if err := writer.Write(row); err != nil {
			fmt.Println("Error writing record:", err)
			return
		}
	}
	fmt.Println("CSV file successfully created.")
}
