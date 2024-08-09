package query

import (
	"fmt"

	"github.com/michaeldcanady/servicenow-sdk-go/internal"
)

type query struct {
	list internal.LinkedList[any]
}

func Query(opts ...option) string {
	query := new()

	for _, opt := range opts {
		opt(query)
	}
	return query.String()
}

func new() *query {
	return &query{
		list: internal.NewLinkedList[any](),
	}
}

func (q *query) AddValue(val any) {
	q.list.AddNode(internal.NewNode(val))
}

func (q *query) GetHead() any {
	return q.list.GetHead().GetValue()
}

func (q *query) GetTail() any {
	return q.list.GetTail().GetValue()
}

func (q *query) String() string {
	query := ""
	node := q.list.GetHead()
	for node != nil && node.GetValue() != nil {
		query += fmt.Sprintf("%s", node.GetValue())
		node = node.GetNext()
	}
	return query
}
