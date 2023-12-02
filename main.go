package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Проверяем аргументы командной строки
	var inputFile, outputFile string

	if len(os.Args) >= 2 {
		inputFile = os.Args[1]
	} else {
		inputFile = "in.txt"
	}

	if len(os.Args) >= 3 {
		outputFile = os.Args[2]
	} else {
		outputFile = "out.txt"
	}

	// Читаем содержимое файла с выражениями
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Ошибка при чтении исходного файла: %v\n", err)
		return
	}

	// Открываем файл для записи результатов
	output, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Printf("Ошибка при открытии файла для записи результатов: %v\n", err)
		return
	}
	defer output.Close()

	// Создаем буфер для записи результатов
	writer := bufio.NewWriter(output)

	// Разделяем содержимое файла на строки
	lines := strings.Split(string(content), "\n")

	// Создаем регулярное выражение для поиска математических выражений в строках
	regex := regexp.MustCompile(`(\d+)([+\-*\/])(\d+)=\?`)

	// Обработка каждой строки
	for _, line := range lines {
		// Проверяем, является ли строка математическим выражением
		if regex.MatchString(line) {
			// Извлекаем операнды и операторы из строки
			matches := regex.FindStringSubmatch(line)
			operand1, _ := strconv.Atoi(matches[1])
			operator := matches[2]
			operand2, _ := strconv.Atoi(matches[3])

			// Вычисляем результат
			result := 0
			switch operator {
			case "+":
				result = operand1 + operand2
			case "-":
				result = operand1 - operand2
			case "*":
				result = operand1 * operand2
			case "/":
				result = operand1 / operand2
			}

			// Формируем строку с результатом
			resultLine := fmt.Sprintf("%d%s%d=%d\n", operand1, operator, operand2, result)

			// Записываем строку в буфер
			writer.WriteString(resultLine)
		}
	}

	// Очищаем буфер и записываем результаты в файл
	writer.Flush()

	fmt.Println("Результаты записываем в ", outputFile)
}
