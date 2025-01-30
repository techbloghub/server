package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

func (Tag) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		SoftDeleteMixin{},
	}
}

func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
	}
}

func (Tag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Unique().
			Annotations(
				entsql.IndexTypes(map[string]string{
					"postgres": "GIN",
				}),
				entsql.OpClass("gin_trgm_ops"),
			),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return nil
}
