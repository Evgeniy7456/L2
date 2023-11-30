package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

Фабричный метод - это паттерн, который определяет интерфейс для создания объектов некоторого класса, но непосредственное
решение о том, объект какого класса создавать происходит в подклассах. То есть паттерн предполагает, что базовый класс
делегирует создание объектов классам-наследникам.

Когда применять?
	- Когда заранее неизвестно, объекты каких типов необходимо создавать.
	- Когда система должна быть независимой от процесса создания новых объектов и расширяемой: в нее можно легко вводить
	новые классы, объекты которых система должна создавать.
	- Когда создание новых объектов необходимо делегировать из базового класса классам наследникам.

Плюсы:
	- избавляет главный класс от привязки к конкретным типам объектов;
	- выделяет код производства объектов в одно место, упрощая поддержку кода;
	- упрощает добавление новых типов объектов в программу;
	- реализует принцип открытости/закрытости.
Минусы:
	- может привести к созданию больших параллельных иерархий классов, так как для каждого типа объекта надо создать свой
	подкласс создателя.

Ниже приведена реализация паттерна "Фабричный метод" на примере строительства домов. Интерфейсы Developer и House
описывают строителя домов и дом, который он может построить. Структура ConcreteDeveloper реализует интерфейс Developer.
Структуры PanelHouse и WoodHouse реализуют интерфейс House. В зависимости от типа дома, который передается методу Create
структуры ConcreteDeveloper в качестве параметра, строитель может построить панельный или деревянный дом.
*/

// Тип дома
type HouseType string

// Возможные типы домов
const (
	PanelType HouseType = "Панельный дом"
	WoodType  HouseType = "Деревянный дом"
)

// Интерфейс Дом
type House interface {
	HouseType() HouseType
}

// Интерфейс Строитель
type Developer interface {
	Create(houseType HouseType) House
}

// Реализация интерфейса Строитель
type ConcreteDeveloper struct{}

// Создание строителя
func NewDeveloper() Developer {
	return &ConcreteDeveloper{}
}

// Постройка дома
func (developer *ConcreteDeveloper) Create(houseType HouseType) House {
	var house House
	switch houseType { // Сравнение типа дома с возможными типами
	case PanelType:
		house = NewPanelHouse()
	case WoodType:
		house = NewWoodHouse()
	default:
		fmt.Println("Невозможно построить дом типа:", houseType)
	}
	return house
}

// Панельный дом
type PanelHouse struct {
	Type HouseType
}

// Построить панельный дом
func NewPanelHouse() House {
	return &PanelHouse{PanelType}
}

// Получение типа дома
func (panel *PanelHouse) HouseType() HouseType {
	return panel.Type
}

// Деревянный дом
type WoodHouse struct {
	Type HouseType
}

// Построить деревянный дом
func NewWoodHouse() House {
	return &WoodHouse{WoodType}
}

// Получение типа дома
func (wood *WoodHouse) HouseType() HouseType {
	return wood.Type
}

// Строительство домов клиентом
// func main() {
// 	developer := NewDeveloper()

// 	house := developer.Create(PanelType)
// 	fmt.Println(house.HouseType())

// 	house = developer.Create(WoodType)
// 	fmt.Println(house.HouseType())
// }

// Результат выполнения программы
/*
Панельный дом
Деревянный дом
*/
