package query

import ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

type Uint interface {
	~uint | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type Int interface {
	~int | ~int32 | ~int64
}

type Numeric interface {
	Int | Float | Uint
}

type QueryBuilder struct {
	query           ast.Node
	logicalOperator ast.Operator
}

func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		query:           nil,
		logicalOperator: ast.Operator("^"),
	}
}

func (qB *QueryBuilder) AddFilter(field string, filter func(string) ast.Node) *QueryBuilder {
	newExpr := filter(field)
	if qB.query != nil {
		newExpr = &ast.BinaryNode{
			LeftExpression:  qB.query,
			Operator:        qB.logicalOperator,
			Position:        qB.query.Right(),
			RightExpression: newExpr,
		}
	}
	qB.query = newExpr
	return qB
}

func (qB *QueryBuilder) group(op ast.Operator, groupFunc func(q *QueryBuilder)) *QueryBuilder {
	subBuilder := &QueryBuilder{
		query:           nil,
		logicalOperator: op,
	}
	groupFunc(subBuilder)

	if subBuilder.query == nil {
		return qB
	}

	newQuery := subBuilder.query

	if qB.query != nil {
		newQuery = &ast.BinaryNode{
			LeftExpression:  qB.query,
			Operator:        qB.logicalOperator,
			Position:        qB.query.Right(),
			RightExpression: subBuilder.query,
		}
	}

	qB.query = newQuery

	return qB
}

func (qB *QueryBuilder) OrGroup(group func(q *QueryBuilder)) *QueryBuilder {
	return qB.group("^OR", group)
}

func (qB *QueryBuilder) AndGroup(group func(q *QueryBuilder)) *QueryBuilder {
	return qB.group("^", group)
}

func (qB *QueryBuilder) Build() string {
	visitor := NewStringerVisitor()
	visitor.Visit(qB.query)

	return visitor.String()
}
