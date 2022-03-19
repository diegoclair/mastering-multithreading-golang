package common

import "sync"

type Train struct {
	Id            int
	TrainLength   int
	PositionFront int
}

type Intersection struct {
	Id       int
	Mutex    sync.Mutex
	LockedBy int
}

type Crossing struct {
	Position     int
	Intersection *Intersection
}
