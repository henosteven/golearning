package main

import "sync"

type SafeMap struct {
	rw sync.RWMutex
	Map map[int]int
}

func main() {
	sMap := newSafeMap()
	for i := 0; i < 10000; i++ {
		go sWriteMap(sMap, i, i)
		go sReadMap(sMap, i)
	}
}

func newSafeMap() *SafeMap{
	return &SafeMap{
		rw: sync.RWMutex{},
		Map: make(map[int]int),
	}
}

func sReadMap(sMap *SafeMap, i int ) int {
	sMap.rw.RLock()
	defer sMap.rw.RUnlock()
	res := sMap.Map[i]
	return res
}

func sWriteMap(sMap *SafeMap, i int, j int){
	sMap.rw.Lock()
	defer sMap.rw.Unlock()
	sMap.Map[i] = j
}
