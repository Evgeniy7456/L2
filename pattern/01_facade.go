package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
"Фасад" представляет собой структурный шаблон проектирования, который дает возможность скрыть сложность системы посредством
сведения всевозможных внешних вызовов к одному объекту, делегирующему эти вызовы соответствующим объектам системы.

Когда применять?
	- Когда имеется сложная система, и необходимо упростить с ней работу. Фасад позволит определить одну точку
	взаимодействия между клиентом и системой.
	- Когда надо уменьшить количество зависимостей между клиентом и сложной системой. Фасадные объекты позволяют отделить,
	изолировать компоненты системы от клиента и развивать и работать с ними независимо.
	- Когда нужно определить подсистемы компонентов в сложной системе. Создание фасадов для компонентов каждой отдельной
	подсистемы позволит упростить взаимодействие между ними и повысить их независимость друг от друга.

Плюсы:
	- упрощение использования сложной подсистемы.
Минусы:
	- фасад может стать объектом от которого зависит работа всей программы.

Ниже приведена реализация паттерна "Фасад" на примере запуска компьютера. Компьютер состоит из таких подсистем как процессор,
оперативная память, жесткий диск. Они представлены в виде структур со своими методами. Запуск компьютера заключается в
использовании подсистем в определенной последовательности. В качестве "фасада" выступает структура computer. Пользователю для
запуска компьютера необходимо только получить экземпляр структуры computer с помощью функции NewComputer и применить метод
StartComputer. Функция NewComputer инициализирует подсистемы компьютера, а метод StartComputer используя методы подсистем
запускает компьютер. Пользователю предоставляется простой интерфейс, позволяющий запустить компьютер, при этом сложная
реализация скрыта от него.
*/

// Структура процессор
type cpu struct{}

func (cpu *cpu) freeze() {}

func (cpu *cpu) jump(address int) {}

func (cpu *cpu) execute() {}

// Структура оперативная память
type memory struct{}

func (memory *memory) load(position int, data []byte) {}

// Структура жесткий диск
type hardDriver struct{}

func (hardDriver *hardDriver) read(lba int, size int) []byte { return []byte("") }

// Структура компьютер
type computer struct {
	CPU          *cpu
	Memory       *memory
	HardDriver   *hardDriver
	BOOT_ADDRESS int
	BOOT_SECTOR  int
	SECTOR_SIZE  int
}

// Создание компьютера
func NewComputer() *computer {
	return &computer{
		CPU:        &cpu{},
		Memory:     &memory{},
		HardDriver: &hardDriver{},
	}
}

// Метод запускающий компьютер
func (computer *computer) StartComputer() {
	computer.CPU.freeze()
	computer.Memory.load(computer.BOOT_ADDRESS, computer.HardDriver.read(computer.BOOT_SECTOR, computer.SECTOR_SIZE))
	computer.CPU.jump(computer.BOOT_ADDRESS)
	computer.CPU.execute()
}

// Запуск компьютера у клиента
// func main() {
// 	computer := NewComputer()
// 	computer.StartComputer()
// }