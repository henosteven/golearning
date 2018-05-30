package concurrent_map

import "sync"

var SHARED_COUNT int = 32

type Cmap []Sharedmap

type Sharedmap struct {
	sync.RWMutex
	m map[string]interface{}
}

func NewCmap() Cmap {
	cmap := make([]Sharedmap, SHARED_COUNT)
	for i := 0; i < SHARED_COUNT; i++ {
		smap := Sharedmap{m: make(map[string]interface{})}
		cmap[i] = smap
	}
	return cmap
}

func (cmap Cmap) Set(key string, value interface{}) {
	choosekey := hash(key)
	shareMap := cmap[choosekey]
	shareMap.Lock()
	shareMap.m[key] = value
	shareMap.Unlock()
}

func (cmap Cmap) Get(key string) interface{} {
	choosekey := hash(key)
	shareMap := cmap[choosekey]
	shareMap.RLock()
	val, _ := shareMap.m[key]
	shareMap.RUnlock()
	return val
}

func (cmap Cmap) Keys() []interface{} {
	return nil
}

func (cmap Cmap) Values() []interface{} {
	return nil
}

func (cmap Cmap) Len() int {
	return 0
}

func hash(key string) int {
	return 0
}