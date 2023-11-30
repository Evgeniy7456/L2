Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				fmt.Println("a" + v)
				c <- v
			case v := <-b:
			fmt.Println("a" + v)
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Вывод: числа от 1 до 8 в случайном порядке и бесконечный вывод 0.

С помощью функции asChan создаются два канала, которые запускают горутины и пишут в них в канал переданные
в качестве аргумента числа. После отправки числа в канал они спят рандомное время от 0 до 999 миллисекунд
включительно. Функция merge объединяет каналы при этом после закрытия этих каналов выхода из цикла for
не происходит. Поэтому горутина созданная в этой функции не завершает свою работу. Так как каналы a и b
закрыты, то при чтении из этих каналов получаем нулевое значение типа int, т.е. 0. Поэтому будет бесконечный
вывод 0.

```
