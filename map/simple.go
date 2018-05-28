package main

// you will get a fatal error
// fatal error: concurrent map writes
//
//
//Maps are not safe for concurrent use:
// it's not defined what happens when you read and write to them simultaneously.
// If you need to read from and write to a map from concurrently executing goroutines,
// the accesses must be mediated by some kind of synchronization mechanism.
// One common way to protect maps is with sync.RWMutex.
func main() {
	var Map map[int]int
	Map = make(map[int]int)
	for i := 0; i < 20; i++ {
		go writeMap(Map, i, i)
		go readMap(Map, i)
	}

}

func readMap(Map map[int]int, i int ) int {
	return Map[i]
}

func writeMap(Map map[int]int, i int, j int){
	Map[i] = j
}

