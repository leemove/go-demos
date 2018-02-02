package main

import (
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func lookUp(key string) (v string) {
	cache.Lock()
	v = cache.mapping[key]
	cache.Unlock()
	return
}

func main() {

}
