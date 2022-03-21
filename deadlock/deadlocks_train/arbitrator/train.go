package arbitrator

import (
	"sync"
	"time"

	"github.com/diegoclair/mastering-multithreading-golang/deadlocks_train/common"
)

var (
	controller = sync.Mutex{}
	cond       = sync.NewCond(&controller)
)

func allFree(intersectionsToLock []*common.Intersection) bool {
	for _, it := range intersectionsToLock {
		if it.LockedBy >= 0 {
			return false
		}
	}
	return true
}

func lockIntersectionsInDistance(id, reserveStart int, reserveEnd int, crossings []*common.Crossing) {
	var intersectionsToLock []*common.Intersection
	for _, crossing := range crossings {
		if reserveEnd >= crossing.Position &&
			reserveStart <= crossing.Position &&
			crossing.Intersection.LockedBy != id {
			intersectionsToLock = append(intersectionsToLock, crossing.Intersection)
		}
	}
	controller.Lock()
	for !allFree(intersectionsToLock) {
		cond.Wait() //this will put the thread to sleep until it receive a signal. When it receives a signal it will check again fot the allFree function
	}
	for _, it := range intersectionsToLock {
		it.LockedBy = id
		time.Sleep(10 * time.Millisecond)
	}
	controller.Unlock()
}

func MoveTrain(train *common.Train, distance int, crossings []*common.Crossing) {
	for train.PositionFront < distance {
		train.PositionFront += 1
		for _, crossing := range crossings {
			if train.PositionFront == crossing.Position {
				lockIntersectionsInDistance(train.Id, crossing.Position, crossing.Position+train.TrainLength, crossings)
			}
			back := train.PositionFront - train.TrainLength
			if back == crossing.Position {
				controller.Lock()
				crossing.Intersection.LockedBy = -1
				cond.Broadcast() //it with send a signal to all threads that are on cond.Wait() process
				controller.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
