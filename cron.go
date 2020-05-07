package gocron

type Croner interface {
	AddCron(string, string, taskFunc) error
	ModifyCron(string, ...interface{}) error
	DeleteCron(string)
	PauseCron(string)
	RestoreCron(string)
}

type cron struct {
	shard [MaxShard]*shard
	hash  fnv64a
}

func NewCron() Croner {
	cr := &cron{}
	cr.hash = newDefaultHasher()

	var i = MaxShard
	for ; i >= 1; i-- {
		cr.shard[i] = newShard()
	}
	return cr
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
