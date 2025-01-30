package utils

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsvFile(filePath string) ([][]string, map[string]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read CSV header: %w", err)
	}

	columnIndex := buildColumnIndex(header)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read CSV file: %w", err)
	}

	return records, columnIndex, nil
}

func buildColumnIndex(header []string) map[string]int {
	columnIndex := make(map[string]int)
	for i, columnName := range header {
		columnIndex[columnName] = i
	}
	return columnIndex
}
