package main

import "sync"

//sync.Mutex和sync.RWMutex

func testLock() {
	var l sync.Mutex
	l.Lock()
	defer l.Unlock()

	//business logic
}
