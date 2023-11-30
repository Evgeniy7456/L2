package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Переменные для флагов
var (
	k                   int
	n, r, u, M, b, c, h bool
	in, o               string
)

// Удаление повторяющихся строк при использовании флага -u
func RemoveDuplicates(lines []string) []string {
	m := make(map[string]struct{}, len(lines)) // Структура map, которая используется в качестве set
	for _, item := range lines {               // Перебор всех строк и добавление в map
		m[item] = struct{}{}
	}

	result := make([]string, 0, len(m)) // Создание нового слайса и добавление строк без дубликатов
	for item := range m {
		result = append(result, item)
	}

	return result
}

// Получение из двумерного слайса одномерного
func Slice2ToSlice(lines2 [][]string) []string {
	var result []string
	for _, item := range lines2 { // Перебор всех слайсов и объединение их в одну строку с разделителем пробел
		result = append(result, strings.Join(item, " "))
	}

	return result
}

// Получение из слайса строк одной строки - результата сортировки
func SliceToString(lines []string, data []byte) string {
	dataString := string(data)            // Строка с исходными данными
	result := strings.Join(lines, "\r\n") // Объединение слайса строк в одну строку. Разделители - перенос строки.

	if c { // Если используется флаг -c
		if dataString == result { // Если исходная строка равна отсортированной
			return "Файл отсортирован"
		}
		return "Файл не отсортирован"
	}

	return result // Если флаг -c не используется возвращается отсортированная строка
}

// Сортировка строк без флага -k
func StringLineSort(lines []string) []string {
	if b { // Если используется флаг -b, то у строк отрезаются хвостовые пробелы
		sort.Slice(lines, func(i, j int) bool {
			x := strings.TrimRight(lines[i], " ")
			y := strings.TrimRight(lines[j], " ")
			return x < y
		})
	} else {
		sort.Slice(lines, func(i, j int) bool {
			return lines[i] < lines[j]
		})
	}
	fmt.Println(lines)
	return lines
}

// Сортировка строк по столбцу
func StringColumnSort(lines2 [][]string) []string {
	if b {
		sort.Slice(lines2, func(i, j int) bool {
			x := strings.TrimRight(lines2[i][k-1], " ")
			y := strings.TrimRight(lines2[j][k-1], " ")
			return x < y
		})
	} else {
		sort.Slice(lines2, func(i, j int) bool {
			return lines2[i][k-1] < lines2[j][k-1]
		})
	}

	return Slice2ToSlice(lines2)
}

// Сортировка при использовании флага -n
func IntLineSort(lines []string) []string {
	if b {
		sort.Slice(lines, func(i, j int) bool {
			xTrim := strings.TrimRight(lines[i], " ")
			yTrim := strings.TrimRight(lines[j], " ")
			x, _ := strconv.Atoi(xTrim)
			y, _ := strconv.Atoi(yTrim)
			return x < y
		})
	} else {
		sort.Slice(lines, func(i, j int) bool {
			x, _ := strconv.Atoi(lines[i])
			y, _ := strconv.Atoi(lines[j])
			return x < y
		})
	}

	return lines
}

// Сортировка при использовании флагов -n и -k
func IntColumnSort(lines2 [][]string) []string {
	if b {
		sort.Slice(lines2, func(i, j int) bool {
			xTrim := strings.TrimRight(lines2[i][k-1], " ")
			yTrim := strings.TrimRight(lines2[j][k-1], " ")
			x, _ := strconv.Atoi(xTrim)
			y, _ := strconv.Atoi(yTrim)
			return x < y
		})
	} else {
		sort.Slice(lines2, func(i, j int) bool {
			x, _ := strconv.Atoi(lines2[i][k-1])
			y, _ := strconv.Atoi(lines2[j][k-1])
			return x < y
		})
	}

	return Slice2ToSlice(lines2)
}

// Сортировка при использовании флага -h
func FloatLineSort(lines []string) []string {
	if b {
		sort.Slice(lines, func(i, j int) bool {
			xTrim := strings.TrimRight(lines[i], " ")
			yTrim := strings.TrimRight(lines[j], " ")
			x, _ := strconv.ParseFloat(xTrim, 32)
			y, _ := strconv.ParseFloat(yTrim, 32)
			return x > y
		})
	} else {
		sort.Slice(lines, func(i, j int) bool {
			x, _ := strconv.ParseFloat(lines[i], 32)
			y, _ := strconv.ParseFloat(lines[j], 32)
			return x < y
		})
	}

	return lines
}

