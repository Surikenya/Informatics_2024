package laba8

import (
	"fmt"
	"os"
)

func CreateFile(fileName string) error {
	_, err := os.Stat(fileName)
	if err == nil {
		return fmt.Errorf("файл с таким названием уже существует")
	}

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("не удалось создать файл ошибка: %w", err)
	}
	defer file.Close()

	return nil
}
