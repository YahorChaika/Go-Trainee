package main

import (
	"flag"
	"fmt"
	"log"
	"math"
)


type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	shape := flag.String("shape", "", "circle или rectangle")
	width := flag.Float64("width", 0, "Ширина")
	height := flag.Float64("height", 0, "Высота")
	radius := flag.Float64("radius", 0, "Радиус")
	flag.Parse()

	var s Shape

	if *shape == "circle" {
		if *radius <= 0 {
			log.Fatal("Ошибка: нужен положительный -radius")
		}
		s = Circle{*radius}
	} else if *shape == "rectangle" {
		if *width <= 0 || *height <= 0 {
			log.Fatal("Ошибка: нужны положительные -width и -height")
		}
		s = Rectangle{*width, *height}
	} else {
		log.Fatal("Ошибка: укажите -shape circle или rectangle")
	}

	fmt.Println(s.Area())
}
