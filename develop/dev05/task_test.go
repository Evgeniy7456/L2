package main

import (
	"testing"
)

func TestTextSearch(t *testing.T) {
	t.Run("search", func(t *testing.T) {
		lines := []string{"text text text", "text find text", "text find", "text text text"}
		text := "find"
		result := map[int]struct{}{
			1: {},
			2: {},
		}

		realResult := TextSearch(lines, text)

		if len(realResult) != len(result) {
			t.Error("expected result != real result")
		}
		for item := range result {
			if _, ok := realResult[item]; !ok {
				t.Error("expected result != real result")
			}
		}
	})

	t.Run("flag i", func(t *testing.T) {
		i = true

		lines := []string{"text text text", "text FIND text", "text find", "text text text"}
		text := "Find"
		result := map[int]struct{}{
			1: {},
			2: {},
		}

		realResult := TextSearch(lines, text)

		if len(realResult) != len(result) {
			t.Error("expected result != real result")
		}
		for item := range result {
			if _, ok := realResult[item]; !ok {
				t.Error("expected result != real result")
			}
		}

		i = false
	})

	t.Run("flag F", func(t *testing.T) {
		F = true

		lines := []string{"text text text", "text find text", "text find", "text text text", "find"}
		text := "find"
		result := map[int]struct{}{
			4: {},
		}

		realResult := TextSearch(lines, text)

		if len(realResult) != len(result) {
			t.Error("expected result != real result")
		}
		for item := range result {
			if _, ok := realResult[item]; !ok {
				t.Error("expected result != real result")
			}
		}

		F = false
	})
}
