package deadlock

import (
	"time"

	"github.com/diegoclair/mastering-multithreading-golang/deadlocks_train/common"
)

func MoveTrain(train *common.Train, distance int, crossings []*common.Crossing) {
	for train.PositionFront < distance {
		train.PositionFront += 1
		for _, crossing := range crossings {
			if train.PositionFront == crossing.Position {
				crossing.Intersection.Mutex.Lock()
				crossing.Intersection.LockedBy = train.Id
			}
			back := train.PositionFront - train.TrainLength
			if back == crossing.Position {
				crossing.Intersection.LockedBy = -1
				crossing.Intersection.Mutex.Unlock()
			}
		}
		time.Sleep(30 * time.Millisecond)
	}
}
