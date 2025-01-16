package schema

import (
	"net/url"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

func (Company) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Company) Fields() []ent.Field {
	urlField := func(name string) ent.Field {
		return field.String(name).
			GoType(&url.URL{}).
			ValueScanner(field.BinaryValueScanner[*url.URL]{})
	}

	return []ent.Field{
		field.String("name"),
		urlField("logo_url"),
		urlField("blog_url"),
		urlField("rss_url"),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return nil
}
