// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/direktiv/direktiv/pkg/flow/ent/events"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/instanceruntime"
	"github.com/direktiv/direktiv/pkg/flow/ent/logmsg"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/revision"
	"github.com/direktiv/direktiv/pkg/flow/ent/varref"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// InstanceCreate is the builder for creating a Instance entity.
type InstanceCreate struct {
	config
	mutation *InstanceMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ic *InstanceCreate) SetCreatedAt(t time.Time) *InstanceCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *InstanceCreate) SetNillableCreatedAt(t *time.Time) *InstanceCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *InstanceCreate) SetUpdatedAt(t time.Time) *InstanceCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *InstanceCreate) SetNillableUpdatedAt(t *time.Time) *InstanceCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetEndAt sets the "end_at" field.
func (ic *InstanceCreate) SetEndAt(t time.Time) *InstanceCreate {
	ic.mutation.SetEndAt(t)
	return ic
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (ic *InstanceCreate) SetNillableEndAt(t *time.Time) *InstanceCreate {
	if t != nil {
		ic.SetEndAt(*t)
	}
	return ic
}

// SetStatus sets the "status" field.
func (ic *InstanceCreate) SetStatus(s string) *InstanceCreate {
	ic.mutation.SetStatus(s)
	return ic
}

// SetAs sets the "as" field.
func (ic *InstanceCreate) SetAs(s string) *InstanceCreate {
	ic.mutation.SetAs(s)
	return ic
}

// SetErrorCode sets the "errorCode" field.
func (ic *InstanceCreate) SetErrorCode(s string) *InstanceCreate {
	ic.mutation.SetErrorCode(s)
	return ic
}

// SetNillableErrorCode sets the "errorCode" field if the given value is not nil.
func (ic *InstanceCreate) SetNillableErrorCode(s *string) *InstanceCreate {
	if s != nil {
		ic.SetErrorCode(*s)
	}
	return ic
}

// SetErrorMessage sets the "errorMessage" field.
func (ic *InstanceCreate) SetErrorMessage(s string) *InstanceCreate {
	ic.mutation.SetErrorMessage(s)
	return ic
}

// SetNillableErrorMessage sets the "errorMessage" field if the given value is not nil.
func (ic *InstanceCreate) SetNillableErrorMessage(s *string) *InstanceCreate {
	if s != nil {
		ic.SetErrorMessage(*s)
	}
	return ic
}

// SetInvoker sets the "invoker" field.
func (ic *InstanceCreate) SetInvoker(s string) *InstanceCreate {
	ic.mutation.SetInvoker(s)
	return ic
}

// SetNillableInvoker sets the "invoker" field if the given value is not nil.
func (ic *InstanceCreate) SetNillableInvoker(s *string) *InstanceCreate {
	if s != nil {
		ic.SetInvoker(*s)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *InstanceCreate) SetID(u uuid.UUID) *InstanceCreate {
	ic.mutation.SetID(u)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *InstanceCreate) SetNillableID(u *uuid.UUID) *InstanceCreate {
	if u != nil {
		ic.SetID(*u)
	}
	return ic
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (ic *InstanceCreate) SetNamespaceID(id uuid.UUID) *InstanceCreate {
	ic.mutation.SetNamespaceID(id)
	return ic
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (ic *InstanceCreate) SetNamespace(n *Namespace) *InstanceCreate {
	return ic.SetNamespaceID(n.ID)
}

// SetWorkflowID sets the "workflow" edge to the Workflow entity by ID.
func (ic *InstanceCreate) SetWorkflowID(id uuid.UUID) *InstanceCreate {
	ic.mutation.SetWorkflowID(id)
	return ic
}

// SetNillableWorkflowID sets the "workflow" edge to the Workflow entity by ID if the given value is not nil.
func (ic *InstanceCreate) SetNillableWorkflowID(id *uuid.UUID) *InstanceCreate {
	if id != nil {
		ic = ic.SetWorkflowID(*id)
	}
	return ic
}

// SetWorkflow sets the "workflow" edge to the Workflow entity.
func (ic *InstanceCreate) SetWorkflow(w *Workflow) *InstanceCreate {
	return ic.SetWorkflowID(w.ID)
}

// SetRevisionID sets the "revision" edge to the Revision entity by ID.
func (ic *InstanceCreate) SetRevisionID(id uuid.UUID) *InstanceCreate {
	ic.mutation.SetRevisionID(id)
	return ic
}

// SetNillableRevisionID sets the "revision" edge to the Revision entity by ID if the given value is not nil.
func (ic *InstanceCreate) SetNillableRevisionID(id *uuid.UUID) *InstanceCreate {
	if id != nil {
		ic = ic.SetRevisionID(*id)
	}
	return ic
}

// SetRevision sets the "revision" edge to the Revision entity.
func (ic *InstanceCreate) SetRevision(r *Revision) *InstanceCreate {
	return ic.SetRevisionID(r.ID)
}

// AddLogIDs adds the "logs" edge to the LogMsg entity by IDs.
func (ic *InstanceCreate) AddLogIDs(ids ...uuid.UUID) *InstanceCreate {
	ic.mutation.AddLogIDs(ids...)
	return ic
}

// AddLogs adds the "logs" edges to the LogMsg entity.
func (ic *InstanceCreate) AddLogs(l ...*LogMsg) *InstanceCreate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ic.AddLogIDs(ids...)
}

// AddVarIDs adds the "vars" edge to the VarRef entity by IDs.
func (ic *InstanceCreate) AddVarIDs(ids ...uuid.UUID) *InstanceCreate {
	ic.mutation.AddVarIDs(ids...)
	return ic
}

// AddVars adds the "vars" edges to the VarRef entity.
func (ic *InstanceCreate) AddVars(v ...*VarRef) *InstanceCreate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return ic.AddVarIDs(ids...)
}

// SetRuntimeID sets the "runtime" edge to the InstanceRuntime entity by ID.
func (ic *InstanceCreate) SetRuntimeID(id uuid.UUID) *InstanceCreate {
	ic.mutation.SetRuntimeID(id)
	return ic
}

// SetRuntime sets the "runtime" edge to the InstanceRuntime entity.
func (ic *InstanceCreate) SetRuntime(i *InstanceRuntime) *InstanceCreate {
	return ic.SetRuntimeID(i.ID)
}

// AddChildIDs adds the "children" edge to the InstanceRuntime entity by IDs.
func (ic *InstanceCreate) AddChildIDs(ids ...uuid.UUID) *InstanceCreate {
	ic.mutation.AddChildIDs(ids...)
	return ic
}

// AddChildren adds the "children" edges to the InstanceRuntime entity.
func (ic *InstanceCreate) AddChildren(i ...*InstanceRuntime) *InstanceCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ic.AddChildIDs(ids...)
}

// AddEventlistenerIDs adds the "eventlisteners" edge to the Events entity by IDs.
func (ic *InstanceCreate) AddEventlistenerIDs(ids ...uuid.UUID) *InstanceCreate {
	ic.mutation.AddEventlistenerIDs(ids...)
	return ic
}

// AddEventlisteners adds the "eventlisteners" edges to the Events entity.
func (ic *InstanceCreate) AddEventlisteners(e ...*Events) *InstanceCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return ic.AddEventlistenerIDs(ids...)
}

