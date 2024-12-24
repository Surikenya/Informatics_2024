package laba8

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"isuct.ru/informatics2022/laba4ab"
)

func recordingData(fileName string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		newLine, _ := reader.ReadString('\n')
		newLine = strings.TrimSpace(newLine)
		if newLine == "Все" {
			fmt.Println("Данные успешно записаны")
			break
		}

		err := WriteFile(fileName, newLine)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func ReadyLab8() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите название файла:")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)
	err := CreateFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Введите текст для записи в файл (Если закончили - напишите 'Все')")
	recordingData(fileName)

	file, err := ReadForLaba4(fileName)
	if err != nil {
		panic("Произошла ошибка при чтении файла")
	}
	fmt.Println("В файле: \n", file)

	fmt.Println("Что искать в файле:")
	searchText, _ := reader.ReadString('\n')
	searchText = strings.TrimSpace(searchText)
	err = FindText(fileName, searchText)
	if err != nil {
		fmt.Println(err)
		return
	}

	variables, err := ReadFile(fileName)
	if err != nil {
		panic("Произошла ошибка при чтении файла")
	}

	fmt.Print("Задача А\n")
	fmt.Print(laba4ab.CalculateA(variables[1], variables[2], variables[3], variables[0]), "\n")
	fmt.Print("Задача В\n")
	fmt.Print(laba4ab.CalculateB(variables[4:], variables[0]), "\n")
}
