package pool

import "time"

const (
	maxTTL     = 1000 * time.Second
	defaultTTL = 10 * time.Second
)

func MaxTTL(ttl int64) *pool {
	setMaxTTL(ttl)
	return getPool()
}

func (p *pool) MaxTTL(ttl int64) *pool {
	setMaxTTL(ttl)
	return getPool()
}

func setMaxTTL(ttl int64) {
	if ttl <= 0 {
		ttl = int64(defaultTTL)
	}
	if ttl > int64(maxTTL) {
		ttl = int64(maxTTL)
	}
	getPool().ttl = time.Duration(ttl) * time.Second
}

func getTTL() time.Duration {
	return getPool().ttl
}
