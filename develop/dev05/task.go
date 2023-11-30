package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	A, B, C       int
	c, i, v, F, n bool
)

// Поиск слова в строке
func TextSearch(lines []string, text string) map[int]struct{} {
	result := make(map[int]struct{}) // Переменная map с результатом поиска

	if F { // Если используется флаг -F
		for index, line := range lines {
			if line == text { // Если строка полностью совпадает с искомым текстом, то индекс данной строки добавляется в map
				result[index] = struct{}{}
			}
		}
	} else {
		if i { // Если используется флаг -i
			text = strings.ToLower(text) // Искомый текст приводится к нижнему регистру
			for index, line := range lines {
				lines[index] = strings.ToLower(line) // Строки, в которых производится поиск, приводятся к нижнему регистру
			}
		}
		for index, line := range lines {
			if strings.Contains(line, text) { // Если искомая подстрока содержится в строке, то индекс данной строки добавляется в map
				result[index] = struct{}{}
			}
		}
	}

	return result
}

// Функция поиска
func Search(data []byte) {
	lines := strings.Split(string(data), "\n") // Разбиение данных на строки. Разделитель - перенос строки.

	found := TextSearch(lines, flag.Arg(0)) // Поиск подстроки в строке

	if c { // Если используется флаг -c
		if v { // Если используется флаг -v
			io.WriteString(os.Stdout, fmt.Sprint(len(lines)-len(found))) // Вывод количества строк, которые не совпали
			return
		} else {
			io.WriteString(os.Stdout, fmt.Sprint(len(found))) // Вывод количества найденных совпадений
			return
		}
	}

	if v { // Если используется флаг -v
		if n { // Если используется флаг -n
			for index, item := range lines {
				if _, ok := found[index]; !ok { // Если в строке есть совпадение, то она не выводится
					io.WriteString(os.Stdout, fmt.Sprintf("%d:%s\n", index+1, item)) // Вывод строк с номерами
				}
			}
			return
		} else {
			for index, item := range lines {
				if _, ok := found[index]; !ok {
					io.WriteString(os.Stdout, item+"\n") // Вывод строк без номеров
				}
			}
			return
		}
	}

	if C > 0 { // Если используется флаг -C
		if C > A { // Если значение флага -C больше значения флага -A
			A = C // Значение флага -C присваивается флагу -A
		}
		if C > B { // Если значение флага -C больше значения флага -B
			B = C // Значение флага -C присваивается флагу -B
		}
	}

	if A > 0 || B > 0 { // Если используется флаг -A или флаг -B
		linesCount := len(lines)  // Общее количество строк
		for item := range found { // Перебор строк с совпадениями
			for i := item - B; i <= item+A; i++ { // Перебор индексов от индекс строки с совпадением минус значение флага -B до плюс значение флага -A
				if i >= 0 && i < linesCount {
					found[i] = struct{}{}
				}
			}
		}
	}

	result := make([]int, 0, len(found))
	for item := range found { // Индексы найденных строк из map добавляются в слайс
		result = append(result, item)
	}

	slices.Sort(result) // Сортировка индексов

	if n { // Если используется флаг -n
		for _, item := range result {
			io.WriteString(os.Stdout, fmt.Sprintf("%d:%s\n", item+1, lines[item])) // Вывод строк с номерами
		}
	} else {
		for _, item := range result {
			io.WriteString(os.Stdout, (lines[item])+"\n") // Вывод строк без номеров
		}
	}
}

func main() {
	// Флаги
	flag.IntVar(&A, "A", 0, "\"after\" печатать +N строк после совпадения")
	flag.IntVar(&B, "B", 0, "\"before\" печатать +N строк до совпадения")
	flag.IntVar(&C, "C", 0, "\"context\" (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&c, "c", false, "\"count\" (количество строк)")
	flag.BoolVar(&i, "i", false, "\"ignore-case\" (игнорировать регистр)")
	flag.BoolVar(&v, "v", false, "\"invert\" (вместо совпадения, исключать)")
	flag.BoolVar(&F, "F", false, "\"fixed\", точное совпадение со строкой, не паттерн")
	flag.BoolVar(&n, "n", false, "\"line num\", печатать номер строки")
	flag.Parse()         // Парсинг флагов
	input := flag.Arg(1) // Путь к файлу

	if input == "" {
		io.WriteString(os.Stderr, "укажите путь к файлу\n")
		return
	}

	data, err := os.ReadFile(input)
	if err != nil {
		io.WriteString(os.Stderr, "указанного файла не существует")
		return
	}

	Search(data)
}
