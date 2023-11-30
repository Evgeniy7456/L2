package main

import (
	"testing"
)

func TestUnpacking(t *testing.T) {
	t.Run("Test 1", func(t *testing.T) {
		s := "a4bc2d5e"
		result := "aaaabccddddde"

		realResult, err := Unpacking(s)
		if err != nil {
			t.Errorf("error != nil: %v", err)
		}
		if realResult != result {
			t.Errorf("expected result: %s != %s real result", result, realResult)
		}
	})

	t.Run("Test 2", func(t *testing.T) {
		s := "abcd"
		result := "abcd"

		realResult, err := Unpacking(s)
		if err != nil {
			t.Errorf("error != nil: %v", err)
		}
		if realResult != result {
			t.Errorf("expected result: %s != %s real result", result, realResult)
		}
	})

	t.Run("Test 3", func(t *testing.T) {
		s := "45"
		result := ""

		realResult, err := Unpacking(s)
		if err == nil {
			t.Errorf("expected error: некорректная строка")
		}
		if realResult != result {
			t.Errorf("expected result: %s != %s real result", result, realResult)
		}
	})

	t.Run("Test 4", func(t *testing.T) {
		s := ""
		result := ""

		realResult, err := Unpacking(s)
		if err != nil {
			t.Errorf("error != nil: %v", err)
		}
		if realResult != result {
			t.Errorf("expected result: %s != %s real result", result, realResult)
		}
	})

	t.Run("Test 5", func(t *testing.T) {
		s := `qwe\4\5`
		result := "qwe45"

		realResult, err := Unpacking(s)
		if err != nil {
			t.Errorf("error != nil: %v", err)
		}
		if realResult != result {
			t.Errorf("expected result: %s != %s real result", result, realResult)
		}
	})

	t.Run("Test 6", func(t *testing.T) {
		s := `qwe\45`
		result := "qwe44444"

		realResult, err := Unpacking(s)
		if err != nil {
			t.Errorf("error != nil: %v", err)
		}
		if realResult != result {
			t.Errorf("expected result: %s != %s real result", result, realResult)
		}
	})

	t.Run("Test 7", func(t *testing.T) {
		s := `qwe\\5`
		result := `qwe\\\\\`

		realResult, err := Unpacking(s)
		if err != nil {
			t.Errorf("error != nil: %v", err)
		}
		if realResult != result {
			t.Errorf("expected result: %s != %s real result", result, realResult)
		}
	})
}
