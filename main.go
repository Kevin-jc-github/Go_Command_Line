package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	// Start CPU profiling and defer stop profiling until program ends.
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatalf("Unable to create CPU profile: %v\n", err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatalf("Unable to start CPU profile: %v\n", err)
	}
	defer pprof.StopCPUProfile()

	// Record the start time for performance monitoring.
	start := time.Now()

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

	// Create a CSV reader to read the input file.
	reader := csv.NewReader(csvFile)

	// Read all records from the CSV file.
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

	// Monitor memory usage.
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	log.Printf("Initial memory usage: Alloc = %v KB\n", memStats.Alloc/1024)

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
		jsonBytes, err := json.Marshal(jsonData)
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

	// Record and print the total elapsed time for performance monitoring.
	elapsed := time.Since(start)
	log.Printf("CSV to JSON Lines conversion completed in %s\n", elapsed)

	// Monitor final memory usage.
	runtime.ReadMemStats(&memStats)
	log.Printf("Final memory usage: Alloc = %v KB\n", memStats.Alloc/1024)

	// Print a success message to indicate the conversion is complete.
	fmt.Println("CSV file has been successfully converted to JSON Lines file")
}

/*
Explanation of the code:

1. **Imports and Package Declaration**:
   - Import necessary packages: `encoding/csv` for CSV file handling, `encoding/json` for JSON encoding, `fmt` for formatted I/O, `os` for file handling, `log` for logging errors, `runtime/pprof` for profiling, and `time` for monitoring execution duration.

2. **Start CPU Profiling**:
   - Use `pprof.StartCPUProfile()` to start CPU profiling and generate a profile file (`cpu.prof`). Profiling is deferred to stop at the end of the program.

3. **Record Start Time**:
   - Use `time.Now()` to record the start time for performance monitoring.

4. **Command-Line Argument Check**:
   - The program checks if there are exactly two command-line arguments (input and output file paths). If not, it prints usage instructions and exits.

5. **Define File Paths**:
   - Get the input and output file paths from the command-line arguments.

6. **Open Input CSV File**:
   - Open the input CSV file for reading using `os.Open()`. If the file cannot be opened, log a fatal error and exit.
   - Use `defer csvFile.Close()` to ensure the file is properly closed at the end.

7. **Create CSV Reader and Read Data**:
   - Use `csv.NewReader()` to create a reader for reading the CSV file.
   - Use `reader.ReadAll()` to read all records from the CSV file.
   - If an error occurs during reading, log a fatal error and exit.

8. **Check CSV Data Validity**:
   - Ensure that the CSV file contains at least a header row and one data row. If not, log a fatal error and exit.

9. **Extract Headers**:
   - Use the first row of the CSV file as the header names, which will be used as keys in the JSON objects.

10. **Open Output JSON Lines File**:
    - Create the output JSON Lines file using `os.Create()`. If the file cannot be created, log a fatal error and exit.
    - Use `defer jsonlFile.Close()` to ensure the file is properly closed at the end.

11. **Monitor Initial Memory Usage**:
    - Use `runtime.ReadMemStats()` to capture memory statistics before processing records.

12. **Convert Records to JSON and Write to Output File**:
    - Iterate through each record (excluding the header) and convert it to a JSON object.
    - Use a map to create key-value pairs where the key is the header and the value is the corresponding column value.
    - Use `json.Marshal()` to convert the map to a JSON string.
    - Write the JSON string to the output file in JSON Lines format (one JSON object per line).

13. **Record and Print Elapsed Time**:
    - Use `time.Since()` to measure the elapsed time since the start of the program and print it for performance monitoring.

14. **Monitor Final Memory Usage**:
    - Use `runtime.ReadMemStats()` to capture memory statistics after processing records.

15. **Print Success Message**:
    - Print a success message to indicate that the conversion process is complete.
*/
