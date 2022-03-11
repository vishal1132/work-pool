package pool

import (
	"sync"
	"time"
)

const (
	maxGoRoutines     = 10000
	defaultGoRoutines = 100
)

func SetMaxGoRoutines(goroutines uint32) *pool {
	p := getPool()
	p.setMaxGoRoutines(goroutines)
	return p
}

func (p *pool) SetMaxGoRoutines(goroutines uint32) *pool {
	p = getPool()
	p.setMaxGoRoutines(goroutines)
	return p
}

func (p *pool) setMaxGoRoutines(goroutines uint32) {
	if goroutines > maxGoRoutines {
		goroutines = maxGoRoutines
	}
	if goroutines == 0 {
		goroutines = defaultGoRoutines
	}
	p.MaxGoRoutines = uint32(goroutines)
}

func (p *pool) enableMaxGoRoutines() *pool {
	if p == nil {
		p = getPool()
	}
	for i := 0; i < int(p.MaxGoRoutines); i++ {
		p.wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			for w := range getPool().workChan {
				if time.Since(w.createdAt) > p.ttl {
					continue
				}
				w.f(w.args...)
			}
		}(&p.wg)
	}
	return p
}
