// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ITK13201/portfolio/backend/ent/abouttopic"
	"github.com/ITK13201/portfolio/backend/ent/image"
)

// AboutTopicCreate is the builder for creating a AboutTopic entity.
type AboutTopicCreate struct {
	config
	mutation *AboutTopicMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (atc *AboutTopicCreate) SetTitle(s string) *AboutTopicCreate {
	atc.mutation.SetTitle(s)
	return atc
}

// SetDescriptionJp sets the "description_jp" field.
func (atc *AboutTopicCreate) SetDescriptionJp(s string) *AboutTopicCreate {
	atc.mutation.SetDescriptionJp(s)
	return atc
}

// SetDescriptionEn sets the "description_en" field.
func (atc *AboutTopicCreate) SetDescriptionEn(s string) *AboutTopicCreate {
	atc.mutation.SetDescriptionEn(s)
	return atc
}

// SetImageID sets the "image_id" field.
func (atc *AboutTopicCreate) SetImageID(i int64) *AboutTopicCreate {
	atc.mutation.SetImageID(i)
	return atc
}

// SetNillableImageID sets the "image_id" field if the given value is not nil.
func (atc *AboutTopicCreate) SetNillableImageID(i *int64) *AboutTopicCreate {
	if i != nil {
		atc.SetImageID(*i)
	}
	return atc
}

// SetPriority sets the "priority" field.
func (atc *AboutTopicCreate) SetPriority(i int) *AboutTopicCreate {
	atc.mutation.SetPriority(i)
	return atc
}

// SetCreatedAt sets the "created_at" field.
func (atc *AboutTopicCreate) SetCreatedAt(t time.Time) *AboutTopicCreate {
	atc.mutation.SetCreatedAt(t)
	return atc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (atc *AboutTopicCreate) SetNillableCreatedAt(t *time.Time) *AboutTopicCreate {
	if t != nil {
		atc.SetCreatedAt(*t)
	}
	return atc
}

// SetUpdatedAt sets the "updated_at" field.
func (atc *AboutTopicCreate) SetUpdatedAt(t time.Time) *AboutTopicCreate {
	atc.mutation.SetUpdatedAt(t)
	return atc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (atc *AboutTopicCreate) SetNillableUpdatedAt(t *time.Time) *AboutTopicCreate {
	if t != nil {
		atc.SetUpdatedAt(*t)
	}
	return atc
}

// SetID sets the "id" field.
func (atc *AboutTopicCreate) SetID(i int64) *AboutTopicCreate {
	atc.mutation.SetID(i)
	return atc
}

// SetImage sets the "image" edge to the Image entity.
func (atc *AboutTopicCreate) SetImage(i *Image) *AboutTopicCreate {
	return atc.SetImageID(i.ID)
}

// Mutation returns the AboutTopicMutation object of the builder.
func (atc *AboutTopicCreate) Mutation() *AboutTopicMutation {
	return atc.mutation
}

// Save creates the AboutTopic in the database.
func (atc *AboutTopicCreate) Save(ctx context.Context) (*AboutTopic, error) {
	var (
		err  error
		node *AboutTopic
	)
	atc.defaults()
	if len(atc.hooks) == 0 {
		if err = atc.check(); err != nil {
			return nil, err
		}
		node, err = atc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AboutTopicMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = atc.check(); err != nil {
				return nil, err
			}
			atc.mutation = mutation
			if node, err = atc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(atc.hooks) - 1; i >= 0; i-- {
			if atc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = atc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, atc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (atc *AboutTopicCreate) SaveX(ctx context.Context) *AboutTopic {
	v, err := atc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atc *AboutTopicCreate) Exec(ctx context.Context) error {
	_, err := atc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atc *AboutTopicCreate) ExecX(ctx context.Context) {
	if err := atc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (atc *AboutTopicCreate) defaults() {
	if _, ok := atc.mutation.CreatedAt(); !ok {
		v := abouttopic.DefaultCreatedAt()
		atc.mutation.SetCreatedAt(v)
	}
	if _, ok := atc.mutation.UpdatedAt(); !ok {
		v := abouttopic.DefaultUpdatedAt()
		atc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (atc *AboutTopicCreate) check() error {
	if _, ok := atc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "AboutTopic.title"`)}
	}
	if v, ok := atc.mutation.Title(); ok {
		if err := abouttopic.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "AboutTopic.title": %w`, err)}
		}
	}
	if _, ok := atc.mutation.DescriptionJp(); !ok {
		return &ValidationError{Name: "description_jp", err: errors.New(`ent: missing required field "AboutTopic.description_jp"`)}
	}
	if v, ok := atc.mutation.DescriptionJp(); ok {
		if err := abouttopic.DescriptionJpValidator(v); err != nil {
			return &ValidationError{Name: "description_jp", err: fmt.Errorf(`ent: validator failed for field "AboutTopic.description_jp": %w`, err)}
		}
	}
	if _, ok := atc.mutation.DescriptionEn(); !ok {
		return &ValidationError{Name: "description_en", err: errors.New(`ent: missing required field "AboutTopic.description_en"`)}
	}
	if v, ok := atc.mutation.DescriptionEn(); ok {
		if err := abouttopic.DescriptionEnValidator(v); err != nil {
			return &ValidationError{Name: "description_en", err: fmt.Errorf(`ent: validator failed for field "AboutTopic.description_en": %w`, err)}
		}
	}
	if _, ok := atc.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`ent: missing required field "AboutTopic.priority"`)}
	}
	if _, ok := atc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AboutTopic.created_at"`)}
	}
	if _, ok := atc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AboutTopic.updated_at"`)}
	}
	return nil
}

