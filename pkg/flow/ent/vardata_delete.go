// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/predicate"
	"github.com/direktiv/direktiv/pkg/flow/ent/vardata"
)

// VarDataDelete is the builder for deleting a VarData entity.
type VarDataDelete struct {
	config
	hooks    []Hook
	mutation *VarDataMutation
}

// Where appends a list predicates to the VarDataDelete builder.
func (vdd *VarDataDelete) Where(ps ...predicate.VarData) *VarDataDelete {
	vdd.mutation.Where(ps...)
	return vdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vdd *VarDataDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vdd.hooks) == 0 {
		affected, err = vdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VarDataMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vdd.mutation = mutation
			affected, err = vdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vdd.hooks) - 1; i >= 0; i-- {
			if vdd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vdd *VarDataDelete) ExecX(ctx context.Context) int {
	n, err := vdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vdd *VarDataDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: vardata.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: vardata.FieldID,
			},
		},
	}
	if ps := vdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// VarDataDeleteOne is the builder for deleting a single VarData entity.
type VarDataDeleteOne struct {
	vdd *VarDataDelete
}

// Exec executes the deletion query.
func (vddo *VarDataDeleteOne) Exec(ctx context.Context) error {
	n, err := vddo.vdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vardata.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vddo *VarDataDeleteOne) ExecX(ctx context.Context) {
	vddo.vdd.ExecX(ctx)
}
