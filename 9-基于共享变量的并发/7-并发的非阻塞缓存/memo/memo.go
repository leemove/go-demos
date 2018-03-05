package memo

type entry struct {
	res   result
	ready chan struct{}
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type request struct {
	value    interface{}
	response chan<- result
}

type Memo struct {
	request chan request
}

func New(f Func) *Memo {
	// return &Memo{f: f, cache: make(map[string]*entry)}
	memo := &Memo{request: make(chan request)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}
