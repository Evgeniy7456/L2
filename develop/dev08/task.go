package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// Смена директории
func cd(command []string) {
	var path string // Переменная для пути на который нужно поменять

	switch len(command) {
	case 1: // Если длина команды равна 1, т.е. кроме команды echo нет аргументов
		path = "" // Путь равен пустой строке
	case 2: // Если длина пути равна 2
		path = command[1] // Путь равен второму элементу слайса command
	default: // Если длина пути больше 2
		io.WriteString(os.Stderr, "too many arguments\n") // Вывод ошибки: слишком много аргументов
		return                                            // Завершение работы функции
	}

	switch path {
	case "", "~", "~/": // Если путь равен пустой строке, символу "~" или "~/"
		path, _ = os.UserHomeDir() // Присвоение пути адреса домашней директории
	case "-": // Если путь равен символу "-"
		if os.Args[1] == "" { // Если аргумент равен пустой строке
			io.WriteString(os.Stderr, "OLDPWD not set\n") // Вывод ошибки старый путь не задан
			return                                        // Завершение работы функции
		} else {
			path = os.Args[1] // Присвоение пути старого адреса, который был до смены директории
		}
	}

	os.Args[1], _ = syscall.Getwd() // Получение текущего пути и присвоение аргументу

	err := syscall.Chdir(path) // Смена директории
	if err != nil {
		io.WriteString(os.Stderr, err.Error()+"\n")
	}
}

// Показать путь до текущего каталога
func pwd() {
	path, err := syscall.Getwd() // Получение текущего пути
	if err != nil {
		io.WriteString(os.Stderr, err.Error()+"\n")
		return
	}
	fmt.Println(path) // Вывод пути
}

// Вывод аргумента в Stdout
func echo(command []string) {
	if len(command) == 1 { // Если длина слайса command равна 1, т.е. только команда echo
		fmt.Println("") // Вывод пустой строки
	} else {
		fmt.Println(strings.Join(command[1:], " ")) // Вывод элементов слайса command начиная со второго
	}
}

// Завершение работы процесса
func kill(command []string) {
	for _, item := range command[1:] { // Перебор элементов слайса command начиная со второго
		pid, err := strconv.Atoi(item) // Преобразование string в int
		if err != nil {
			io.WriteString(os.Stderr, "argument must be a number\n")
			return
		}
		err = syscall.Kill(pid, syscall.SIGKILL) // Завершение работы процесса
		if err != nil {
			io.WriteString(os.Stderr, err.Error()+"\n")
		}
	}
}

// Вывод общей информации по запущенный процессам
func ps() {
	file, err := os.Open("/proc") // Открытие директории /proc
	if err != nil {
		io.WriteString(os.Stderr, err.Error()+"\n")
		return
	}
	defer file.Close() // Закрытие файла

	fnames, err := file.Readdirnames(-1) // Получение слайса из названий папок
	if err != nil {
		io.WriteString(os.Stderr, err.Error()+"\n")
		return
	}

	pids := make([]int, 0)         // Инициализация слайса int
	for _, fname := range fnames { // Перебор названий папок
		pid, err := strconv.Atoi(fname) // Преобразование string в int
		if err != nil {                 // Если не удалось преобразовать
			continue // Переход к следующей итерации цикла
		}
		pids = append(pids, int(pid)) // Добавление int в слайс
	}

	for _, pid := range pids { // Перебор pid процессов
		path, err := os.Readlink(fmt.Sprintf("/proc/%d/exe", pid)) // Получение пути запущенного процесса
		if err != nil {                                            // Если не удалось получить путь из-за отсутсвия прав доступа
			continue // Переход к следующей итерации цикла
		}
		fmt.Println(pid, path) // Вывод результата в формате: pid путь
	}
}

// Завершение работы утилиты
func exit() {
	os.Exit(0)
}

