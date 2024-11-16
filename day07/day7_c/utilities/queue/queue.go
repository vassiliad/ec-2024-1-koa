package queue

import "container/list"

// Queue is a queue
type Queue interface {
	Front() *list.Element
	Len() int
	Add(interface{})
	Pop() *list.Element
}

type queueImpl struct {
	*list.List
}

func (q *queueImpl) Add(v interface{}) {
	q.PushFront(v)
}

func (q *queueImpl) Pop() *list.Element {
	e := q.Front()
	if e != nil {
		q.List.Remove(e)
	}

	return e
}

func New() Queue {
	return &queueImpl{list.New()}
}
