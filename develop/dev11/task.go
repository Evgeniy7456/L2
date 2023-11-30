package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

type serverConfig struct { // Конфигурация сервера
	Port string // Порт сервера
}

type event struct { // Структура event
	User_id int    // id пользователя
	Date    string // Дата
	Info    string // Информация о событии
}

type report struct { // Отчет о выполнении запроса пользователя
	Type   string // Тип отчета
	Result string // Результат
}

var data = make(map[int]map[string]string) // Структура map содержащая id пользователя и еще одну структуру map с датой и информацией о событии

func main() {
	http.HandleFunc("/create_event", requestLog(createEvent))        // Обработчик - создать событие
	http.HandleFunc("/update_event", requestLog(updateEvent))        // Обработчик - обновить событие
	http.HandleFunc("/delete_event", requestLog(deleteEvent))        // Обработчик - удалить событие
	http.HandleFunc("/events_for_day", requestLog(eventsForDay))     // Обработчик - событие за день
	http.HandleFunc("/events_for_week", requestLog(eventsForWeek))   // Обработчик - события за неделю
	http.HandleFunc("/events_for_month", requestLog(eventsForMonth)) // Обработчик - собятия за месяц

	file, err := os.ReadFile("config.json") // Чтение конфигурационного файла сервера
	if err != nil {
		log.Fatal(err)
	}

	config := &serverConfig{}           // Создание экземпляра структуры конфигурации сервера
	err = json.Unmarshal(file, &config) // Десериализация json
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(":"+config.Port, nil) // Запуск сервера
	if err != nil {
		log.Fatal(err)
	}
}

// Обработчик создания события
func createEvent(w http.ResponseWriter, r *http.Request) {
	event, ok := postDataChecking(w, r) // Проверка полученных от пользователя данных

	if !ok { // Если данные некорректные
		return // Выполнение функции прерывается
	}

	if _, ok := data[event.User_id][event.Date]; ok { // Проверка на существование события
		w.WriteHeader(500)                                            // Возвращение пользователю ошибки 500
		report := getJson(&report{"error", "событие уже существует"}) // Создание отчета
		json.NewEncoder(w).Encode(report)                             // Сериализация json
		log.Println("ERROR: the event already exists")                // Вывод информации в лог
		return
	}

	if data[event.User_id] == nil {
		data[event.User_id] = make(map[string]string) // Инициализация map
	}
	data[event.User_id][event.Date] = event.Info // Добавление события в map

	report := getJson(&report{"result", "событие создано"}) // Создание отчета
	json.NewEncoder(w).Encode(report)                       // Сериализация json

	log.Printf("Create event | user_id %d date %s", event.User_id, event.Date) // Вывод информации в лог
}

// Обработчик обновления события
func updateEvent(w http.ResponseWriter, r *http.Request) {
	event, ok := postDataChecking(w, r) // Проверка полученных от пользователя данных

	if !ok { // Если данные некорректные
		return // Выполнение функции прерывается
	}

	if _, ok := data[event.User_id][event.Date]; !ok { // Проверка на существование события
		w.WriteHeader(500)                                           // Возвращение пользователю ошибки 500
		report := getJson(&report{"error", "события не существует"}) // Создание отчета
		json.NewEncoder(w).Encode(report)                            // Сериализация json
		log.Println("ERROR: event does not exist")                   // Вывод информации в лог
		return
	}

	data[event.User_id][event.Date] = event.Info // Изменение события

	report := getJson(&report{"result", "событие изменено"}) // Создание отчета
	json.NewEncoder(w).Encode(report)                        // Сериализация json

	log.Printf("Update event | user_id %d date %s", event.User_id, event.Date) // Вывод информации в лог
}

// Обработчик удаления события
func deleteEvent(w http.ResponseWriter, r *http.Request) {
	event, ok := postDataChecking(w, r) // Проверка полученных от пользователя данных

	if !ok { // Если данные некорректные
		return // Выполнение функции прерывается
	}

	if _, ok := data[event.User_id][event.Date]; !ok { // Проверка на существование события
		w.WriteHeader(500)                                           // Возвращение пользователю ошибки 500
		report := getJson(&report{"error", "события не существует"}) // Создание отчета
		json.NewEncoder(w).Encode(report)                            // Сериализация json
		log.Println("ERROR: event does not exist")                   // Вывод информации в лог
		return
	}

	delete(data[event.User_id], event.Date) // Удаление события

	report := getJson(&report{"result", "событие удалено"}) // Создание отчета
	json.NewEncoder(w).Encode(report)                       // Сериализация json

	log.Printf("Delete event | user_id %d date %s", event.User_id, event.Date) // Вывод информации в лог
}

