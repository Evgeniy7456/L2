package pattern

import (
	"container/list"
	"fmt"
)

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

Паттерн "Команда" позволяет инкапсулировать запрос на выполнение определенного действия в виде отдельного
объекта. Этот объект запроса на действие и называется командой. При этом объекты, инициирующие запросы на выполнение
действия, отделяются от объектов, которые выполняют это действие. Команды могут использовать параметры, которые передают
ассоциированную с командой информацию. Кроме того, команды могут ставиться в очередь и также могут быть отменены.

Когда применять?
	- Когда надо передавать в качестве параметров определенные действия, вызываемые в ответ на другие действия. То есть
	когда необходимы функции обратного действия в ответ на определенные действия.
	- Когда необходимо обеспечить выполнение очереди запросов, а также их возможную отмену.
	- Когда надо поддерживать логгирование изменений в результате запросов. Использование логов может помочь восстановить
	состояние системы - для этого необходимо будет использовать последовательность запротоколированных команд.

Плюсы:
	- убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют;
	- позволяет реализовать простую отмену и повтор операций;
	- позволяет реализовать отложенный запуск операций;
	- позволяет собирать сложные команды из простых;
	- реализует принцип открытости/закрытости.
Минусы:
	- усложняет код программы из-за введения множества дополнительных классов.

Ниже приведена реализация паттерна "Команда" на примере использования телевизора. Структура TV содержит поля отвечающие
за громкость звука и методы позволяющие включить телевизор, выключить его и прибавить или убавить громкость звука.
Для взаимодействия с телевизором используется пульт. Структура Pult содержит массив с командами, где каждая команда это
экземпляр структуры, реализующей интерфейс Command. Каждую команду можно выполнить и отменить. Для отмены команд в
структуре Pult используется список и принцип LIFO.
*/

// Интерфейс команды
type Command interface {
	Execute() // Выполнить команду
	Undo()    // Отменить команду
}

// Структура Телевизор
type TV struct {
	minVolume int // Минимальная громкость
	maxVolume int // Максимальная громкость
	level     int // Текущий уровень звука
}

// Создать телевизор
func NewTV() *TV {
	return &TV{0, 20, 0}
}

// Включить телевизор
func (tv *TV) On() {
	fmt.Println("Телевизор включен")
}

// Выключить телевизор
func (tv *TV) Off() {
	fmt.Println("Телевизор выключен")
}

// Прибавить громкость
func (tv *TV) RaiseLevel() {
	if tv.level < tv.maxVolume { // Проверка условия, чтобы текущий уровень звука не превысил максимальный
		tv.level++
	}
	fmt.Printf("Уровень звука %d\n", tv.level)
}

// Убавить громкость
func (tv *TV) DropLevel() {
	if tv.level > tv.minVolume { // Проверка условия, чтобы текущий уровень звука не был ниже минимального
		tv.level--
	}
	fmt.Printf("Уровень звука %d\n", tv.level)
}

// Команда для включения телевизора
type TVOnCommand struct {
	tv *TV // Телевизор к которому будет применяться команда
}

// Создать команду
func NewTVOnCommand(tv *TV) *TVOnCommand {
	return &TVOnCommand{tv}
}

// Включить телевизор
func (tvOnCommand *TVOnCommand) Execute() {
	tvOnCommand.tv.On()
}

// Отменить команду, т.е. выключить телевизор
func (tvOnCommand *TVOnCommand) Undo() {
	tvOnCommand.tv.Off()
}

// Команда для регулирования громкости
type VolumeCommand struct {
	tv *TV // Телевизор к которому будет применяться команда
}

// Создать команду
func NewVolumeCommand(tv *TV) *VolumeCommand {
	return &VolumeCommand{tv}
}

// Прибавить громкость
func (volumeCommand *VolumeCommand) Execute() {
	volumeCommand.tv.RaiseLevel()
}

// Отменить команду, т.е. убавить громкость
func (volumeCommand *VolumeCommand) Undo() {
	volumeCommand.tv.DropLevel()
}

// Структура Пульть
type Pult struct {
	buttons         [2]Command // Массив возможных команд
	commandsHistory list.List  // История вызова команд
}

// Создать пульт
func NewPult() *Pult {
	return &Pult{}
}

// Добавить команду в массив
func (pult *Pult) SetCommand(number int, com Command) {
	pult.buttons[number] = com
}

// Выполнить команду
func (pult *Pult) PressButton(number int) {
	pult.buttons[number].Execute()                      // Выполнение команды
	pult.commandsHistory.PushBack(pult.buttons[number]) // Добавление команды историю вызова
}

// Отменить команду
func (pult *Pult) PressUndoButton() {
	command := pult.commandsHistory.Back() // Получение последней вызванной команды
	command.Value.(Command).Undo()         // Вызов метода отмены команды
	pult.commandsHistory.Remove(command)   // Удаление последней команды из истории вызова
}

// Использование телевизора клиентом
// func main() {
// 	tv := NewTV()
// 	pult := NewPult()
// 	pult.SetCommand(0, NewTVOnCommand(tv))
// 	pult.SetCommand(1, NewVolumeCommand(tv))
// 	pult.PressButton(0)
// 	pult.PressButton(1)
// 	pult.PressButton(1)
// 	pult.PressUndoButton()
// 	pult.PressUndoButton()
// 	pult.PressUndoButton()
// }

// Результат выполнения программы
/*
Телевизор включен
Уровень звука 1
Уровень звука 2
Уровень звука 1
Уровень звука 0
Телевизор выключен
*/