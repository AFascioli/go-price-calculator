package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("failed to read a line in the file")
	}

	file.Close()
	return lines, nil
}

// could write type any
func WriteJSON(path string, data interface{}) error {
	file, err := os.Create(path)

	if err != nil {
		return errors.New("could not create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to convert data to json")
	}

	file.Close()
	return nil
}
