package main

import "sync"

type SafeMap struct {
	rw sync.RWMutex
	Map map[int]int
}

func main() {
	sMap := newSafeMap2()
	for i := 0; i < 10000; i++ {
		go sMap.sWriteMap(i, i)
		go sMap.sReadMap(i)
	}
}

func newSafeMap2() *SafeMap{
	return &SafeMap{
		rw: sync.RWMutex{},
		Map: make(map[int]int),
	}
}

func (sMap *SafeMap) sReadMap(i int ) int {
	sMap.rw.RLock()
	defer sMap.rw.RUnlock()
	res := sMap.Map[i]
	return res
}

func (sMap *SafeMap)  sWriteMap(i int, j int){
	sMap.rw.Lock()
	defer sMap.rw.Unlock()
	sMap.Map[i] = j
}
