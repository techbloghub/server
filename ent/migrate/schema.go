// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CompaniesColumns holds the columns for the "companies" table.
	CompaniesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString},
		{Name: "logo_url", Type: field.TypeString},
		{Name: "blog_url", Type: field.TypeString},
		{Name: "rss_url", Type: field.TypeString},
	}
	// CompaniesTable holds the schema information for the "companies" table.
	CompaniesTable = &schema.Table{
		Name:       "companies",
		Columns:    CompaniesColumns,
		PrimaryKey: []*schema.Column{CompaniesColumns[0]},
	}
	// PostingsColumns holds the columns for the "postings" table.
	PostingsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString},
		{Name: "url", Type: field.TypeString, Unique: true},
		{Name: "published_at", Type: field.TypeTime},
		{Name: "tags", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"postgres": "text[]"}},
		{Name: "company_postings", Type: field.TypeInt, Nullable: true},
	}
	// PostingsTable holds the schema information for the "postings" table.
	PostingsTable = &schema.Table{
		Name:       "postings",
		Columns:    PostingsColumns,
		PrimaryKey: []*schema.Column{PostingsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "postings_companies_postings",
				Columns:    []*schema.Column{PostingsColumns[7]},
				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "delete_time", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CompaniesTable,
		PostingsTable,
		TagsTable,
	}
)

func init() {
	PostingsTable.ForeignKeys[0].RefTable = CompaniesTable
}
