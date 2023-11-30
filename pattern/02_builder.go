package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
"Строитель" это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово. Строитель даёт
возможность использовать один и тот же код строительства для получения разных представлений объектов.

Когда применять?
	- Когда процесс создания нового объекта не должен зависеть от того, из каких частей этот объект состоит и как эти
	части связаны между собой.
	- Когда необходимо обеспечить получение различных вариаций объекта в процессе его создания.

Плюсы:
	- позволяет создавать объекты пошагово;
	- позволяет использовать один и тот же код для создания различных объектов;
	- изолирует сложный код сборки объекта от его основной бизнес-логики.
Минусы:
	- усложняет код программы из-за введения дополнительных классов;
	- клиент может оказаться привязан к конкретным классам строителей, так как в интерфейсе строителя может не быть
	метода получения результата.

Ниже приведена реализация паттерна "Строитель" на примере выпечки хлеба. Хлеб разных сортов состоит из разных
ингредиентов, но при этом процесс выпечки одинаковый. Был определен интерфейс BreadBuilder, описывающий функции
необходимые для выпечки хлеба. Каждый сорт хлеба реализует данный интерфейс по-своему. Структура Baker содержит
экземпляр сорта хлеба и метод Bake(), который задает порядок выпечки.
*/

// Структура Хлеб
type Bread struct {
	Name      string // Название хлеба
	Flour     string // Сорт муки
	Salt      string // Соль
	Additives string // Добавки
}

// Функция возвращающая название хлеба и его состав
func (bread Bread) String() string {
	return fmt.Sprintf("%s. Состав: %s, %s, добавки: %s.", bread.Name, bread.Flour, bread.Salt, bread.Additives)
}

// Константы задающие сорт хлеба
const (
	RyeBreadType   = "rye"   // Ржаной
	WheatBreadType = "wheat" // Пшеничный
)

// Интерфейс описывающий функции для выпечки хлеба
type BreadBuilder interface {
	SetName()        // Задать название хлеба
	SetFlour()       // Добавить муку
	SetSalt()        // Посолить
	SetAdditives()   // Использовать пищевые добавки
	GetBread() Bread // Получить готовый хлеб
}

// Функция возвращающая экземпляр заданного сорта хлеба
func GetBreadBuilder(breadType string) BreadBuilder {
	switch breadType {
	case RyeBreadType:
		return &RyeBreadBuilder{&Bread{}}
	case WheatBreadType:
		return &WheatBreadBuilder{&Bread{}}
	default:
		return nil
	}
}

// Структура Пекарь
type Baker struct {
	BreadBuilder BreadBuilder
}

// Создать пекаря
func NewBaker(breadBuilder BreadBuilder) *Baker {
	return &Baker{breadBuilder}
}

// Изменить сорт выпекаемого хлеба
func (baker *Baker) SetBuilder(breadBuilder BreadBuilder) {
	baker.BreadBuilder = breadBuilder
}

// Испечь хлеб
func (baker Baker) Bake() Bread {
	baker.BreadBuilder.SetName()
	baker.BreadBuilder.SetFlour()
	baker.BreadBuilder.SetSalt()
	baker.BreadBuilder.SetAdditives()
	return baker.BreadBuilder.GetBread()
}

// Реализация интерфейса BreadBuilder для выпечки ржаного хлеба
type RyeBreadBuilder struct {
	Bread *Bread
}

func (rye *RyeBreadBuilder) SetName() {
	rye.Bread.Name = "Ржаной хлеб"
}

func (rye *RyeBreadBuilder) SetFlour() {
	rye.Bread.Flour = "ржаная мука 1 сорт"
}

func (rye *RyeBreadBuilder) SetSalt() {
	rye.Bread.Salt = "соль"
}

func (rye *RyeBreadBuilder) SetAdditives() {
	rye.Bread.Additives = "отсутствуют"
}

func (rye RyeBreadBuilder) GetBread() Bread {
	return *rye.Bread
}

// Реализация интерфейса BreadBuilder для выпечки пшеничного хлеба
type WheatBreadBuilder struct {
	Bread *Bread
}

func (wheat *WheatBreadBuilder) SetName() {
	wheat.Bread.Name = "Пшеничный хлеб"
}

func (wheat *WheatBreadBuilder) SetFlour() {
	wheat.Bread.Flour = "пшеничная мука высший сорт"
}

func (wheat *WheatBreadBuilder) SetSalt() {
	wheat.Bread.Salt = "соль"
}

func (wheat *WheatBreadBuilder) SetAdditives() {
	wheat.Bread.Additives = "улучшитель хлебопекарный"
}

func (wheat WheatBreadBuilder) GetBread() Bread {
	return *wheat.Bread
}

// Выпечка хлеба клиентом
// func main() {
// 	builder := GetBreadBuilder(RyeBreadType)
// 	baker := NewBaker(builder)
// 	ryeBread := baker.Bake()
// 	fmt.Println(ryeBread)

// 	builder = GetBreadBuilder(WheatBreadType)
// 	baker.SetBuilder(builder)
// 	wheatBread := baker.Bake()
// 	fmt.Println(wheatBread)
// }

// Результат выполнения программы
/*
Ржаной хлеб. Состав: ржаная мука 1 сорт, соль, добавки: отсутствуют.
Пшеничный хлеб. Состав: пшеничная мука высший сорт, соль, добавки: улучшитель хлебопекарный.
*/
