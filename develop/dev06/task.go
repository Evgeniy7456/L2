package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	f int
	d string
	s bool
)

func Cut(str string) string {
	lines := strings.Split(str, "\n") // Разбиение на нексколько строк по разделителю переноса строки
	var result strings.Builder

	if s { // Если используется флаг -s
		for _, line := range lines {
			fields := strings.Split(line, d) // Разбиение строк по разделителю заданному флагом -d
			if len(fields) > 1 {             // Если строка разделилась
				if f <= len(fields) { // Если номер выводимого элемента не превышает длину строки
					result.WriteString(fields[f-1]) // Добавление элемента в результат
					result.WriteString("\n")
				}
			}
		}
	} else {
		for _, line := range lines {
			fields := strings.Split(line, d) // Разбиение строк по разделителю заданному флагом -d
			if len(fields) == 1 {            // Если строка не разделилась
				result.WriteString(fields[0]) // Добавление строки в результат
				result.WriteString("\n")
			} else {
				if f <= len(fields) { // Если номер выводимого элемента не превышает длину строки
					result.WriteString(fields[f-1]) // Добавление элемента в результат
					result.WriteString("\n")
				} else {
					result.WriteString("\n") // Перенос строки
				}
			}
		}
	}

	return strings.Trim(result.String(), "\n") // Возвращение результата без последнего символа переноса строки
}

func main() {
	// Флаги
	flag.IntVar(&f, "f", 0, "\"fields\" - выбрать поля (колонки)")
	flag.StringVar(&d, "d", "\t", "\"delimiter\" - использовать другой разделитель")
	flag.BoolVar(&s, "s", false, "\"separated\" - только строки с разделителем")
	flag.Parse() // Парсинг флагов

	if f == 0 { // Если не используется флаг -f выводится ошибка
		io.WriteString(os.Stderr, "выберите поле (колонку)\n")
		return
	}

	// Чтение данных из Stdout
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	result := Cut(scanner.Text())
	fmt.Println(result)
}
