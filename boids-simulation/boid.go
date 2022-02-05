package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	ID       int
	Position Vector2D
	Velocity Vector2D
}

func (b *Boid) calcAcceleration() Vector2D {
	upper, lower := b.Position.AddBoth(viewRadius), b.Position.AddBoth(-viewRadius)
	avgVelocity := Vector2D{0, 0}

	count := 0.0

	//sometimes the "lower X or Y" may is negative
	//and the "upper X or Y" can be more than screenWidth
	for i := math.Max(lower.X, 0); i <= math.Min(upper.X, screenWidth); i++ {
		for j := math.Max(lower.Y, 0); j <= math.Min(upper.Y, screenHeight); j++ {
			if otherBoidID := boidMap[int(i)][int(j)]; otherBoidID != -1 && otherBoidID != b.ID {
				if distance := boids[otherBoidID].Position.Distance(b.Position); distance < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(boids[otherBoidID].Velocity)
				}
			}
		}
	}

	accel := Vector2D{0, 0}
	if count > 0 {
		avgVelocity = avgVelocity.DivisionBoth(count)
		accel = avgVelocity.Subtract(b.Velocity).MultiplyBoth(adjRate)
	}
	return accel
}

func (b *Boid) moveOne() {

	acceleration := b.calcAcceleration()
	b.Velocity = b.Velocity.Add(acceleration).limit(-1, 1)

	boidMap[int(b.Position.X)][int(b.Position.Y)] = -1
	b.Position = b.Position.Add(b.Velocity)
	boidMap[int(b.Position.X)][int(b.Position.Y)] = b.ID

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
	boidMap[int(b.Position.X)][int(b.Position.Y)] = b.ID
	go b.start()
}
