package main

import (
	"fmt"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// Распаковка строки
func Unpacking(s string) (string, error) {
	fmt.Println("text")
	r := []rune(s)             // Получение из строки слайса rune, чтобы можно было по индексу обращаться к предыдущему символу
	var result strings.Builder // Переменная содержащая результат работы функции

	for i := 0; i < len(r); i++ { // Цикл для перебора элементов слайса. Не через range, чтобы можно было пропустить шаг.
		if r[i] > 47 && r[i] < 58 { // Если элемент число (rune от 48 до 57 - числа от 0 до 9)
			if i == 0 { // Если число это первый символ строки, то возвращается ошибка
				return "", fmt.Errorf("некорректная строка")
			}
			for j := 1; j < int(r[i])-48; j++ { // Цикл, который добавляет символ к результату N-1 раз. int(r[i])-48 - получение числа N из rune.
				result.WriteRune(r[i-1])
			}
		} else if r[i] == 92 { // Если символ равен "\"
			if i == len(r)-1 { // Если это последний символ, то возвращается ошибка
				return "", fmt.Errorf("некорректная строка")
			}
			if r[i+1] > 46 && r[i+1] < 58 || r[i+1] == 92 { // Если после символа "\" следует число или "\"
				result.WriteRune(r[i+1]) // Добавление символа
				i++                      // Пропуск 1 шага
			} else { // Если после символа не число или не "\", то возвращается ошибка
				return "", fmt.Errorf("некорректная строка")
			}
		} else { // Добавление символа к результату
			result.WriteRune(r[i])
		}
	}
	return result.String(), nil // Вывод результата
}