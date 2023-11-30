package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var path, site string // Переменные под флаги

	flag.StringVar(&path, "p", "", "download path") // Флаг -p, который указывает путь сохранения скачанного сайта
	flag.Parse()                                    // Считывание флагов

	if len(flag.Args()) == 0 { // Если длина аргументов равна 0
		io.WriteString(os.Stderr, "enter site adress\n") // Ошибка означающая, что адрес сайта не указан
	} else {
		site = flag.Args()[0] // Получение адреса сайта
	}

	re, _ := regexp.Compile(`\w+\.\w+`) // Создание регулярного выражения для получения названия сайта без протокола
	root := re.FindString(site)         // Получение названия сайта

	if path != "" { // Если пользователь указал путь
		os.MkdirAll(path, 0777) // Создание директорий на случай когда указанных директорий не существует
		os.Chdir(path)          // Перемещение в указанную директорию
	}

	os.Mkdir(root, 0777)     // Создание директории с названием сайта
	os.Chdir(root)           // Перемещение в эту директорию
	rootDir, _ := os.Getwd() // Получение пути к текущей директории как к корневой

	var links []string                    // Слайс ссылок
	setLinks := make(map[string]struct{}) // Структура map для проверки на наличие ссылки в links

	site = strings.Trim(site, "/") // Убираем символ "/" в конце ссылки
	links = append(links, site)    // Добавление в links ссылки на сайт
	setLinks[site] = struct{}{}    // Добавление ссылки во множество ссылок

	i := 0               // Переменная для итерации по слайсу ссылок
	for i < len(links) { // Если не все ссылки были просмотрены
		resp, err := http.Get(links[i]) // Использование HTTP-метода GET
		if err != nil {                 // Если получили ошибку
			io.WriteString(os.Stderr, err.Error()+"\n") // Вывод ошибки
			i++                                         // Увеличение значения переменной на 1
			continue                                    // Переход к следующей итерации цикла
		}
		defer resp.Body.Close() // Закрытие запрос

		re, _ := regexp.Compile(`\w+\..+`)  // Получение адреса без протокола
		sitePath := re.FindString(links[i]) // Поиск

		os.Chdir(rootDir)                     // Переход в корневую директорию
		split := strings.Split(sitePath, "/") // Разделение адреса по символу "/"
		var dir string                        // Переменная для создания директории
		if len(split) > 2 {                   // Если адрес разделился на 3 части и более
			dir = strings.Join(split[1:len(split)-1], "/") // Получение строки из слайса с разделителем "/"
			os.MkdirAll(dir, 0777)                         // Создание директорий
			os.Chdir(dir)                                  // Переход в созданную директорию
		}

		var file *os.File    // Переменная для взаимодействия файлом
		if len(split) == 1 { // Если длина разделенного адреса равна 1
			file, _ = os.Create("index.html") // Создание файла "index.html"
		} else {
			file, _ = os.Create(fmt.Sprintf("%s.html", split[len(split)-1])) // Создание файла с названием последнего элмента в адресе
		}
		defer file.Close() // Закрытие файла

		result, _ := io.ReadAll(resp.Body) // Получение результата запроса
		file.Write(result)                 // Запись результата в файл

		re, _ = regexp.Compile(`href="([^"]*)"`) // Поиск ссылок в файле
		hrefs := re.FindAllSubmatch(result, -1)

		for _, item := range hrefs { // Перебор полученных ссылок
			var href string             // Переменная для итоговой ссылки
			findHref := string(item[1]) // Текущая ссылка

			if ok := strings.Contains(findHref, site); ok { // Если в ссылке содержится название сайта
				href = findHref // Итоговая ссылка равна текущей
			}

			if string(findHref[0]) == "/" { // Если ссылка начинается с символа "/"
				href = site + strings.TrimRight(findHref, "/") // Итоговая ссылка равна адрес сайта плюс относительная ссылка
			}

			if _, ok := setLinks[href]; href != "" && !ok { // Если ссылка не пустая строка и ее нет в map setLinks
				setLinks[href] = struct{}{} // Добавление ссылки в map setLinks
				links = append(links, href) // Добавление ссылки в слайс links
			}
		}

		i++ // Увеличение значения переменной на 1
	}
}
