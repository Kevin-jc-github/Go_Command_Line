package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	// Check if command-line arguments are valid.
	// Expect exactly two arguments: the input CSV file path and the output JSON Lines file path.
	if len(os.Args) != 3 {
		fmt.Println("Usage: csvtojl <input_csv_file> <output_jsonl_file>")
		os.Exit(1)
	}

	// Get the input CSV file path and output JSON Lines file path from command-line arguments.
	inputFilePath := os.Args[1]
	outputFilePath := os.Args[2]

	// Open the input CSV file for reading.
	csvFile, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Unable to open input file: %v\n", err)
	}
	defer csvFile.Close()

	// Use `csv.NewReader()` to create a reader for reading the CSV file.
	// Create a CSV reader to read the input file.
	reader := csv.NewReader(csvFile)

	// Use `reader.ReadAll()` to read all records from the CSV file.
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV file: %v\n", err)
	}

	// Check if the CSV file contains enough data (at least headers and one row).
	if len(records) < 2 {
		log.Fatalln("Not enough data in the CSV file")
	}

	// Use the first row as the header names.
	headers := records[0]

	// Open the output JSON Lines file for writing.
	jsonlFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("Unable to create output file: %v\n", err)
	}
	defer jsonlFile.Close()

	// Iterate through all records and convert them to JSON objects.
	for _, record := range records[1:] {
		// Ensure the number of columns matches the number of headers.
		if len(record) != len(headers) {
			fmt.Println("Number of columns in the record does not match headers")
			continue
		}

		// Create a map to store field names and their corresponding values.
		jsonData := make(map[string]string)
		for i, header := range headers {
			jsonData[header] = record[i]
		}

		// Convert the map to a JSON string.
		jsonBytes, err := json.Marshal(jsonData) // Use `json.Marshal()` to convert the map to a JSON string.
		if err != nil {
			fmt.Printf("Error converting to JSON: %v\n", err)
			continue
		}

		// Write the JSON string to the JSON Lines file.
		_, err = jsonlFile.WriteString(string(jsonBytes) + "\n")
		if err != nil {
			log.Fatalf("Error writing to JSON Lines file: %v\n", err)
		}
	}

	// Success message to indicate the conversion is complete.
	fmt.Println("CSV file has been successfully converted to JSON Lines file")
}
