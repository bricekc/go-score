package main

import (
	"fmt"
	"strconv"
)

type BaseShape struct {
	name string
	isPolygon bool
}

func newBaseShape(a string, b bool) BaseShape {
	return BaseShape{name: a, isPolygon: b}
}

func (base BaseShape) GetName() string { 
	return base.name
}
func (base BaseShape) IsPolygon() bool { 
	return base.isPolygon
}

type Rectangle struct {
	BaseShape
	side1 float64
	side2 float64
}

func newRectangle(name string, side1 float64, side2 float64) Rectangle {
	return Rectangle{newBaseShape(name, true), side1, side2}
}

func (rectangle Rectangle) GetArea() float64 {
	return rectangle.side1 * rectangle.side2
}

func (rectangle Rectangle) GetName() string {
	 return rectangle.BaseShape.GetName() 
	}
func (rectangle Rectangle) IsPolygon() bool { 
	return rectangle.BaseShape.IsPolygon()
}

type Square struct {
	Rectangle
}

func newSquare(name string, side float64) Square {
	return Square{newRectangle(name, side, side)}
}

func (square Square) GetArea() float64 {
	return square.side1 * square.side2
}

func (square Square) GetName() string { 
	return square.BaseShape.GetName()
}
func (square Square) IsPolygon() bool { 
	return square.BaseShape.IsPolygon()
}

type Circle struct {
	BaseShape
	radius float64
}

func newCircle(name string, radius float64) Circle {
	return Circle{newBaseShape(name, false), radius}
}

func (circle Circle) GetArea() float64 {
	return 3.141 * circle.radius * circle.radius
}

func (circle Circle) GetName() string { 
	return circle.BaseShape.GetName()
}
func (circle Circle) IsPolygon() bool { 
	return circle.BaseShape.IsPolygon()
}

type Shape interface {
	GetArea() float64
	GetName() string
	IsPolygon() bool
}

func displayShape(shape Shape) {
	polygonStatus := "pas un polygone"
	if shape.IsPolygon() {
		polygonStatus = "un polygone"
	}
	areaStr := strconv.FormatFloat(shape.GetArea(), 'f', 2, 64)
	fmt.Println(shape.GetName() + " est " + polygonStatus + ". Son aire est " + areaStr + ".")}

func main() {
	rect := newRectangle("Rectangle", 4, 5)
	square := newSquare("Carré", 4)
	circle := newCircle("Cercle", 3)

	displayShape(rect)
	displayShape(square)
	displayShape(circle)
}
