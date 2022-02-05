package main

import "math"

type Vector2D struct {
	X float64
	Y float64
}

func (v Vector2D) Add(v2 Vector2D) Vector2D {
	return Vector2D{X: v.X + v2.X, Y: v.Y + v2.Y}
}

func (v Vector2D) Subtract(v2 Vector2D) Vector2D {
	return Vector2D{X: v.X - v2.X, Y: v.Y - v2.Y}
}

func (v Vector2D) Multiply(v2 Vector2D) Vector2D {
	return Vector2D{X: v.X * v2.X, Y: v.Y * v2.Y}
}

func (v Vector2D) AddBoth(value float64) Vector2D {
	return Vector2D{X: v.X + value, Y: v.Y + value}
}

func (v Vector2D) SubtractBoth(value float64) Vector2D {
	return Vector2D{X: v.X - value, Y: v.Y - value}
}

func (v Vector2D) MultiplyBoth(value float64) Vector2D {
	return Vector2D{X: v.X * value, Y: v.Y * value}
}

func (v Vector2D) DivisionBoth(value float64) Vector2D {
	return Vector2D{X: v.X / value, Y: v.Y / value}
}

func (v Vector2D) limit(lower, upper float64) Vector2D {
	return Vector2D{X: math.Min(math.Max(v.X, lower), upper), Y: math.Min(math.Max(v.Y, lower), upper)}
}

func (v Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v.X-v2.X, 2) + math.Pow(v.Y-v2.Y, 2))
}