// Обработчик получения информации о событии за день
func eventsForDay(w http.ResponseWriter, r *http.Request) {
	id, date, ok := getDataChecking(w, r) // Проверка полученных от пользователя данных

	if !ok { // Если данные некорректные
		return // Выполнение функции прерывается
	}

	event, ok := data[id][date] // Получение информации о событии
	if !ok {                    // Если события не существует
		w.WriteHeader(500)                                         // Возвращение пользователю ошибки 500
		report := getJson(&report{"error", "событие отсутствует"}) // Создание отчета
		json.NewEncoder(w).Encode(report)                          // Сериализация json
		log.Println("ERROR: event does not exist")                 // Вывод информации в лог
		return
	}

	report := getJson(&report{"result", fmt.Sprintf("date: %s | event: %s", date, event)}) // Создание отчета
	json.NewEncoder(w).Encode(report)                                                      // Сериализация json

	log.Printf("Events for day | user_id %d date %s", id, date) // Вывод информации в лог
}

// Обработчик получения информации о событии за неделю
func eventsForWeek(w http.ResponseWriter, r *http.Request) {
	id, date, ok := getDataChecking(w, r) // Проверка полученных от пользователя данных

	if !ok { // Если данные некорректные
		return // Выполнение функции прерывается
	}

	layout := "2006-01-02"                  // Формат даты
	dateTime, _ := time.Parse(layout, date) // Получение даты в соответствии с форматом

	weekDay := int(dateTime.Weekday())                                  // Получение номера для недели
	dateTime = dateTime.Add(-24 * time.Duration(weekDay-1) * time.Hour) // Получение даты понедельника

	var result string        // Переменная для результата
	for i := 0; i < 7; i++ { // Цикл для прохода по дням недели
		date = dateTime.Format(layout) // Преобразование из time.Time в string в соответствии с форматом даты

		if event, ok := data[id][date]; ok { // Если событие существует
			result += fmt.Sprintf("date: %s | event: %s\n", date, event) // Добавление события в результат
		}

		dateTime = dateTime.Add(24 * time.Hour) // Получение даты следующего дня
	}

	if result == "" { // Если результат пустая строка
		w.WriteHeader(500)                                         // Возвращение пользователю ошибки 500
		report := getJson(&report{"error", "события отсутствуют"}) // Создание отчета
		json.NewEncoder(w).Encode(report)                          // Сериализация json
		log.Println("ERROR: no events exist")                      // Вывод информации в лог
		return
	}

	report := getJson(&report{"result", result}) // Создание отчета
	json.NewEncoder(w).Encode(report)            // Сериализация json

	log.Printf("Events for week | user_id %d date %s", id, date) // Вывод информации в лог
}

// Обработчик получения информации о событии за месяц
func eventsForMonth(w http.ResponseWriter, r *http.Request) {
	id, date, ok := getDataChecking(w, r) // Проверка полученных от пользователя данных

	if !ok { // Если данные некорректные
		return // Выполнение функции прерывается
	}

	layout := "2006-01-02"                  // Формат даты
	timeDate, _ := time.Parse(layout, date) // Получение даты в соответствии с форматом

	var result string          // Переменная для результата
	for i := 1; i <= 31; i++ { // Цикл для прохода по дням недели
		date := time.Date(timeDate.Year(), timeDate.Month(), i, 0, 0, 0, 0, time.Local).Format(layout) // Получение даты в формате string

		if event, ok := data[id][date]; ok { // Если событие существует
			result += fmt.Sprintf("date: %s | event: %s\n", date, event) // Добавление события в результат
		}
	}

	if result == "" { // Если результат пустая строка
		w.WriteHeader(500)                                               // Возвращение пользователю ошибки 500
		report := getJson(&report{"error", "событий в этом месяце нет"}) // Создание отчета
		json.NewEncoder(w).Encode(report)                                // Сериализация json
		log.Println("ERROR: no events exist")                            // Вывод информации в лог
		return
	}

	report := getJson(&report{"result", result}) // Создание отчета
	json.NewEncoder(w).Encode(report)            // Сериализация json

	log.Printf("Events for month | user_id %d date %s", id, date) // Вывод информации в лог
}

