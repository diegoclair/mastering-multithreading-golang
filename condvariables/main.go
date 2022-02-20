package main

import (
	"sync"
	"time"
)

/*
	Condition variable to do a multithread synchronization

	without a condition variable this program results in a negative balance

	with a condition variable, when the condition of balance on spendy is < than 0, then it will wait until it receives a signal
	So, each time of the stingy put some money in balance, then it will send a signal to a specific thread that are waiting to check the condition again
*/
var (
	money          = 100
	lock           = sync.Mutex{}
	moneyDeposited = sync.NewCond(&lock)
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 10
		println("Stingy sees balance of", money)
		moneyDeposited.Signal()
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Stingy Done")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		for money-20 < 0 {
			moneyDeposited.Wait()
		}
		money -= 20
		println("Spendy sees balance of", money)
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Spendy Done")
}

func main() {
	go stingy()
	go spendy()
	time.Sleep(3000 * time.Millisecond)
	print(money)
}
