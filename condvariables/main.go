package main

import (
	"sync"
	"time"
)

/*
	without a condition variable this program results in a negative balance
*/
var (
	money = 100
	lock  = sync.Mutex{}
)

func stingy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
		money += 10
		println("Stingy sees balance of", money)
		lock.Unlock()
		time.Sleep(1 * time.Millisecond)
	}
	println("Stingy Done")
}

func spendy() {
	for i := 1; i <= 1000; i++ {
		lock.Lock()
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
