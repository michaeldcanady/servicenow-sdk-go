//go:build preview.query

package query

type logicalConditionBuilder interface {
	And() *FieldBuilder
	Or() *FieldBuilder
}
