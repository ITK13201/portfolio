// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ITK13201/portfolio/backend/ent/image"
	"github.com/ITK13201/portfolio/backend/ent/predicate"
	"github.com/ITK13201/portfolio/backend/ent/work"
)

// WorkUpdate is the builder for updating Work entities.
type WorkUpdate struct {
	config
	hooks    []Hook
	mutation *WorkMutation
}

// Where appends a list predicates to the WorkUpdate builder.
func (wu *WorkUpdate) Where(ps ...predicate.Work) *WorkUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetTitle sets the "title" field.
func (wu *WorkUpdate) SetTitle(s string) *WorkUpdate {
	wu.mutation.SetTitle(s)
	return wu
}

// SetDescriptionJp sets the "description_jp" field.
func (wu *WorkUpdate) SetDescriptionJp(s string) *WorkUpdate {
	wu.mutation.SetDescriptionJp(s)
	return wu
}

// SetDescriptionEn sets the "description_en" field.
func (wu *WorkUpdate) SetDescriptionEn(s string) *WorkUpdate {
	wu.mutation.SetDescriptionEn(s)
	return wu
}

// SetLanguageID sets the "language_id" field.
func (wu *WorkUpdate) SetLanguageID(i int64) *WorkUpdate {
	wu.mutation.SetLanguageID(i)
	return wu
}

// SetNillableLanguageID sets the "language_id" field if the given value is not nil.
func (wu *WorkUpdate) SetNillableLanguageID(i *int64) *WorkUpdate {
	if i != nil {
		wu.SetLanguageID(*i)
	}
	return wu
}

// ClearLanguageID clears the value of the "language_id" field.
func (wu *WorkUpdate) ClearLanguageID() *WorkUpdate {
	wu.mutation.ClearLanguageID()
	return wu
}

// SetLink sets the "link" field.
func (wu *WorkUpdate) SetLink(s string) *WorkUpdate {
	wu.mutation.SetLink(s)
	return wu
}

// SetPriority sets the "priority" field.
func (wu *WorkUpdate) SetPriority(i int) *WorkUpdate {
	wu.mutation.ResetPriority()
	wu.mutation.SetPriority(i)
	return wu
}

// AddPriority adds i to the "priority" field.
func (wu *WorkUpdate) AddPriority(i int) *WorkUpdate {
	wu.mutation.AddPriority(i)
	return wu
}

// SetCreatedAt sets the "created_at" field.
func (wu *WorkUpdate) SetCreatedAt(t time.Time) *WorkUpdate {
	wu.mutation.SetCreatedAt(t)
	return wu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wu *WorkUpdate) SetNillableCreatedAt(t *time.Time) *WorkUpdate {
	if t != nil {
		wu.SetCreatedAt(*t)
	}
	return wu
}

// SetUpdatedAt sets the "updated_at" field.
func (wu *WorkUpdate) SetUpdatedAt(t time.Time) *WorkUpdate {
	wu.mutation.SetUpdatedAt(t)
	return wu
}

// SetLanguage sets the "language" edge to the Image entity.
func (wu *WorkUpdate) SetLanguage(i *Image) *WorkUpdate {
	return wu.SetLanguageID(i.ID)
}

// Mutation returns the WorkMutation object of the builder.
func (wu *WorkUpdate) Mutation() *WorkMutation {
	return wu.mutation
}

