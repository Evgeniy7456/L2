package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

Цепочка обязанностей (Chain of responsibility) - поведенческий шаблон проектирования, который позволяет избежать жесткой
привязки отправителя запроса к получателю. Все возможные обработчики запроса образуют цепочку, а сам запрос перемещается
по этой цепочке. Каждый объект в этой цепочке при получении запроса выбирает, либо закончить обработку запроса, либо
передать запрос на обработку следующему по цепочке объекту.

Когда применять?
	- Когда имеется более одного объекта, который может обработать определенный запрос.
	- Когда надо передать запрос на выполнение одному из нескольких объект, точно не определяя, какому именно объекту.
	- Когда набор объектов задается динамически.

Плюсы:
	- уменьшает зависимость между клиентом и обработчиками;
	- реализует принцип единственной обязанности;
	- реализует принцип открытости/закрытости.
Минусы:
	- запрос может остаться никем не обработанным.

Ниже приведена реализация паттерна "Цепочка обязанностей" на примере денежного перевода. Получатель может получить
перевод денег только на PayPal. Банковского счета у него нет. При этом отправителю необязательно это знать. Он отправляет
деньги банковским переводом. Если у получателя нет банковского счета, тогда обработка перевода от банка передается дальше
по цепочке.
*/

// Структура Получателя
type Receiver struct {
	BankTransfer   bool // Счет в банке
	MoneyTransfer  bool // Платежные системы WesternUnion, Unistream
	PayPalTransfer bool // Счет в PayPal
}

// Создание получателя
func NewReceiver(BankTransfer, MoneyTransfer, PayPalTransfer bool) *Receiver {
	return &Receiver{BankTransfer, MoneyTransfer, PayPalTransfer}
}

// Интерфейс обработки денежного перевода
type PaymentHandler interface {
	Handle(receiver *Receiver)
}

// Обработчик банковского перевода
type BankPaymentHandler struct {
	Successor PaymentHandler // Следующий обработчик
}

// Создание обработчика
func NewBankPaymentHandler() *BankPaymentHandler {
	return &BankPaymentHandler{}
}

// Обработка банковского перевода
func (bank *BankPaymentHandler) Handle(receiver *Receiver) {
	if receiver.BankTransfer { // Проверка наличия банковского счета у получателя
		fmt.Println("Банковский перевод выполнен")
	} else if bank.Successor != nil { // Если счета нет, то вызывается следующий обработчик
		bank.Successor.Handle(receiver)
	}
}

// Обработчик PayPal
type PayPalPaymentHandler struct {
	Successor PaymentHandler // Следующий обработчик
}

// Создание обработчика
func NewPayPalPlaymentHandler() *PayPalPaymentHandler {
	return &PayPalPaymentHandler{}
}

// Обработка перевода через PayPal
func (paypal *PayPalPaymentHandler) Handle(receiver *Receiver) {
	if receiver.PayPalTransfer { // Проверка наличия счета в PayPal у получателя
		fmt.Println("Перевод через PayPal выполнен")
	} else if paypal.Successor != nil { // Если счета нет, то вызывается следующий обработчик
		paypal.Successor.Handle(receiver)
	}
}

// Обработчик системы денежных переводов
type MoneyPaymentHandler struct {
	Successor PaymentHandler // Следующий обработчик
}

// Создание обработчика
func NewMoneyPaymentHandler() *MoneyPaymentHandler {
	return &MoneyPaymentHandler{}
}

// Обработка перевода через системы денежных переводов
func (money *MoneyPaymentHandler) Handle(receiver *Receiver) {
	if receiver.MoneyTransfer { // Проверка наличия счета в системах денежного перевода у получателя
		fmt.Println("Перевод через системы денежных переводов выполнен")
	} else if money.Successor != nil { // Если счета нет, то вызывается следующий обработчик
		money.Successor.Handle(receiver)
	}
}

// Выполнение перевода клиентом
// func main() {
// 	receiver := NewReceiver(false, false, true)

// 	bankPaymentHandler := NewBankPaymentHandler()
// 	moneyPaymentHandler := NewMoneyPaymentHandler()
// 	paypalPaymentHandler := NewPayPalPlaymentHandler()

// 	bankPaymentHandler.Successor = paypalPaymentHandler
// 	paypalPaymentHandler.Successor = moneyPaymentHandler

// 	bankPaymentHandler.Handle(receiver)
// }

// Результат выполнения программы
/*
Перевод через PayPal выполнен
*/
