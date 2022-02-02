package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	ID       int
	Position Vector2D
	Velocity Vector2D
}

func (b *Boid) moveOne() {
	b.Position = b.Position.Add(b.Velocity)
	nextPosition := b.Position.Add(b.Velocity)
	if nextPosition.X >= screenWidth || nextPosition.X < 0 {
		b.Velocity = Vector2D{X: -b.Velocity.X, Y: b.Velocity.Y}
	}

	if nextPosition.Y >= screenHeight || nextPosition.Y < 0 {
		b.Velocity = Vector2D{X: b.Velocity.X, Y: -b.Velocity.Y}
	}
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(boidID int) {
	b := Boid{
		Position: Vector2D{X: rand.Float64() * screenWidth, Y: rand.Float64() * screenHeight},
		Velocity: Vector2D{X: (rand.Float64() * 2) - 1.0, Y: (rand.Float64() * 2) - 1.0},
		ID:       boidID,
	}
	boids[boidID] = &b
	go b.start()
}
