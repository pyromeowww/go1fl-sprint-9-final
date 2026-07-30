package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements_Success(t *testing.T) { // тест на успешную генерацию случайных чисел
	numbers, err := generateRandomElements(10) // генерируем 10 случайных чисел
	assert.NoError(t, err)                     // проверяем, что нет ошибок
	assert.Len(t, numbers, 10)                 // проверяем, что количество чисел равно 10
}

func TestGenerateRandomElements_InvalidSize(t *testing.T) { // тест на некорректный размер
	testCases := []struct { // тестовые случаи
		size     int
		expected string
	}{
		{0, "size must be greater than 0"},  // тестовый случай с size=0
		{-1, "size must be greater than 0"}, // тестовый случай с size=-1
		{-5, "size must be greater than 0"}, // тестовый случай с size=-5
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("size %d", tc.size), func(t *testing.T) {
			numbers, err := generateRandomElements(tc.size)
			assert.Error(t, err, "expected error for size %d", tc.size)                                                         // проверяем, что ошибка есть
			assert.Nil(t, numbers, "expected nil for numbers for size %d", tc.size)                                             // проверяем, что numbers равен nil
			assert.Contains(t, err.Error(), tc.expected, "error message should contain '%s' for size %d", tc.expected, tc.size) // проверяем, что ошибка содержит expected
		})
	}
}

func TestGenerateRandomElements_ValueRange(t *testing.T) {
	numbers, err := generateRandomElements(10000)
	assert.NoError(t, err)
	for _, num := range numbers {
		assert.GreaterOrEqual(t, num, 0)
		assert.Less(t, num, 1000)
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct { // тестовые случаи
		name    string
		data    []int
		want    int
		wantErr string
	}{
		{name: "empty slice", data: []int{}, wantErr: "slice is empty"},             // тестовый случай с пустым слайсом
		{name: "single element", data: []int{1}, want: 1},                           // тестовый случай с одним числом
		{name: "multiple elements", data: []int{1, 2, 3}, want: 3},                  // тестовый случай с несколькими числами
		{name: "first element is maximum", data: []int{3, 2, 1}, want: 3},           // тестовый случай с максимальным числом в начале
		{name: "middle element is maximum", data: []int{1, 3, 2}, want: 3},          // тестовый случай с максимальным числом в середине
		{name: "last element is maximum", data: []int{1, 2, 3}, want: 3},            // тестовый случай с максимальным числом в конце
		{name: "all elements are equal", data: []int{1, 1, 1, 1, 1, 1, 1}, want: 1}, // тестовый случай с равными числами
		{name: "negative numbers", data: []int{-10, -5, -20, -1, -30}, want: -1},    // тестовый случай с отрицательными числами
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // запускаем каждый тестовый случай как отдельный подтест
			result, err := maximum(tt.data)
			if tt.wantErr != "" {
				assert.Error(t, err)                        // проверяем, что ошибка есть
				assert.Contains(t, err.Error(), tt.wantErr) // проверяем, что текст ошибки содержит подстроку "slice is empty"
				assert.Zero(t, result)                      // проверяем, что результат равен 0
			} else {
				assert.NoError(t, err)           // проверяем, что нет ошибки
				assert.Equal(t, tt.want, result) // проверяем правильный результат
			}
		})
	}
}

func TestMaximumChunks(t *testing.T) {
	tests := []struct { // тестовые случаи
		name    string
		data    []int
		want    int
		wantErr string
	}{
		{name: "empty slice", data: []int{}, wantErr: "slice is empty"},    // тестовый случай с пустым слайсом
		{name: "single element", data: []int{1}, want: 1},                  // тестовый случай с одним числом
		{name: "multiple elements", data: []int{10, 2, 14}, want: 14},      // тестовый случай с несколькими числами
		{name: "first element is maximum", data: []int{3, 2, 1}, want: 3},  // тестовый случай с максимальным числом в начале
		{name: "middle element is maximum", data: []int{1, 3, 2}, want: 3}, // тестовый случай с максимальным числом в середине
		{name: "last element is maximum", data: []int{1, 2, 3}, want: 3},   // тестовый случай с максимальным числом в конце
		{name: "less than chunks", data: []int{1, 10}, want: 10},           // тестовый случай с меньшим количеством чисел
		{name: "more than chunks", data: []int{10, 21, 1, 2, 3, 5, 10, 11, 96, 100, 21, 45, 724, 97, 12, 54, 62, 67, 67, 12, 11, 800, 11, 7, 10, 21, 1, 2, 3, 5, 10}, want: 800}, // тестовый случай с большим количеством чисел
		{name: "all elements are equal", data: []int{1, 1, 1, 1, 1, 1, 1}, want: 1},                                                                                              // тестовый случай с равными числами
		{name: "negative numbers", data: []int{-10, -5, -20, -1, -30}, want: -1},                                                                                                 // тестовый случай с отрицательными числами. 0 потому что при отрицательном значение возвращается 0
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // запускаем каждый тестовый случай как отдельный подтест
			result, err := maxChunks(tt.data)
			if tt.wantErr != "" {
				assert.Error(t, err)                        // проверяем, что ошибка есть
				assert.Contains(t, err.Error(), tt.wantErr) // проверяем, что текст ошибки содержит подстроку "slice is empty"
				assert.Zero(t, result)                      // проверяем, что результат равен 0
			} else {
				assert.NoError(t, err)           // проверяем, что нет ошибки
				assert.Equal(t, tt.want, result) // проверяем правильный результат
			}
		})
	}
}
