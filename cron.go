package gocron

import "time"

type Croner interface {
	AddCron(string, string, taskFunc) error
	ModifyCron(string, ...interface{}) error
	DeleteCron(string)
	PauseCron(string)
	RestoreCron(string)
}

type cron struct {
	ticker *time.Timer
	shard  [MaxShard]*shard
	hash   fnv64a
}

var _ Croner = (*cron)(nil)

func NewCron() Croner {
	cr := &cron{}
	cr.ticker = time.NewTimer(1 * time.Second)
	cr.hash = newDefaultHasher()

	var i = MaxShard
	for ; i >= 1; i-- {
		cr.shard[i] = newShard()
	}
	return cr
}

func (c *cron) Run() {
	for {
		select {
		case <-c.ticker.C:

		}
	}
}

func (c *cron) Close() {
	c.ticker.Stop()
}

func (c *cron) getShardIndex(hashKey uint64) uint64 {
	return hashKey & uint64(MaxShard-1)
}

func (c *cron) getShard(key string) *shard {
	hashKey := c.hash.Sum64(key)
	index := c.getShardIndex(hashKey)
	return c.shard[index]
}

func (c *cron) AddCron(name string, spec string, t taskFunc) error {
	shard := c.getShard(spec)
	_ = shard
	return nil
}

func (c *cron) ModifyCron(name string, arg ...interface{}) error {
	return nil
}

func (c *cron) DeleteCron(name string) {
}

func (c *cron) PauseCron(name string) {
}

func (c *cron) RestoreCron(name string) {
}