func (atc *AboutTopicCreate) sqlSave(ctx context.Context) (*AboutTopic, error) {
	_node, _spec := atc.createSpec()
	if err := sqlgraph.CreateNode(ctx, atc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (atc *AboutTopicCreate) createSpec() (*AboutTopic, *sqlgraph.CreateSpec) {
	var (
		_node = &AboutTopic{config: atc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: abouttopic.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: abouttopic.FieldID,
			},
		}
	)
	if id, ok := atc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := atc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: abouttopic.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := atc.mutation.DescriptionJp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: abouttopic.FieldDescriptionJp,
		})
		_node.DescriptionJp = value
	}
	if value, ok := atc.mutation.DescriptionEn(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: abouttopic.FieldDescriptionEn,
		})
		_node.DescriptionEn = value
	}
	if value, ok := atc.mutation.Priority(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: abouttopic.FieldPriority,
		})
		_node.Priority = value
	}
	if value, ok := atc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: abouttopic.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := atc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: abouttopic.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := atc.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   abouttopic.ImageTable,
			Columns: []string{abouttopic.ImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: image.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ImageID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AboutTopicCreateBulk is the builder for creating many AboutTopic entities in bulk.
type AboutTopicCreateBulk struct {
	config
	builders []*AboutTopicCreate
}

// Save creates the AboutTopic entities in the database.
func (atcb *AboutTopicCreateBulk) Save(ctx context.Context) ([]*AboutTopic, error) {
	specs := make([]*sqlgraph.CreateSpec, len(atcb.builders))
	nodes := make([]*AboutTopic, len(atcb.builders))
	mutators := make([]Mutator, len(atcb.builders))
	for i := range atcb.builders {
		func(i int, root context.Context) {
			builder := atcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AboutTopicMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, atcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, atcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, atcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (atcb *AboutTopicCreateBulk) SaveX(ctx context.Context) []*AboutTopic {
	v, err := atcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (atcb *AboutTopicCreateBulk) Exec(ctx context.Context) error {
	_, err := atcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (atcb *AboutTopicCreateBulk) ExecX(ctx context.Context) {
	if err := atcb.Exec(ctx); err != nil {
		panic(err)
	}
}
