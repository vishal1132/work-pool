package pool

const (
	maxQueueSize     = 10000
	defaultQueueSize = 100
)

func MaxQueueSize(size uint32) *pool {
	setMaxQueueSize(size)
	return getPool()
}

func (p *pool) MaxQueueSize(size uint32) *pool {
	setMaxQueueSize(size)
	return getPool()
}

func setMaxQueueSize(size uint32) {
	p = getPool()
	if p.queueSize == 0 {
		size = defaultQueueSize
	}
	if size > maxQueueSize {
		size = maxQueueSize
	}
	p.queueSize = size
}

func getQueueSize() uint32 {
	return getPool().queueSize
}

func (p *pool) enableQueue() *pool {
	p = getPool()
	if p.queueSize == 0 {
		setMaxQueueSize(0)
	}
	p.workChan = make(chan work, getQueueSize())
	return p
}