// ClearLanguage clears the "language" edge to the Image entity.
func (wu *WorkUpdate) ClearLanguage() *WorkUpdate {
	wu.mutation.ClearLanguage()
	return wu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WorkUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	wu.defaults()
	if len(wu.hooks) == 0 {
		if err = wu.check(); err != nil {
			return 0, err
		}
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wu.check(); err != nil {
				return 0, err
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			if wu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WorkUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WorkUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WorkUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wu *WorkUpdate) defaults() {
	if _, ok := wu.mutation.UpdatedAt(); !ok {
		v := work.UpdateDefaultUpdatedAt()
		wu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WorkUpdate) check() error {
	if v, ok := wu.mutation.Title(); ok {
		if err := work.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Work.title": %w`, err)}
		}
	}
	if v, ok := wu.mutation.DescriptionJp(); ok {
		if err := work.DescriptionJpValidator(v); err != nil {
			return &ValidationError{Name: "description_jp", err: fmt.Errorf(`ent: validator failed for field "Work.description_jp": %w`, err)}
		}
	}
	if v, ok := wu.mutation.DescriptionEn(); ok {
		if err := work.DescriptionEnValidator(v); err != nil {
			return &ValidationError{Name: "description_en", err: fmt.Errorf(`ent: validator failed for field "Work.description_en": %w`, err)}
		}
	}
	if v, ok := wu.mutation.Link(); ok {
		if err := work.LinkValidator(v); err != nil {
			return &ValidationError{Name: "link", err: fmt.Errorf(`ent: validator failed for field "Work.link": %w`, err)}
		}
	}
	return nil
}

func (wu *WorkUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   work.Table,
			Columns: work.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: work.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldTitle,
		})
	}
	if value, ok := wu.mutation.DescriptionJp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldDescriptionJp,
		})
	}
	if value, ok := wu.mutation.DescriptionEn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldDescriptionEn,
		})
	}
	if value, ok := wu.mutation.Link(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldLink,
		})
	}
	if value, ok := wu.mutation.Priority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: work.FieldPriority,
		})
	}
	if value, ok := wu.mutation.AddedPriority(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: work.FieldPriority,
		})
	}
	if value, ok := wu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: work.FieldCreatedAt,
		})
	}
	if value, ok := wu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: work.FieldUpdatedAt,
		})
	}
	if wu.mutation.LanguageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   work.LanguageTable,
			Columns: []string{work.LanguageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: image.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.LanguageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   work.LanguageTable,
			Columns: []string{work.LanguageColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{work.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WorkUpdateOne is the builder for updating a single Work entity.
type WorkUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WorkMutation
}

// SetTitle sets the "title" field.
func (wuo *WorkUpdateOne) SetTitle(s string) *WorkUpdateOne {
	wuo.mutation.SetTitle(s)
	return wuo
}

// SetDescriptionJp sets the "description_jp" field.
func (wuo *WorkUpdateOne) SetDescriptionJp(s string) *WorkUpdateOne {
	wuo.mutation.SetDescriptionJp(s)
	return wuo
}

// SetDescriptionEn sets the "description_en" field.
func (wuo *WorkUpdateOne) SetDescriptionEn(s string) *WorkUpdateOne {
	wuo.mutation.SetDescriptionEn(s)
	return wuo
}

// SetLanguageID sets the "language_id" field.
func (wuo *WorkUpdateOne) SetLanguageID(i int64) *WorkUpdateOne {
	wuo.mutation.SetLanguageID(i)
	return wuo
}

// SetNillableLanguageID sets the "language_id" field if the given value is not nil.
func (wuo *WorkUpdateOne) SetNillableLanguageID(i *int64) *WorkUpdateOne {
	if i != nil {
		wuo.SetLanguageID(*i)
	}
	return wuo
}

// ClearLanguageID clears the value of the "language_id" field.
func (wuo *WorkUpdateOne) ClearLanguageID() *WorkUpdateOne {
	wuo.mutation.ClearLanguageID()
	return wuo
}

// SetLink sets the "link" field.
func (wuo *WorkUpdateOne) SetLink(s string) *WorkUpdateOne {
	wuo.mutation.SetLink(s)
	return wuo
}

// SetPriority sets the "priority" field.
func (wuo *WorkUpdateOne) SetPriority(i int) *WorkUpdateOne {
	wuo.mutation.ResetPriority()
	wuo.mutation.SetPriority(i)
	return wuo
}

// AddPriority adds i to the "priority" field.
func (wuo *WorkUpdateOne) AddPriority(i int) *WorkUpdateOne {
	wuo.mutation.AddPriority(i)
	return wuo
}

// SetCreatedAt sets the "created_at" field.
func (wuo *WorkUpdateOne) SetCreatedAt(t time.Time) *WorkUpdateOne {
	wuo.mutation.SetCreatedAt(t)
	return wuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wuo *WorkUpdateOne) SetNillableCreatedAt(t *time.Time) *WorkUpdateOne {
	if t != nil {
		wuo.SetCreatedAt(*t)
	}
	return wuo
}

// SetUpdatedAt sets the "updated_at" field.
func (wuo *WorkUpdateOne) SetUpdatedAt(t time.Time) *WorkUpdateOne {
	wuo.mutation.SetUpdatedAt(t)
	return wuo
}

// SetLanguage sets the "language" edge to the Image entity.
func (wuo *WorkUpdateOne) SetLanguage(i *Image) *WorkUpdateOne {
	return wuo.SetLanguageID(i.ID)
}

// Mutation returns the WorkMutation object of the builder.
func (wuo *WorkUpdateOne) Mutation() *WorkMutation {
	return wuo.mutation
}

// ClearLanguage clears the "language" edge to the Image entity.
func (wuo *WorkUpdateOne) ClearLanguage() *WorkUpdateOne {
	wuo.mutation.ClearLanguage()
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WorkUpdateOne) Select(field string, fields ...string) *WorkUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Work entity.
func (wuo *WorkUpdateOne) Save(ctx context.Context) (*Work, error) {
	var (
		err  error
		node *Work
	)
	wuo.defaults()
	if len(wuo.hooks) == 0 {
		if err = wuo.check(); err != nil {
			return nil, err
		}
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wuo.check(); err != nil {
				return nil, err
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			if wuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WorkUpdateOne) SaveX(ctx context.Context) *Work {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WorkUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WorkUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuo *WorkUpdateOne) defaults() {
	if _, ok := wuo.mutation.UpdatedAt(); !ok {
		v := work.UpdateDefaultUpdatedAt()
		wuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WorkUpdateOne) check() error {
	if v, ok := wuo.mutation.Title(); ok {
		if err := work.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Work.title": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.DescriptionJp(); ok {
		if err := work.DescriptionJpValidator(v); err != nil {
			return &ValidationError{Name: "description_jp", err: fmt.Errorf(`ent: validator failed for field "Work.description_jp": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.DescriptionEn(); ok {
		if err := work.DescriptionEnValidator(v); err != nil {
			return &ValidationError{Name: "description_en", err: fmt.Errorf(`ent: validator failed for field "Work.description_en": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.Link(); ok {
		if err := work.LinkValidator(v); err != nil {
			return &ValidationError{Name: "link", err: fmt.Errorf(`ent: validator failed for field "Work.link": %w`, err)}
		}
	}
	return nil
}

func (wuo *WorkUpdateOne) sqlSave(ctx context.Context) (_node *Work, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   work.Table,
			Columns: work.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: work.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Work.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, work.FieldID)
		for _, f := range fields {
			if !work.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != work.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldTitle,
		})
	}
	if value, ok := wuo.mutation.DescriptionJp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldDescriptionJp,
		})
	}
	if value, ok := wuo.mutation.DescriptionEn(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldDescriptionEn,
		})
	}
	if value, ok := wuo.mutation.Link(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: work.FieldLink,
		})
	}
	if value, ok := wuo.mutation.Priority(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: work.FieldPriority,
		})
	}
	if value, ok := wuo.mutation.AddedPriority(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: work.FieldPriority,
		})
	}
	if value, ok := wuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: work.FieldCreatedAt,
		})
	}
	if value, ok := wuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: work.FieldUpdatedAt,
		})
	}
	if wuo.mutation.LanguageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   work.LanguageTable,
			Columns: []string{work.LanguageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: image.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.LanguageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   work.LanguageTable,
			Columns: []string{work.LanguageColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Work{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{work.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
