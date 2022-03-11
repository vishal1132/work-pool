package pool

import "runtime"

// Default Procs is 1

const (
	defaultProcs = 1
)

func (p *pool) enableMaxProcs() *pool {
	if p == nil {
		p = getPool()
	}
	runtime.GOMAXPROCS(int(p.Procs))
	return p
}

func SetMaxProcs(procs uint8) *pool {
	p := getPool()
	p.setMaxProcs(procs)
	return p
}

func (p *pool) SetMaxProcs(procs uint8) *pool {
	p = getPool()
	p.setMaxProcs(procs)
	return p
}

// setMaxProcs sets the max procs to the given value if the given value lies between (0,runtime.NumCPU()]
// if the given value is more than runtime.NumCPU(), it sets the max procs to the number of CPUs
// if the given value is 0 it will set the procs to defaultProcs value
func (p *pool) setMaxProcs(procs uint8) {
	if procs == 0 {
		procs = defaultProcs
	}
	if procs > uint8(runtime.NumCPU()) {
		procs = uint8(runtime.NumCPU())
	}

	p.Procs = procs
}

func (p *pool) getProcs() uint8 {
	if p.Procs == 0 {
		p.setMaxProcs(0)
	}
	return p.Procs
}

func GetProcs() uint8 {
	return getPool().getProcs()
}
