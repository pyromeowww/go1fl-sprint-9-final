package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) ([]int, error) {
	if size < 1 { // проверка на корректность входных данных, что size больше 0
		return nil, fmt.Errorf("size must be greater than 0")
	}
	numbers := make([]int, size) // слайс цифр

	r := rand.New(rand.NewSource(time.Now().UnixNano())) // генератор случайных чисел
	for i := 0; i < size; i++ {
		numbers[i] = r.Intn(1000) // генерируем случайное число и добавляем его в слайс
	}
	return numbers, nil
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int, error) {
	if len(data) == 0 { // проверка на пустоту слайса
		return 0, fmt.Errorf("slice is empty")
	}

	maxNumber := data[0]              // первый элемент как начальный максимум
	for _, number := range data[1:] { // начинаем со второго элемента
		if number > maxNumber {
			maxNumber = number // если текущий больше максимального, то обновляем максимум
		}
	}
	return maxNumber, nil
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) (int, error) {
	if len(data) == 0 { // проверка на пустоту слайса
		return 0, fmt.Errorf("slice is empty")
	}

	var wg sync.WaitGroup
	maxNumber := make([]int, CHUNKS) // максимальные числа в каждом куске
	for i := range maxNumber {
		maxNumber[i] = math.MinInt32 // Инициализируем каждый элемент минимальным возможным значением int32
	}

	chunkSize := (len(data) + CHUNKS - 1) / CHUNKS // размер каждого куска

	for i := 0; i < CHUNKS; i++ { // проходимся по каждому куску
		start := i * chunkSize // старт каждого куска
		if start >= len(data) {
			continue // не запускаем горутину для пустого чанка
		}
		wg.Add(1)
		go func(ind int) {
			defer wg.Done()
			start := ind * chunkSize // старт каждого куска
			end := start + chunkSize // конец каждого куска
			if end > len(data) {     // если конец больше длины слайса, то устанавливаем конец на длину слайса
				end = len(data)
			}
			localMax := data[start]                 // максимальное число в каждом куске
			for _, v := range data[start+1 : end] { // проходимся по каждому числу в каждом куске
				if v > localMax { // если текущее число больше максимального, то обновляем максимум
					localMax = v
				}
			}
			maxNumber[ind] = localMax // максимальное число в каждом куске
		}(i)
	}
	wg.Wait()

	overallMax := maxNumber[0]
	for _, v := range maxNumber[1:] { // проходимся по каждому максимальному числу в каждом куске
		if v > overallMax { // если текущее число больше максимального, то обновляем максимум
			overallMax = v
		}
	}
	return overallMax, nil
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numbers, err := generateRandomElements(SIZE) // Генерируем 1000 целых чисел
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()          // замеряем время поиска
	max, err := maximum(numbers) // ищем максимальное значение в один поток
	if err != nil {
		fmt.Println(err)
		return
	}
	elapsed := time.Since(start).Milliseconds() // замеряем время поиска

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()            // замеряем время поиска
	max, err = maxChunks(numbers) // ищем максимальное значение в CHUNKS потоков
	if err != nil {
		fmt.Println(err)
		return
	}
	elapsed = time.Since(start).Milliseconds() // замеряем время поиска

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