// Запуск процесса с помощью системных вызовов
func forkExec(file string, command []string, procAttr *syscall.ProcAttr) {
	path, err := exec.LookPath(file) // Поиск пути для запуска процесса по его названию
	if err != nil {
		io.WriteString(os.Stderr, fmt.Sprintf("%s: command not found\n", file))
		return
	}
	syscall.ForkExec(path, command, procAttr) // Запуск процесса
	time.Sleep(1 * time.Millisecond)          // Пауза для ожидания завершения работы процесса, чтобы результат был выведен до вывода символов ">> "
}

// Вызов команды
func execute(command []string, procAttr *syscall.ProcAttr) {
	switch command[0] { // Сравнение введенной пользователем команды с существующими функциями
	case "cd":
		cd(command)
	case "pwd":
		pwd()
	case "echo":
		echo(command)
	case "kill":
		kill(command)
	case "ps":
		ps()
	case "exit":
		exit()
	default: // Если пользователь ввел команду, которая не реализована, то вызов переадресовывается системе
		forkExec(command[0], command, procAttr)
	}
}

func main() {
	var input string                      // Переменная для ввода пользователя
	os.Args = append(os.Args, "")         // Добавление аргумента old pwd для функции cd
	scanner := bufio.NewScanner(os.Stdin) // Scanner ввода пользователя

	for { // Бесконечный цикл до ввода команды exit
		fmt.Print(">> ")
		scanner.Scan() // Сканирование ввода пользователя

		input = strings.Trim(scanner.Text(), " ") // Обрезание лишних пробелов в начале и конце строки
		pipeline := strings.Split(input, "|")     // Разделение строки по символу | для поддержки pipeline

		switch len(pipeline) {
		case 1: // Если длина слайса pipeline равна 1, значит введена одиночная команда
			var command []string                             // Слайс для команды и аргументов
			for _, item := range strings.Split(input, " ") { // Разделение строки по пробелам
				if item != "" { // Если элемент не пустая строка
					command = append(command, item) // Добавление элемента в слайс
				}
			}

			if len(command) == 0 { // Если длина слайса command равна 0
				fmt.Print("") // Вывод пустой строки
				continue      // Переход к следующей итерации цикла
			}

			procAttr := &syscall.ProcAttr{Files: []uintptr{0, 1, 2}} // Так как команда одиночная, то используются стандартные дескрипторы файла
			execute(command, procAttr)                               // Выполнение команды
		default:
			osInput := os.Stdin // Сохранение стандартных ввода, вывода и вывода ошибок
			osOutput := os.Stdout
			errOutput := os.Stderr

			errReader, errWriter, _ := os.Pipe() // Создание дескриптора для вывода ошибок
			os.Stderr = errWriter                // Замена стандартного дескриптора вывода ошибок на новый

			for index, item := range pipeline { // Перебор строк слайса pipeline
				line := strings.Trim(item, " ") // Обрезание лишних пробелов в начале и конце строки

				var command []string                            // Слайс для команды и аргументов
				for _, elem := range strings.Split(line, " ") { // Разделение строки по пробелам
					if elem != "" { // Если элемент не пустая строка
						command = append(command, elem) // Добавление элемента в слайс
					}
				}

				if len(command) == 0 { // Если длина слайса command равна 0
					fmt.Print("") // Вывод пустой строки
					continue      // Переход к следующей итерации цикла
				}

				pReader, pWriter, _ := os.Pipe() // Создание дескриптора для вывода
				os.Stdout = pWriter              // Замена стандартного дескриптора вывода на новый

				if index == len(pipeline)-1 { // Если это последняя команда
					os.Stdout = osOutput  // Возвращение стандартного вывода
					os.Stderr = errOutput // Возвращение стандартного вывода ошибок
				}

				procAttr := &syscall.ProcAttr{Files: []uintptr{os.Stdin.Fd(), os.Stdout.Fd(), os.Stderr.Fd()}} // Использование измененных дескрипторов

				execute(command, procAttr) // Выполнение команды

				pWriter.Close()    // Закрытие дескриптора вывода
				os.Stdin = pReader // Изменение дескриптора ввода
			}

			os.Stdin = osInput            // Возвращение стандартного дескриптора ввода
			errWriter.Close()             // Закрытие дескриптора вывода
			io.Copy(errOutput, errReader) // Вывод ошибок
		}
	}
}
