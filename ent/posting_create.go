// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/techbloghub/server/ent/company"
	"github.com/techbloghub/server/ent/posting"
	"github.com/techbloghub/server/internal/schemasupport"
)

// PostingCreate is the builder for creating a Posting entity.
type PostingCreate struct {
	config
	mutation *PostingMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (pc *PostingCreate) SetCreateTime(t time.Time) *PostingCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *PostingCreate) SetNillableCreateTime(t *time.Time) *PostingCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *PostingCreate) SetUpdateTime(t time.Time) *PostingCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *PostingCreate) SetNillableUpdateTime(t *time.Time) *PostingCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetTitle sets the "title" field.
func (pc *PostingCreate) SetTitle(s string) *PostingCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetURL sets the "url" field.
func (pc *PostingCreate) SetURL(u *url.URL) *PostingCreate {
	pc.mutation.SetURL(u)
	return pc
}

// SetPublishedAt sets the "published_at" field.
func (pc *PostingCreate) SetPublishedAt(t time.Time) *PostingCreate {
	pc.mutation.SetPublishedAt(t)
	return pc
}

// SetTags sets the "tags" field.
func (pc *PostingCreate) SetTags(st *schemasupport.PostingTags) *PostingCreate {
	pc.mutation.SetTags(st)
	return pc
}

// SetCompanyID sets the "company" edge to the Company entity by ID.
func (pc *PostingCreate) SetCompanyID(id int) *PostingCreate {
	pc.mutation.SetCompanyID(id)
	return pc
}

// SetNillableCompanyID sets the "company" edge to the Company entity by ID if the given value is not nil.
func (pc *PostingCreate) SetNillableCompanyID(id *int) *PostingCreate {
	if id != nil {
		pc = pc.SetCompanyID(*id)
	}
	return pc
}

// SetCompany sets the "company" edge to the Company entity.
func (pc *PostingCreate) SetCompany(c *Company) *PostingCreate {
	return pc.SetCompanyID(c.ID)
}

// Mutation returns the PostingMutation object of the builder.
func (pc *PostingCreate) Mutation() *PostingMutation {
	return pc.mutation
}

// Save creates the Posting in the database.
func (pc *PostingCreate) Save(ctx context.Context) (*Posting, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostingCreate) SaveX(ctx context.Context) *Posting {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PostingCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PostingCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PostingCreate) defaults() {
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := posting.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := posting.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PostingCreate) check() error {
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Posting.create_time"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Posting.update_time"`)}
	}
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Posting.title"`)}
	}
	if _, ok := pc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Posting.url"`)}
	}
	if _, ok := pc.mutation.PublishedAt(); !ok {
		return &ValidationError{Name: "published_at", err: errors.New(`ent: missing required field "Posting.published_at"`)}
	}
	return nil
}

func (pc *PostingCreate) sqlSave(ctx context.Context) (*Posting, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec, err := pc.createSpec()
	if err != nil {
		return nil, err
	}
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PostingCreate) createSpec() (*Posting, *sqlgraph.CreateSpec, error) {
	var (
		_node = &Posting{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(posting.Table, sqlgraph.NewFieldSpec(posting.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.SetField(posting.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.SetField(posting.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.SetField(posting.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := pc.mutation.URL(); ok {
		vv, err := posting.ValueScanner.URL.Value(value)
		if err != nil {
			return nil, nil, err
		}
		_spec.SetField(posting.FieldURL, field.TypeString, vv)
		_node.URL = value
	}
	if value, ok := pc.mutation.PublishedAt(); ok {
		_spec.SetField(posting.FieldPublishedAt, field.TypeTime, value)
		_node.PublishedAt = value
	}
	if value, ok := pc.mutation.Tags(); ok {
		_spec.SetField(posting.FieldTags, field.TypeOther, value)
		_node.Tags = value
	}
	if nodes := pc.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   posting.CompanyTable,
			Columns: []string{posting.CompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(company.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.company_postings = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec, nil
}

// PostingCreateBulk is the builder for creating many Posting entities in bulk.
type PostingCreateBulk struct {
	config
	err      error
	builders []*PostingCreate
}

// Save creates the Posting entities in the database.
func (pcb *PostingCreateBulk) Save(ctx context.Context) ([]*Posting, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Posting, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i], err = builder.createSpec()
				if err != nil {
					return nil, err
				}
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PostingCreateBulk) SaveX(ctx context.Context) []*Posting {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PostingCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PostingCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
