package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	var timeout int // Переменные под флаги
	var ip string
	var host string
	var port string
	flag.IntVar(&timeout, "timeout", 10, "timeout connection") // Флаг -timeout задает таймаут на подключение к серверу
	flag.StringVar(&ip, "ip", "", "server ip address")         // Флаг -ip используется для получения ip адреса сервера
	flag.StringVar(&host, "host", "", "server hostname")       // Флаг -host используется для получения hostname сервера
	flag.StringVar(&port, "port", "", "server port")           // Флаг -port используется для получения порта сервера
	flag.Parse()

	var address string // Переменная для адреса подключения
	if ip != "" {      // Если используется флаг -ip
		address = ip // Адрес равен ip
	} else if host != "" { // Если используется флаг -host
		address = host // Адрес равен host
	} else { // Если не указаны ip и host
		io.WriteString(os.Stderr, "enter ip or host\n") // Вывод ошибки: введите host или ip
		return                                          // Завершение работы программы
	}

	if port != "" { // Если используется флаг -port
		address += ":" + port // К адресу добавляется ":port"
	} else {
		io.WriteString(os.Stderr, "enter ip or host\n") // Вывод ошибки: введите host или ip
		return                                          // Завершение работы программы
	}

	var conn net.Conn                                      // Переменная для подключения
	var err error                                          // Переменная для ошибки
	to := time.After(time.Duration(timeout) * time.Second) // Задаем таймаут подключения
	done := make(chan bool, 1)                             // Канал для сообщения об успешном или не успешном подключении

	go func() { // Горутина для подключения к серверу
		for { // Бесконечный цикл
			select {
			case <-to: // Если таймаут истек
				io.WriteString(os.Stderr, "failed to connect to server\n") // Ошибка: не удалось подключиться к серверу
				done <- false                                              // В канал отправляется значение false
				return                                                     // Завершение работы горутины
			default: // Если таймаут не истек
				conn, err = net.Dial("tcp", address) // Подключение к серверу
				if err == nil {                      // Если удалось подключиться
					done <- true // В канал отправляется значение true
					return       // Завершение работы горутины
				}
			}
		}
	}()

	result := <-done // Получаем сообщение из горутины, которая подключается к серверу

	if result { // Если удалось подключиться
		defer func() { // Действия при завершении работы программы
			fmt.Println("\nЗавершение работы программы") // Вывод сообщения
			conn.Close()                                 // Закрываем соединение с сервером
		}()

		for { // Бесконечный цикл
			reader := bufio.NewReader(os.Stdin) // Создание объекта для чтение Stdin
			fmt.Print("Введите сообщение: ")    // Сообщения для пользователя с предложением ввести текст

			message, err := reader.ReadString('\n') // Чтение из Stdin
			if err == io.EOF {                      // Проверка на нажатие пользователем сочетания клавиш Ctrl+D
				return // Завершение работы программы
			}
			_, err = fmt.Fprint(conn, message) // Отправка сообщения на сервер
			if err != nil {                    // Если соединение с сервером прервано
				io.WriteString(os.Stderr, err.Error()) // Вывод ошибки
				return                                 // Завершение работы программы
			}
		}
	}
}
