package csvutils

import (
	"encoding/csv"
	"os"
)

// ToArrays reads a CSV file and returns an array
// with each line of the file as array of strings
//
// Parameters
// - filePath string - The path to the CSV file
// Returns
// - [][]string - The lines as array of string
func ToArrays(filePath string) ([][]string, error) {
	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read the CSV records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
