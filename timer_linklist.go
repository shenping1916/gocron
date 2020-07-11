package gocron

import (
	"sync/atomic"
	"unsafe"
)

type ThreadSafeList struct {
	head *listNode
}

type listNode struct {
	markableNext *markablePointer
	object       unsafe.Pointer
}

type markablePointer struct {
	marked bool
	next   *listNode
}

func newLinkList() ThreadSafeList {
	return new(ThreadSafeList).init()
}

func (t *ThreadSafeList) InsertObject(object unsafe.Pointer, lessThanFn func(unsafe.Pointer, unsafe.Pointer) bool) {
	currentHeadAddress := &t.head
	currentHead := t.head

	if currentHead == nil || lessThanFn(object, currentHead.object) {
		newNode := listNode{
			object: object,
			markableNext: &markablePointer{
				next: currentHead,
			},
		}

		compareSuccess := atomic.CompareAndSwapPointer(
			(*unsafe.Pointer)(unsafe.Pointer(currentHeadAddress)),
			unsafe.Pointer(currentHead),
			unsafe.Pointer(&newNode),
		)

		if !compareSuccess {
			t.InsertObject(object, lessThanFn)
		}

		return
	}

	cursor := t.head
	for {
		if cursor.markableNext.next == nil || lessThanFn(object, currentHead.markableNext.next.object) {
			cursor = cursor.markableNext.next
		}
	}
}

func (t *ThreadSafeList) DeleteObject(object unsafe.Pointer) {
	var previous *listNode
	currentHeadAddress := &t.head
	currentHead := t.head

	cursor := currentHead
	for {
		if cursor == nil {
			break
		}

		if cursor.object == object {
			nextNode := cursor.markableNext.next
			
			if previous != nil {

			} else {

			}
		}

		previous = cursor
		cursor = cursor.markableNext.next
	}
}

func (t *ThreadSafeList) Iter() <-chan unsafe.Pointer {
	ch := make(chan unsafe.Pointer)
	go func() {
		cursor := t.head
		for {
			if cursor == nil {
				break
			}

			ch <- cursor.object
			cursor = cursor.markableNext.next
		}

		close(ch)
	}()
	return ch
}

func (t *ThreadSafeList) init() ThreadSafeList {
	return ThreadSafeList{}
}
