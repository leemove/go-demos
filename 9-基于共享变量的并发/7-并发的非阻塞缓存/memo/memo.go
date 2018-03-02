package memo

import (
	"fmt"
	"sync"
)

type entry struct {
	res   result
	ready chan struct{}
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		j := <-e.ready
		fmt.Println(j)
	}
	return e.res.value, e.res.err
	// memo.mu.Lock()
	// res, ok := memo.cache[key]
	// if !ok {
	// 	res.value, res.err = memo.f(key)
	// 	memo.cache[key] = res
	// }
	// memo.mu.Unlock()
	// return res.value, res.err
}
