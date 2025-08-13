package query

import ast "github.com/michaeldcanady/servicenow-sdk-go/internal/ast"

type QueryBuilder struct {
	query           ast.Node
	logicalOperator ast.Operator
}

// NewQueryBuilder Instantiates new *QueryBuilder
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{
		query:           nil,
		logicalOperator: ast.OperatorAnd,
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
	return qB.group(ast.OperatorOr, group)
}

func (qB *QueryBuilder) AndGroup(group func(q *QueryBuilder)) *QueryBuilder {
	return qB.group(ast.OperatorAnd, group)
}

func (qB *QueryBuilder) String() string {
	visitor := NewStringerVisitor()
	visitor.Visit(qB.query)

	return visitor.String()
}
