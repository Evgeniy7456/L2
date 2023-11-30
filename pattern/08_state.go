package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

Состояние - шаблон проектирования, который позволяет объекту изменять свое поведение в зависимости от внутреннего состояния.

Когда применять?
	- Когда поведение объекта должно зависеть от его состояния и может изменяться динамически во время выполнения.
	- Когда в коде методов объекта используются многочисленные условные конструкции, выбор которых зависит от текущего
	состояния объекта.

Плюсы:
	- избавляет от множества больших условных операторов машины состояний;
	- концентрирует в одном месте код, связанный с определённым состоянием;
	- упрощает код контекста.
Минусы:
	- может неоправданно усложнить код, если состояний мало и они редко меняются.

Ниже приведена реализация паттерна "Состояние" на примере изменений состояния воды. Структура Water содержит состояние
воды. Состояние воды реализует интерфейс WaterState, который описывает метод нагрева Heat и охлаждения Frost. Вода
может быть в состоянии жидкости, пара или льда. В зависимости от состояния воды методы заморозки и охлаждения будут
давать разный результат. Каждое состояние воды реализует интерфейс WaterState. Вода может изменить свое состояние
после применения к ней одного из методов.
*/

// Структура Вода
type Water struct {
	State WaterState // Состояние воды
}

// Создать воду в определенном состоянии
func NewWater(waterState WaterState) *Water {
	return &Water{waterState}
}

// Нагреть воду
func (water *Water) Heat() {
	water.State.Heat(water)
}

// Остудить воду
func (water *Water) Frost() {
	water.State.Frost(water)
}

// Интерфейс действий с водой
type WaterState interface {
	Heat(water *Water) // Нагреть
	Frost(wate *Water) // Охладить
}

// Структура Лед
type SolidWaterState struct{}

// Создать состояние льда
func NewSolidWaterState() *SolidWaterState {
	return &SolidWaterState{}
}

// Нагреть лед
func (solid *SolidWaterState) Heat(water *Water) {
	fmt.Println("Превращаем лед в жидкость")
	water.State = NewLiquidWaterState()
}

// Охладить лед
func (solid *SolidWaterState) Frost(water *Water) {
	fmt.Println("Продолжаем заморозку льда")
}

// Структура Жидкость
type LiquidWaterState struct{}

// Создать состояние жидкости
func NewLiquidWaterState() *LiquidWaterState {
	return &LiquidWaterState{}
}

// Нагреть жидкость
func (liquid *LiquidWaterState) Heat(water *Water) {
	fmt.Println("Превращаем жидкость в пар")
	water.State = NewGasWaterState()
}

// Охладить жидкость
func (liquid *LiquidWaterState) Frost(water *Water) {
	fmt.Println("Превращаем жидкость в лед")
	water.State = NewSolidWaterState()
}

// Структура Пар
type GasWaterState struct{}

// Создать состояние водяного пара
func NewGasWaterState() *GasWaterState {
	return &GasWaterState{}
}

// Нагреть пар
func (gas *GasWaterState) Heat(water *Water) {
	fmt.Println("Повышаем температуру водяного пара")
}

// Охладить пар
func (gas *GasWaterState) Frost(water *Water) {
	fmt.Println("Превращаем водяной пар в жидкость")
	water.State = NewLiquidWaterState()
}

// Взаимодействие клиента с водой
// func main() {
// 	water := NewWater(NewLiquidWaterState())
// 	water.Heat()
// 	water.Frost()
// 	water.Frost()
// }

// Результат работы программы
/*
Превращаем жидкость в пар
Превращаем водяной пар в жидкость
Превращаем жидкость в лед
*/
