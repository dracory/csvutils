package csvutils

import "strings"

// ToMaps reads a CSV file and returns an array
// with each line of the file as a key-value map
//
// Parameters
// - filePath string - The path to the CSV file
// Returns
// - []map[string]string - The lines each one as a map
func ToMaps(filePath string, replacements map[string]string) (returnMap []map[string]string, err error) {
	records, err := ToArrays(filePath)

	if err != nil {
		return returnMap, err
	}

	if len(records) < 2 {
		return returnMap, err
	}

	header := records[0]

	for i := 0; i < len(header); i++ {
		header[i] = strings.TrimSpace(header[i])

		for k, v := range replacements {
			if header[i] == k {
				header[i] = v
			}
		}
	}

	records = records[1:]

	for _, record := range records {
		line := map[string]string{}
		for i := 0; i < len(record); i++ {
			line[header[i]] = record[i]
		}
		returnMap = append(returnMap, line)
	}

	return returnMap, nil
}
