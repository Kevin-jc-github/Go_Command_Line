# CSV to JSON Lines Converter

This is a Go command-line application that converts data from a CSV file to a JSON Lines (JSONL) file format. JSON Lines is a convenient format for storing structured data that is easy to process and parse, commonly used for streaming applications, databases, and document stores.

## Overview

The application reads data from a CSV file, where the first row represents the headers, and subsequent rows represent the data entries. The program then converts each row of data into a JSON object and writes these objects into an output file, one JSON object per line.

### Roles of Programs and Data

- **Go Application**: The main program (`csvtojl`) reads the CSV file, converts each row to JSON format, and writes it to an output JSONL file.
- **Input Data (CSV)**: The CSV file (`housesInput.csv`) contains the data to be converted, with headers defining the keys for each JSON object.
- **Output Data (JSONL)**: The output file (`housesOutput.jl`) contains each row from the CSV file in JSON format, separated by newlines.

## Requirements

- [Go](https://golang.org/doc/install) version 1.16 or above.

## Usage Instructions

### Step 1: Clone the Repository

First, clone the repository containing this program to your local machine:

```sh
git clone https://github.com/Kevin-jc-github/Go_Command_Line.git
cd Go_Command_Line
```

### Step 2: Prepare Input Data

Ensure that your CSV file (`housesInput.csv`) is placed in the root of the project directory. The CSV file must have:

- A header row, which defines the column names (used as JSON keys).
- Data rows, which represent the actual data values.

Example CSV (`housesInput.csv`):

```csv
value	income	age	rooms	bedrooms	pop	   hh
452600	8.3252	41	 880	  129	    322	   126
358500	8.3014	21	 7099	  1106	    2401   1138
```

### Step 3: Build the Program

To compile the program, use the following command to build it:

```sh
go build -o csvtojl
```
### Step 4: Create an Executable File for Your OS

For Windows: To create a .exe file that can be executed on Windows, use the following command:

GOOS=windows GOARCH=amd64 go build -o csvtojl.exe

For MacOS: To create an .app file for MacOS, use the following command:

GOOS=darwin GOARCH=amd64 go build -o csvtojl.app

These commands set the target operating system and architecture, allowing you to build platform-specific executable files.

This command creates an executable file named `csvtojl` in the current directory.

### Step 5: Run the Application

To run the program and convert a CSV file to a JSON Lines file, use the following command:

```sh
./csvtojl housesInput.csv housesOutput.jl
```

This command takes two arguments:

1. **Input CSV file**: The path to the CSV file to be read (`housesInput.csv`).
2. **Output JSON Lines file**: The path to the JSONL file to be written (`housesOutput.jl`).

If the command runs successfully, it will output:

```
CSV file has been successfully converted to JSON Lines file
```

The output JSON Lines file (`housesOutput.jl`) will contain each row of data in JSON format, written one per line.

### Example JSON Lines Output (`housesOutput.jl`):

```json
{"age":"41","bedrooms":"129","hh":"126","income":"8.3252","pop":"322","rooms":"880","value":"452600"}
{"age":"21","bedrooms":"1106","hh":"1138","income":"8.3014","pop":"2401","rooms":"7099","value":"358500"}
```

## Testing the Application

1. **Test Input Validity**: Use a CSV file with at least one header and one data row.
2. **Error Handling**: The application will output errors if the input CSV file is not formatted correctly, if the number of columns in data rows does not match the headers, or if there are file read/write issues.

### Example Scenarios:

- **Valid CSV File**: If your CSV file is well-formatted, the program will convert it to JSON Lines format and display a success message.
- **Invalid File Paths**: If an invalid file path is provided, the program will log an error message and exit.

## Conclusion

This application is designed to simplify the process of converting structured CSV data to JSON Lines format, making it easy to handle large datasets for downstream applications. By following the steps above, you should be able to successfully run and test this command-line utility. If you encounter any issues or have questions, feel free to consult Go documentation or ask for help.