// Mutation returns the InstanceMutation object of the builder.
func (ic *InstanceCreate) Mutation() *InstanceMutation {
	return ic.mutation
}

// Save creates the Instance in the database.
func (ic *InstanceCreate) Save(ctx context.Context) (*Instance, error) {
	var (
		err  error
		node *Instance
	)
	ic.defaults()
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstanceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			if node, err = ic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			if ic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Instance)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from InstanceMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InstanceCreate) SaveX(ctx context.Context) *Instance {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *InstanceCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *InstanceCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *InstanceCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := instance.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := instance.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		v := instance.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InstanceCreate) check() error {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Instance.created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Instance.updated_at"`)}
	}
	if _, ok := ic.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Instance.status"`)}
	}
	if _, ok := ic.mutation.As(); !ok {
		return &ValidationError{Name: "as", err: errors.New(`ent: missing required field "Instance.as"`)}
	}
	if _, ok := ic.mutation.NamespaceID(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`ent: missing required edge "Instance.namespace"`)}
	}
	if _, ok := ic.mutation.RuntimeID(); !ok {
		return &ValidationError{Name: "runtime", err: errors.New(`ent: missing required edge "Instance.runtime"`)}
	}
	return nil
}

func (ic *InstanceCreate) sqlSave(ctx context.Context) (*Instance, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ic *InstanceCreate) createSpec() (*Instance, *sqlgraph.CreateSpec) {
	var (
		_node = &Instance{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: instance.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: instance.FieldID,
			},
		}
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instance.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instance.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ic.mutation.EndAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: instance.FieldEndAt,
		})
		_node.EndAt = value
	}
	if value, ok := ic.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instance.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := ic.mutation.As(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instance.FieldAs,
		})
		_node.As = value
	}
	if value, ok := ic.mutation.ErrorCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instance.FieldErrorCode,
		})
		_node.ErrorCode = value
	}
	if value, ok := ic.mutation.ErrorMessage(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instance.FieldErrorMessage,
		})
		_node.ErrorMessage = value
	}
	if value, ok := ic.mutation.Invoker(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instance.FieldInvoker,
		})
		_node.Invoker = value
	}
	if nodes := ic.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   instance.NamespaceTable,
			Columns: []string{instance.NamespaceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: namespace.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.namespace_instances = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.WorkflowIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   instance.WorkflowTable,
			Columns: []string{instance.WorkflowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: workflow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.workflow_instances = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.RevisionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   instance.RevisionTable,
			Columns: []string{instance.RevisionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: revision.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.revision_instances = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.LogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   instance.LogsTable,
			Columns: []string{instance.LogsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: logmsg.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.VarsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   instance.VarsTable,
			Columns: []string{instance.VarsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: varref.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.RuntimeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   instance.RuntimeTable,
			Columns: []string{instance.RuntimeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instanceruntime.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   instance.ChildrenTable,
			Columns: []string{instance.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instanceruntime.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.EventlistenersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   instance.EventlistenersTable,
			Columns: []string{instance.EventlistenersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: events.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InstanceCreateBulk is the builder for creating many Instance entities in bulk.
type InstanceCreateBulk struct {
	config
	builders []*InstanceCreate
}

// Save creates the Instance entities in the database.
func (icb *InstanceCreateBulk) Save(ctx context.Context) ([]*Instance, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Instance, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InstanceMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InstanceCreateBulk) SaveX(ctx context.Context) []*Instance {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *InstanceCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *InstanceCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
