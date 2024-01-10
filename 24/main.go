package main

import (
	"fmt"
	"math"
)

// Point Структура представляет точку с координатами x и y
type Point struct {
	x, y float64
}

// NewPoint Конструктор
func NewPoint(x, y float64) Point {
	return Point{x, y}
}

// DistanceTo Метод для вычисления расстояния между двумя точками
func (p Point) DistanceTo(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)

	// Вычисляем расстояние между точками
	distance := point1.DistanceTo(point2)

	// Выводим результат
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
