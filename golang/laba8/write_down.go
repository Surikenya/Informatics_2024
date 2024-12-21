package laba8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ReadFile1(fileName string) (string, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return "", fmt.Errorf("не удается прочитать файл: %w", err)
	}
	return string(file), nil
}

func ReadFile2(fileName string) ([]float64, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("не удается открыть файл")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	var variables []float64

	for fileScanner.Scan() {
		number, err := strconv.ParseFloat(fileScanner.Text(), 64)
		if err != nil {
			return nil, fmt.Errorf("не удается перевести число")
		}
		variables = append(variables, number)
	}

	if len(variables) < 5 {
		return variables, fmt.Errorf("недостаточно значений")
	}

	return variables, fileScanner.Err()
}
