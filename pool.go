package pool

import (
	"errors"
	"sync"
	"time"
)

var once sync.Once

type workableUnit func(args ...interface{})

var (
	emptyChanError = errors.New("pool: worker queue is empty")
)

type work struct {
	f         workableUnit
	createdAt time.Time
	args      []interface{}
}

type pool struct {
	// uint8 coz the procs can't be negative and the max size of 255 seems reasonable for procs
	Procs         uint8
	MaxGoRoutines uint32
	queueSize     uint32
	workChan      chan work
	wg            sync.WaitGroup
	ttl           time.Duration
}

var p *pool

func GetPool() *pool {
	once.Do(func() {
		p = &pool{}
	})
	return p
}

func getPool() *pool {
	if p == nil {
		p = GetPool()
	}
	return p
}

func Start() {
	p := getPool()
	p.enableMaxProcs()
	p.enableQueue()
	p.enableMaxGoRoutines()
}

func Submit(f func(args ...interface{}), args ...interface{}) error {
	p := getPool()
	return p.Submit(f, args)
}

func Stop() {
	p := getPool()
	p.Stop()
}

func (p *pool) Stop() {
	close(p.workChan)
	p.wg.Wait()
}

func (p *pool) Submit(f workableUnit, args []interface{}) error {
	if p.workChan == nil {
		return emptyChanError
	}
	p.workChan <- work{f, time.Now(), args}
	return nil
}
