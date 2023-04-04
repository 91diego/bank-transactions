package utils

import (
	"encoding/csv"
	"mime/multipart"
)

func ReadCSV(file multipart.File) ([][]string, error) {

	r := csv.NewReader(file)
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}
