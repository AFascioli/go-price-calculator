package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	// will execute when parent/outer/enclosing function finishes
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		// file.Close()
		return nil, errors.New("failed to read a line in the file")
	}

	// file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	// could write type any
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("could not create file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		// file.Close()
		return errors.New("failed to convert data to json")
	}

	// file.Close()
	return nil
}
