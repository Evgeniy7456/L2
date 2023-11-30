package main

import (
	"testing"
)

func TestCut(t *testing.T) {
	t.Run("flag f", func(t *testing.T) {
		str := "a\tb\tc\nd\te\tf"
		result := "b\ne"
		f = 2
		d = "\t"

		realResult := Cut(str)

		if realResult != result {
			t.Error("expected result != real result")
		}
	})

	t.Run("flag d", func(t *testing.T) {
		str := "a b c\nd e f"
		result := "c\nf"
		f = 3
		d = " "

		realResult := Cut(str)

		if realResult != result {
			t.Error("expected result != real result")
		}

		d = "\t"
	})

	t.Run("flag s", func(t *testing.T) {
		str := "a\tb\tc\nd e f\ng\th\ti"
		result := "a\ng"
		f = 1
		s = true

		realResult := Cut(str)

		if realResult != result {
			t.Error("expected result != real result")
		}
	})
}
