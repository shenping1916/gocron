package gocron

type array []*scheduler

func (a array) insert(s *scheduler) {

}

type linkList struct {
	prev *scheduler
	next *scheduler

	ga []*array
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
