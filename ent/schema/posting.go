package schema

import (
	"net/url"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/techbloghub/server/internal/schemasupport"
)

// Posting holds the schema definition for the Posting entity.
type Posting struct {
	ent.Schema
}

// Fields of the Posting.
func (Posting) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("url").
			GoType(&url.URL{}).
			ValueScanner(field.BinaryValueScanner[*url.URL]{}).
			Unique(),
		field.Time("published_at"),
		field.Other("tags", &schemasupport.PostingTags{}).
			SchemaType(map[string]string{
				dialect.Postgres: "text[]",
			}).
			Optional(),
	}
}

// Edges of the Posting.
func (Posting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).
			Ref("postings").
			Unique(),
	}
}
