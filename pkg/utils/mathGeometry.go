package utils

import (
	"math"
)

type Point struct {
	X float64
	Y float64
	R float64
}

//here are all the auxiliary function to calculate trilateration for the ship location
func Square(v float64) float64 {
	return v * v
}

func Normalize(p Point) float64 {
	return math.Sqrt(Square(p.X) + Square(p.Y))
}

func Dot(p1, p2 Point) float64 {
	return p1.X*p2.X + p1.Y*p2.Y
}

func Subtract(p1, p2 Point) Point {
	return Point{
		X: p1.X - p2.X,
		Y: p1.Y - p2.Y,
	}
}

func Add(p1, p2 Point) Point {
	return Point{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}

func Divide(p Point, v float64) Point {
	return Point{
		X: p.X / v,
		Y: p.Y / v,
	}
}

func Multiply(p Point, v float64) Point {
	return Point{
		X: p.X * v,
		Y: p.Y * v,
	}
}

func RoundUp(input Point, places int) Point {
	pow := math.Pow(10, float64(places))
	input.X = pow * input.X
	input.Y = pow * input.Y
	input.X = (math.Ceil(input.X)) / pow
	input.Y = (math.Ceil(input.Y)) / pow
	return input
}