// Сортировка при использовании флагов -h и -k
func FloatColumnSort(lines2 [][]string) []string {
	if b {
		sort.Slice(lines2, func(i, j int) bool {
			xTrim := strings.TrimRight(lines2[i][k-1], " ")
			yTrim := strings.TrimRight(lines2[j][k-1], " ")
			x, _ := strconv.ParseFloat(xTrim, 32)
			y, _ := strconv.ParseFloat(yTrim, 32)
			return x > y
		})
	} else {
		sort.Slice(lines2, func(i, j int) bool {
			x, _ := strconv.ParseFloat(lines2[i][k-1], 32)
			y, _ := strconv.ParseFloat(lines2[j][k-1], 32)
			return x < y
		})
	}

	return Slice2ToSlice(lines2)
}

// Структура map с названиями месяцев и их номерами
var month = map[string]int{
	"JAN": 1,
	"FEB": 2,
	"MAR": 3,
	"APR": 4,
	"MAY": 5,
	"JUN": 6,
	"JUL": 7,
	"AUG": 8,
	"SEP": 9,
	"OCT": 10,
	"NOB": 11,
	"DEC": 12,
}

// Сортировка при использовании флага -M
func MonthLineSort(lines []string) []string {
	sort.Slice(lines, func(i, j int) bool {
		x := month[strings.ToUpper(lines[i][:3])]
		y := month[strings.ToUpper(lines[j][:3])]
		return x < y
	})

	return lines
}

// Сортировка при использовании флагов -M и -k
func MonthColumnSort(lines2 [][]string) []string {
	k := k

	sort.Slice(lines2, func(i, j int) bool {
		x := month[strings.ToUpper(lines2[i][k-1][:3])]
		y := month[strings.ToUpper(lines2[j][k-1][:3])]
		return x < y
	})

	return Slice2ToSlice(lines2)
}

// Функция вызывающая сортировку без флага -k
func LineSort(lines []string, data []byte) string {
	var sortSlice []string
	switch { // Вызов сортировки в зависимости от используемых флагов
	case n:
		sortSlice = IntLineSort(lines)
	case h:
		sortSlice = FloatLineSort(lines)
	case M:
		sortSlice = MonthLineSort(lines)
	default:
		sortSlice = StringLineSort(lines)
	}

	if r { // При использовании флага -r элементы в слайсе переворачиваются
		slices.Reverse(sortSlice)
	}

	return SliceToString(sortSlice, data) // Получение строки из слайса
}

// Функция вызывающая сортировку с флагом -k
func ColumnSort(lines []string, data []byte) string {
	var lines2 [][]string
	for i := 0; i < len(lines); i++ { // Перебор строк и разделение их по пробелу
		lines2 = append(lines2, strings.Split(lines[i], " "))
	}

	var sortSlice []string
	switch {
	case n:
		sortSlice = IntColumnSort(lines2)
	case h:
		sortSlice = FloatColumnSort(lines2)
	case M:
		sortSlice = MonthColumnSort(lines2)
	default:
		sortSlice = StringColumnSort(lines2)
	}

	if r {
		slices.Reverse(sortSlice)
	}

	return SliceToString(sortSlice, data)
}

// Функция сортировки
func Sort(data []byte) string {
	lines := strings.Split(string(data), "\n") // Получение из строки слайса. Разделители - перенос строки.

	if u { // При использовании флага -u удаляются повторяющиеся строки
		lines = RemoveDuplicates(lines)
	}

	var result string
	if k > 0 { // В зависимости от использования флага -k вызваются разные функции сортировок
		result = ColumnSort(lines, data)
	} else {
		result = LineSort(lines, data)
	}

	return result
}

func main() {
	// Флаги
	flag.IntVar(&k, "k", 0, "указание колонки для сортировки")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&M, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&b, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&c, "c", false, "проверять отсортированы ли данные")
	flag.BoolVar(&h, "h", false, "сортировать по числовому значению с учётом суффиксов")
	flag.StringVar(&in, "in", "", "файл для сортировки")
	flag.StringVar(&o, "o", "", "файл для сохранения результата")
	flag.Parse() // Парсинг флагов

	if o != "" && c { // Вывод ошибки в случае одновременного использования флаго -o и -c
		io.WriteString(os.Stderr, "варианты несовместимы")
		return
	}

	if in == "" { // Вывод ошибки, если не используется флаг -in
		io.WriteString(os.Stderr, "укажите путь к файлу\n")
		return
	}

	data, err := os.ReadFile(in) // Чтение данных из файла
	if err != nil {
		io.WriteString(os.Stderr, "указанного файла не существует\n")
		return
	}

	result := Sort(data) // Сортировка данных

	if o == "" { // Если используется флаг -o происходит запись результата в файл
		io.WriteString(os.Stderr, result+"\n")
	} else {
		file, err := os.Create(o) // Создание файла для сохранения результата сортировки
		if err != nil {
			io.WriteString(os.Stderr, fmt.Sprintf("не удалось создать файл: %s\n", o))
			return
		}
		defer file.Close() // Закрываем файл

		_, err = file.WriteString(result) // Запись данных в файл
		if err != nil {
			io.WriteString(os.Stderr, "не удалось записать данные в файл\n")
			return
		}
	}
}
