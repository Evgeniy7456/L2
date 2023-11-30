Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Вывод: error

Эта программа похожа на 3. Интерфейс также не равен nil, потому что поле с типом не равно nil.
Оно равно *customError. Поэтому поле значение равное nil не влияет на сравнение, так как оно
происходит по первому полю.

```
