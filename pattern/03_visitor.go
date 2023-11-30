package pattern

import (
	"fmt"
	"strings"
)

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

Паттерн "Посетитель" позволяет определить операцию для объектов других классов без изменения этих классов.

Когда применять?
	- Когда имеется много объектов разнородных классов с разными интерфейсами, и требуется выполнить ряд операций над
	каждым из этих объектов.
	- Когда классам необходимо добавить одинаковый набор операций без изменения этих классов.
	- Когда часто добавляются новые операции к классам, при этом общая структура классов стабильна и практически не изменяется.

Плюсы:
	- упрощает добавление операций, работающих со сложными структурами объектов;
	- объединяет родственные операции в одном классе;
	- посетитель может накапливать состояние при обходе структуры элементов.
Минусы:
	- паттерн не оправдан, если иерархия элементов часто меняется;
	- может привести к нарушению инкапсуляции элементов.

Ниже приведена реализация паттерна "Посетитель" на примере добавления функционала сериализации данных пользователей
банка в HTML и XML. Структура Bank содержит информацию о пользователях и методы для взаимодействия с ними. Структура
Person описывает пользователей физ. лиц, а структура Company - юр. лиц. При вызове метода Accept у структуры Bank с
аргументом, реализующим интерфейс Visitor, вызывается метод Accept у всех пользователей банка. У пользователей
Person метод Accept вызывает метод для Person у "посетителя", а у Company - метод для Company. Таким образом, для
добавления нового функционала достаточно создать новую структуру и реализовать интерфейс Visitor. При этом остальные
структуры не изменяются.
*/

// Интерфейс для посетителей
type Visitor interface {
	VisitPersonAccount(account Person)
	VisitCompanyAccount(account Company)
}

// Реализация сериализации в формат HTML
type HtmlVisitor struct{}

// Метод для пользователей Person
func (visitor HtmlVisitor) VisitPersonAccount(account Person) {
	var result strings.Builder
	result.WriteString("<table><tr><td>Свойство<td><td>Значение</td></tr>\n")
	result.WriteString(fmt.Sprintf("<tr><td>Name<td><td>%s</td></tr>\n", account.Name))
	result.WriteString(fmt.Sprintf("<tr><td>Number<td><td>%s</td></tr>\n", account.Number))
	fmt.Println(result.String())
}

// Метод для пользователей Company
func (visitor HtmlVisitor) VisitCompanyAccount(account Company) {
	var result strings.Builder
	result.WriteString("<table><tr><td>Свойство<td><td>Значение</td></tr>\n")
	result.WriteString(fmt.Sprintf("<tr><td>Name<td><td>%s</td></tr>\n", account.Name))
	result.WriteString(fmt.Sprintf("<tr><td>RegNumber<td><td>%s</td></tr>\n", account.RegNumber))
	result.WriteString(fmt.Sprintf("<tr><td>Number<td><td>%s</td></tr>\n", account.Number))
	fmt.Println(result.String())
}

// Создание нового HTML посетителя
func NewHtmlVisitor() HtmlVisitor {
	return HtmlVisitor{}
}

// Реализация сериализации в формат XML
type XmlVisitor struct{}

// Метод для пользователей Person
func (visitor XmlVisitor) VisitPersonAccount(account Person) {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("<Person><Name>%s</Name>", account.Name))
	result.WriteString(fmt.Sprintf("<Number>%s</Number><Person>\n", account.Number))
	fmt.Println(result.String())
}

// Метод для пользователей Company
func (visitor XmlVisitor) VisitCompanyAccount(account Company) {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("<Company><Name>%s</Name>", account.Name))
	result.WriteString(fmt.Sprintf("<RegNumber>%s</RegNumber>", account.RegNumber))
	result.WriteString(fmt.Sprintf("<Number>%s</Number><Company>\n", account.Number))
	fmt.Println(result.String())
}

// Создание нового XML посетителя
func NewXmlVisitor() XmlVisitor {
	return XmlVisitor{}
}

// Структура Bank
type Bank struct {
	Account map[Account]struct{} // Map с пользователями банка
}

// Создание нового банка
func NewBank() Bank {
	return Bank{make(map[Account]struct{})}
}

// Добавление в банк пользователя
func (bank *Bank) Add(account Account) {
	bank.Account[account] = struct{}{}
}

// Удаление пользователя
func (bank *Bank) Remove(account Account) {
	delete(bank.Account, account)
}

// Вызов у каждого пользователя его реализации метода Accept
func (bank *Bank) Accept(visitor Visitor) {
	for item := range bank.Account {
		item.Accept(visitor)
	}
}

// Интерфейс для пользователей
type Account interface {
	Accept(visitor Visitor)
}

// Структура пользователя банка - физ. лица
type Person struct {
	Name   string
	Number string
}

// Создание пользователя Person
func NewPerson(name, number string) Person {
	return Person{name, number}
}

// Вызов у посетителя метода для структуры Person
func (person Person) Accept(visitor Visitor) {
	visitor.VisitPersonAccount(person)
}

// Структура пользователя банка - юр. лица
type Company struct {
	Name      string
	RegNumber string
	Number    string
}

// Создание пользователя Company
func NewCompany(name, regNumber, number string) Company {
	return Company{name, regNumber, number}
}

// Вызов у посетителя метода для структуры Company
func (company Company) Accept(visitor Visitor) {
	visitor.VisitCompanyAccount(company)
}

// Использование посетителей клиентом
// func main() {
// 	bank := NewBank()
// 	bank.Add(NewPerson("Иванов Иван Иванович", "12345678"))
// 	bank.Add(NewCompany("Company", "reg1234", "87654321"))
// 	bank.Accept(NewHtmlVisitor())
// 	bank.Accept(NewXmlVisitor())
// }

// Результат выполнения программы
/*
<table><tr><td>Свойство<td><td>Значение</td></tr>
<tr><td>Name<td><td>Иванов Иван Иванович</td></tr>
<tr><td>Number<td><td>12345678</td></tr>

<table><tr><td>Свойство<td><td>Значение</td></tr>
<tr><td>Name<td><td>Company</td></tr>
<tr><td>RegNumber<td><td>reg1234</td></tr>
<tr><td>Number<td><td>87654321</td></tr>

<Person><Name>Иванов Иван Иванович</Name><Number>12345678</Number><Person>

<Company><Name>Company</Name><RegNumber>reg1234</RegNumber><Number>87654321</Number><Company>
*/
