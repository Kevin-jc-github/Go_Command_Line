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
	// This program expects exactly two command-line arguments: the input CSV file path and the output JSON Lines file path.
	// If the arguments are not provided correctly, print usage instructions and exit
	if len(os.Args) != 3 {
		fmt.Println("Usage: csvtojl <input_csv_file> <output_jsonl_file>")
		os.Exit(1)
	}

	// Get the input CSV file path and output JSON Lines file path from command-line arguments.
	// The first argument after the program name is the input CSV file path.
	// The second argument is the output JSON Lines file path.
	inputFilePath := os.Args[1]
	outputFilePath := os.Args[2]

	// Open the input CSV file for reading.
	// The CSV file is opened using `os.Open()`, which returns a file descriptor and an error.
	// If the file cannot be opened (e.g., it doesn't exist), the program logs a fatal error and exits.
	csvFile, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Unable to open input file: %v\n", err)
	}
	defer csvFile.Close()

	// Create a CSV reader to read the input file
	reader := csv.NewReader(csvFile)

	// Read all records from the CSV file.
	// `reader.ReadAll()` reads the entire content of the CSV into a slice of string slices.
	// If an error occurs during reading, the program logs a fatal error and exits.
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV file: %v\n", err)
	}

	// Check if the CSV file contains enough data (at least headers and one row)
	if len(records) < 2 {
		log.Fatalln("Not enough data in the CSV file")
	}

	// Use the first row as the header names
	headers := records[0]

	// Open the output JSON Lines file for writing
	jsonlFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("Unable to create output file: %v\n", err)
	}
	defer jsonlFile.Close()

	// Iterate through all records and convert them to JSON objects
	for _, record := range records[1:] {
		// Ensure the number of columns matches the number of headers
		if len(record) != len(headers) {
			fmt.Println("Number of columns in the record does not match headers")
			continue
		}

		// Create a map to store field names and their corresponding values
		jsonData := make(map[string]string)
		for i, header := range headers {
			jsonData[header] = record[i]
		}

		// Convert the map to a JSON string.
		// `json.Marshal()` is used to convert the map to a JSON byte slice.
		// If an error occurs during conversion, print a warning and skip the record.
		jsonBytes, err := json.Marshal(jsonData)
		if err != nil {
			fmt.Printf("Error converting to JSON: %v\n", err)
			continue
		}

		// Write the JSON string to the JSON Lines file
		_, err = jsonlFile.WriteString(string(jsonBytes) + "\n")
		if err != nil {
			log.Fatalf("Error writing to JSON Lines file: %v\n", err)
		}
	}

	// Print a success message to indicate the conversion is complete
	fmt.Println("CSV file has been successfully converted to JSON Lines file")
}
