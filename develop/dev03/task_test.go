package main

import (
	"strings"
	"testing"
)

func getString(data []string) string {
	return strings.Join(data, "\n")
}

func TestSort(t *testing.T) {

	t.Run("string sort", func(t *testing.T) {
		t.Run("line", func(t *testing.T) {
			data := []string{"b", "a", "c"}
			dataByte := []byte(getString(data))
			result := []string{"a", "b", "c"}
			resultString := getString(result)

			realResult := Sort(dataByte)

			if realResult != resultString {
				t.Error("expected result != real result")
			}
		})

		t.Run("column", func(t *testing.T) {
			k = 2
			data := []string{"c b", "b a", "a c"}
			dataByte := []byte(getString(data))
			result := []string{"b a", "c b", "a c"}
			resultString := getString(result)

			realResult := Sort(dataByte)

			if realResult != resultString {
				t.Error("expected result != real result")
			}

			k = 0
		})
	})

	t.Run("int sort", func(t *testing.T) {
		t.Run("line", func(t *testing.T) {
			n = true
			data := []string{"2", "1", "3"}
			dataByte := []byte(getString(data))
			result := []string{"1", "2", "3"}
			resultString := getString(result)

			realResult := Sort(dataByte)

			if realResult != resultString {
				t.Error("expected result != real result")
			}
		})

		t.Run("column", func(t *testing.T) {
			k = 2
			data := []string{"a 2", "c 1", "b 3"}
			dataByte := []byte(getString(data))
			result := []string{"c 1", "a 2", "b 3"}
			resultString := getString(result)

			realResult := Sort(dataByte)

			if realResult != resultString {
				t.Error("expected result != real result")
			}

			n = false
			k = 0
		})
	})

	t.Run("float sort", func(t *testing.T) {
		t.Run("line", func(t *testing.T) {
			h = true
			data := []string{"1.2", "1.3", "1.1"}
			dataByte := []byte(getString(data))
			result := []string{"1.1", "1.2", "1.3"}
			resultString := getString(result)

			realResult := Sort(dataByte)
			if realResult != resultString {
				t.Error("expected result != real result")
			}
		})

		t.Run("column", func(t *testing.T) {
			k = 2
			data := []string{"a 1.2", "c 1.3", "b 1.1"}
			dataByte := []byte(getString(data))
			result := []string{"b 1.1", "a 1.2", "c 1.3"}
			resultString := getString(result)

			realResult := Sort(dataByte)

			if realResult != resultString {
				t.Error("expected result != real result")
			}

			h = false
			k = 0
		})
	})

	t.Run("month sort", func(t *testing.T) {
		t.Run("line", func(t *testing.T) {
			M = true
			data := []string{"february", "december", "january"}
			dataByte := []byte(getString(data))
			result := []string{"january", "february", "december"}
			resultString := getString(result)

			realResult := Sort(dataByte)

			if realResult != resultString {
				t.Error("expected result != real result")
			}
		})

		t.Run("column", func(t *testing.T) {
			k = 2
			data := []string{"3 february", "2 december", "1 january"}
			dataByte := []byte(getString(data))
			result := []string{"1 january", "3 february", "2 december"}
			resultString := getString(result)

			realResult := Sort(dataByte)

			if realResult != resultString {
				t.Error("expected result != real result")
			}

			M = false
			k = 0
		})
	})

	t.Run("no duplicates sort", func(t *testing.T) {
		u = true
		data := []string{"b", "a", "c", "a"}
		data = RemoveDuplicates(data)
		dataByte := []byte(getString(data))
		result := []string{"a", "b", "c"}
		resultString := getString(result)

		realResult := Sort(dataByte)

		if realResult != resultString {
			t.Error("expected result != real result")
		}

		u = false
	})

	t.Run("reverse sort", func(t *testing.T) {
		r = true
		data := []string{"b", "a", "c"}
		dataByte := []byte(getString(data))
		result := []string{"c", "b", "a"}
		resultString := getString(result)

		realResult := Sort(dataByte)

		if realResult != resultString {
			t.Error("expected result != real result")
		}

		r = false
	})

	t.Run("trim sort", func(t *testing.T) {
		n = true
		b = true
		data := []string{"2 ", "3  ", "1 "}
		dataByte := []byte(getString(data))
		result := []string{"1 ", "2 ", "3  "}
		resultString := getString(result)

		realResult := Sort(dataByte)

		if realResult != resultString {
			t.Error("expected result != real result")
		}

		n = false
		b = false
	})

	t.Run("sorted test", func(t *testing.T) {
		t.Run("sorted", func(t *testing.T) {
			c = true
			data := []string{"a", "b", "c"}
			dataByte := []byte(getString(data))
			resultString := "Файл отсортирован"

			realResult := Sort(dataByte)

			if realResult != resultString {
				t.Error("expected result != real result")
			}
		})

		t.Run("not sorted", func(t *testing.T) {
			data := []string{"c", "a", "b"}
			dataByte := []byte(getString(data))
			resultString := "Файл не отсортирован"

			realResult := Sort(dataByte)

			if realResult != resultString {
				t.Error("expected result != real result")
			}

			c = false
		})
	})
}
