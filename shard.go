package gocron

import (
	"container/list"
	"sync"
	"sync/atomic"
)

const (
	// Maximum number of shards
	// default: 1024
	MaxShard uint16 = 1 << 10
)

type shard struct {
	contain container
}

type contain struct {
	// list is a doubly linked list,
	// used to store all running tasks
	list *list.List

	//
	count atomic.Value

	lock sync.RWMutex
}

type container interface {
	getShardElement()
	addShardElement(s *scheduler)
}

func newShard() *shard {
	return &shard{
		contain: newContain(),
	}
}

func newContain() container {
	return &contain{
		list:  list.New(),
		count: atomic.Value{},
	}
}

func (c *contain) getShardElement() {

}

func (c *contain) addShardElement(s *scheduler) {
	//c.lock.Lock()
	//defer c.lock.Unlock()
	//
	//if s.weight > 0 {
	//	element := list.Element{Value: s}
	//	c.list.InsertAfter()
	//} else {
	//	element := list.Element{Value: s}
	//	c.list.PushBack(element)
	//}
}
