package main

import (
	"fmt"
	"math"
)

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
//
//
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func groupTemperatures(temperatures []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, temp := range temperatures {
		// Округляем температуру вниз до ближайшего кратного 10
		key := int(math.Floor(temp/10) * 10)

		// Добавляем температуру в соответствующую группу
		groups[key] = append(groups[key], temp)
	}

	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Группируем температуры
	groupedTemps := groupTemperatures(temperatures)

	// Выводим результаты
	for key, values := range groupedTemps {
		fmt.Printf("%d: %v\n", key, values)
	}
}
