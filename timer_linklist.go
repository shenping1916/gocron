package gocron

import "sync"

type array []*scheduler

func (a array) insert(s *scheduler) {

}

type linkList struct {
	sync.Mutex
	prev *scheduler
	next *scheduler
	ga   []*array
}

func newLinkList() *linkList {
	return new(linkList).init()
}

func (list *linkList) init() *linkList {
	return &linkList{
		prev: nil,
		next: nil,
	}
}
