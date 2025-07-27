package query

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
	forest  *QueryForest
	current *QueryTree
}

func NewQueryBuilder() *QueryBuilder {
	tree := NewQueryTree()

	forest := NewQueryForest()
	forest.AddTree(tree)

	return &QueryBuilder{
		forest:  forest,
		current: tree,
	}
}

func (qB *QueryBuilder) AddFilter(field string, filter ValueWrapper) *QueryBuilder {
	qB.current.AddCondition(filter.ToCondition(field))
	return qB
}

func (qB *QueryBuilder) group(logicalOperator string, group func(q *QueryBuilder)) *QueryBuilder {
	subBuilder := &QueryBuilder{
		forest: nil,
		current: &QueryTree{
			root: &LogicalGroup{
				Operator: logicalOperator,
				Children: []QueryNode{},
			},
		},
	}
	group(subBuilder)

	qB.current.AddCondition(subBuilder.current)

	return qB
}

func (qB *QueryBuilder) OrGroup(group func(q *QueryBuilder)) *QueryBuilder {
	return qB.group("^OR", group)
}

func (qB *QueryBuilder) AndGroup(group func(q *QueryBuilder)) *QueryBuilder {
	return qB.group("^And", group)
}

func (qB *QueryBuilder) Build() string {
	return qB.forest.Serialize()
}