// Middleware для логирования запроса
func requestLog(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logging HTTP %s request", r.URL) // Вывод информации в лог
		next(w, r)                                   // Вызов обработчика
	}
}

// Парсинг данных
func parse(r io.ReadCloser) (*event, error) {
	event := &event{}                       // Создание экземпляра структуры события
	err := json.NewDecoder(r).Decode(event) // Десериализация json
	if err != nil {                         // Если не удалось получить преобразовать данные
		return nil, fmt.Errorf("ошибка формата данных") // Вывод ошибки
	}

	return event, err
}

// Валидация данных
func valid(event *event) error {
	if event.User_id < 0 { // Если user_id отрицательное число
		return fmt.Errorf("некорректные данные: user_id должен быть положительным") // Вывод ошибки
	}

	layout := "2006-01-02"                                    // Формат даты
	if _, err := time.Parse(layout, event.Date); err != nil { // Если не удалось преобразовать дату в необходимый формат
		return fmt.Errorf(fmt.Sprintf("некорректные данные: дата %s несоответствует формату yyyy-mm-dd", event.Date)) // Вывод ошибки
	}

	return nil
}

// Проверка данных для POST-методов
func postDataChecking(w http.ResponseWriter, r *http.Request) (*event, bool) {
	event, err := parse(r.Body) // Вызов метода парсинга данных

	if err != nil {
		w.WriteHeader(503)                               // Возвращение пользователю ошибки 503
		report := getJson(&report{"error", err.Error()}) // Создание отчета
		json.NewEncoder(w).Encode(report)                // Сериализация json
		return nil, false
	}

	err = valid(event)
	if err != nil {
		w.WriteHeader(400)                               // Возвращение пользователю ошибки 400
		report := getJson(&report{"error", err.Error()}) // Создание отчета
		json.NewEncoder(w).Encode(report)                // Сериализация json
		return nil, false
	}

	return event, true
}

// Проверка данных для GET-методов
func getDataChecking(w http.ResponseWriter, r *http.Request) (int, string, bool) {
	id, err := strconv.Atoi(r.URL.Query().Get("user_id")) // Получение id пользователя из query string
	if err != nil {
		w.WriteHeader(503)                                                 // Возвращение пользователю ошибки 503
		report := getJson(&report{"error", "значение user_id не найдено"}) // Создание отчета
		json.NewEncoder(w).Encode(report)                                  // Сериализация json
		return 0, "", false
	}

	if id < 0 {
		w.WriteHeader(400)                                                       // Возвращение пользователю ошибки 400
		report := getJson(&report{"error", "user_id должен быть положительным"}) // Создание отчета
		json.NewEncoder(w).Encode(report)                                        // Сериализация json
		return 0, "", false
	}

	if _, ok := data[id]; !ok {
		w.WriteHeader(400)                                                                                          // Возвращение пользователю ошибки 400
		report := getJson(&report{"error", fmt.Sprintf("у пользователя с user_id %d нет созданных событий\n", id)}) // Создание отчета
		json.NewEncoder(w).Encode(report)                                                                           // Сериализация json
		return 0, "", false
	}

	date := r.URL.Query().Get("date") // Получение даты из query string
	layout := "2006-01-02"            // Формат даты
	if _, err = time.Parse(layout, date); err != nil {
		w.WriteHeader(400)                                                                                    // Возвращение пользователю ошибки 400
		report := getJson(&report{"error", fmt.Sprintf("дата %s несоответствует формату yyyy-mm-dd;", date)}) // Создание отчета
		json.NewEncoder(w).Encode(report)                                                                     // Сериализация json
		return 0, "", false
	}

	return id, date, true
}

// Сериализация json
func getJson(report *report) []byte {
	data := make(map[string]string) // Создание map для преобразования в json

	if report.Type == "error" { // Если тип отчета "error"
		data["error"] = report.Result // Добавление результата
	} else { // Если тип отчета "result"
		data["result"] = report.Result // Добавление результата
	}

	json_data, _ := json.Marshal(data) // Сериализация json
	return json_data
}
