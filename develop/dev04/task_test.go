package main

import "testing"

func TestAnagramSearch(t *testing.T) {
	t.Run("test 1", func(t *testing.T) {
		words := []string{"пятка", "столик", "пятак", "слиток", "тяпка", "листок"}
		result := map[string][]string{
			"пятак":  {"пятка", "тяпка"},
			"листок": {"слиток", "столик"},
		}

		realResult := AnagramSearch(words)

		for key, value := range realResult {
			for index, item := range result[key] {
				if item != value[index] {
					t.Error("result != real result")
				}
			}
		}
	})

	t.Run("test 2", func(t *testing.T) {
		words := []string{"ПЯТКА", "столик", "Пятак", "слиток", "тяпка", "ЛИСТОК"}
		result := map[string][]string{
			"пятак":  {"пятка", "тяпка"},
			"листок": {"слиток", "столик"},
		}

		realResult := AnagramSearch(words)

		for key, value := range realResult {
			for index, item := range result[key] {
				if item != value[index] {
					t.Error("result != real result")
				}
			}
		}
	})

	t.Run("test 3", func(t *testing.T) {
		words := []string{"пятка", "столик", "пятак", "слиток", "тяпка", "листок", "стол"}
		result := map[string][]string{
			"пятак":  {"пятка", "тяпка"},
			"листок": {"слиток", "столик"},
		}

		realResult := AnagramSearch(words)

		for key, value := range realResult {
			for index, item := range result[key] {
				if item != value[index] {
					t.Error("result != real result")
				}
			}
		}
	})
}
