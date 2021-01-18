package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Dataset holds the schema definition for the Dataset entity.
type Dataset struct {
	ent.Schema
}

// Mixin of the Dataset schema.
func (Dataset) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the Dataset.
func (Dataset) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Dataset.
func (Dataset) Edges() []ent.Edge {
	return nil
}
