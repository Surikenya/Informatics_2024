package laba8

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func PoiskTexta(fileName string, searchText string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("не удается открыть файл")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	naiden := false
	for fileScanner.Scan() {
		if strings.Contains(fileScanner.Text(), searchText) {
			fmt.Print("Заданный текст найден на строке \n")
			naiden = true
		}
	}

	if !naiden {
		fmt.Println("Текст не найден")
	}

	return err
}
