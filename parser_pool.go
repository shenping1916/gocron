package gocron

import "sync"

type parserPool struct {
	pool *sync.Pool
}

func (p parserPool) initPool() parserPool {
	p.pool = &sync.Pool{New: func() interface{} {
		return newParser()
	}}

	return p
}

func (p parserPool) getParser() *parser {
	return p.pool.Get().(*parser)
}

func (p parserPool) putParser()  {
	p.pool.Put(p.getParser())
}


