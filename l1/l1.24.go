package l1

import "math"

// Разработать программу нахождения расстояния между двумя точками на плоскости.
// Точки представлены в виде структуры Point с инкапсулированными (приватными) полями x, y (типа float64) и конструктором.
// Расстояние рассчитывается по формуле между координатами двух точек.

// Подсказка: используйте функцию-конструктор NewPoint(x, y), Point и метод Distance(other Point) float64.

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(other Point) float64 {
	dX := p.x - other.x
	dY := p.y - other.y
	return math.Sqrt(dX * dX + dY * dY)
}
