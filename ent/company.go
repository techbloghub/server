// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/techbloghub/server/ent/company"
)

// Company is the model entity for the Company schema.
type Company struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// DeleteTime holds the value of the "delete_time" field.
	DeleteTime time.Time `json:"delete_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// LogoURL holds the value of the "logo_url" field.
	LogoURL *url.URL `json:"logo_url,omitempty"`
	// BlogURL holds the value of the "blog_url" field.
	BlogURL *url.URL `json:"blog_url,omitempty"`
	// RssURL holds the value of the "rss_url" field.
	RssURL *url.URL `json:"rss_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CompanyQuery when eager-loading is set.
	Edges        CompanyEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CompanyEdges holds the relations/edges for other nodes in the graph.
type CompanyEdges struct {
	// Postings holds the value of the postings edge.
	Postings []*Posting `json:"postings,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PostingsOrErr returns the Postings value or an error if the edge
// was not loaded in eager-loading.
func (e CompanyEdges) PostingsOrErr() ([]*Posting, error) {
	if e.loadedTypes[0] {
		return e.Postings, nil
	}
	return nil, &NotLoadedError{edge: "postings"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Company) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case company.FieldID:
			values[i] = new(sql.NullInt64)
		case company.FieldName:
			values[i] = new(sql.NullString)
		case company.FieldCreateTime, company.FieldUpdateTime, company.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		case company.FieldLogoURL:
			values[i] = company.ValueScanner.LogoURL.ScanValue()
		case company.FieldBlogURL:
			values[i] = company.ValueScanner.BlogURL.ScanValue()
		case company.FieldRssURL:
			values[i] = company.ValueScanner.RssURL.ScanValue()
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Company fields.
func (c *Company) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case company.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case company.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = value.Time
			}
		case company.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				c.UpdateTime = value.Time
			}
		case company.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				c.DeleteTime = value.Time
			}
		case company.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case company.FieldLogoURL:
			if value, err := company.ValueScanner.LogoURL.FromValue(values[i]); err != nil {
				return err
			} else {
				c.LogoURL = value
			}
		case company.FieldBlogURL:
			if value, err := company.ValueScanner.BlogURL.FromValue(values[i]); err != nil {
				return err
			} else {
				c.BlogURL = value
			}
		case company.FieldRssURL:
			if value, err := company.ValueScanner.RssURL.FromValue(values[i]); err != nil {
				return err
			} else {
				c.RssURL = value
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Company.
// This includes values selected through modifiers, order, etc.
func (c *Company) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryPostings queries the "postings" edge of the Company entity.
func (c *Company) QueryPostings() *PostingQuery {
	return NewCompanyClient(c.config).QueryPostings(c)
}

// Update returns a builder for updating this Company.
// Note that you need to call Company.Unwrap() before calling this method if this Company
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Company) Update() *CompanyUpdateOne {
	return NewCompanyClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Company entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Company) Unwrap() *Company {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Company is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Company) String() string {
	var builder strings.Builder
	builder.WriteString("Company(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete_time=")
	builder.WriteString(c.DeleteTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("logo_url=")
	builder.WriteString(fmt.Sprintf("%v", c.LogoURL))
	builder.WriteString(", ")
	builder.WriteString("blog_url=")
	builder.WriteString(fmt.Sprintf("%v", c.BlogURL))
	builder.WriteString(", ")
	builder.WriteString("rss_url=")
	builder.WriteString(fmt.Sprintf("%v", c.RssURL))
	builder.WriteByte(')')
	return builder.String()
}

// Companies is a parsable slice of Company.
type Companies []*Company
