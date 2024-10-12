package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"log"
)

func main() {
	// Check if the command-line arguments are valid
	// Since we do not need any arguments, we only expect the program name
	if len(os.Args) != 1 {
		fmt.Println("Usage: csvtojl (no arguments needed)")
		os.Exit(1)
	}

	// Define the input and output file paths
	// The input CSV file and output JSON Lines file are both located in ~/Desktop/MSDS431
	inputFilePath := "~/Desktop/MSDS431/housesInput.csv"
	outputFilePath := "~/Desktop/MSDS431/housesOutput.jl"

	// Expand the tilde (~) in the file paths to the full home directory path
	inputFilePath = os.ExpandEnv(inputFilePath)
	outputFilePath = os.ExpandEnv(outputFilePath)

	// Open the input CSV file for reading
	csvFile, err := os.Open(inputFilePath)
	if err != nil {
		log.Fatalf("Unable to open input file: %v\n", err)
	}
	defer csvFile.Close()

	// Create a CSV reader to read data from the input CSV file
	reader := csv.NewReader(csvFile)

	// Read all data from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV file: %v\n", err)
	}

	// Ensure there are enough rows in the CSV file (header + data)
	if len(records) < 2 {
		log.Fatalln("Not enough data in the CSV file")
	}

	// Use the first row of the CSV file as the field names (headers)
	headers := records[0]

	// Open the output JSON Lines file for writing
	jsonlFile, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("Unable to create output file: %v\n", err)
	}
	defer jsonlFile.Close()

	// Iterate over each row of data (excluding the header) and convert it to a JSON object
	for _, record := range records[1:] {
		// Ensure that the number of columns matches the number of headers
		if len(record) != len(headers) {
			fmt.Println("Data does not match the number of headers")
			continue
		}

		// Create a map to store field names and their corresponding values
		jsonData := make(map[string]string)
		for i, header := range headers {
			jsonData[header] = record[i]
		}

		// Convert the map to a JSON string
		jsonBytes, err := json.Marshal(jsonData)
		if err != nil {
			fmt.Printf("Error converting to JSON: %v\n", err)
			continue
		}

		// Write the JSON string to the output JSON Lines file
		_, err = jsonlFile.WriteString(string(jsonBytes) + "\n")
		if err != nil {
			log.Fatalf("Error writing to JSON Lines file: %v\n", err)
		}
	}

	// Print success message
	fmt.Println("CSV file has been successfully converted to JSON Lines file")
}"`
    }
  ]
}