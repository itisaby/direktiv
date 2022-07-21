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
	"github.com/direktiv/direktiv/pkg/flow/ent/inode"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/logmsg"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/ref"
	"github.com/direktiv/direktiv/pkg/flow/ent/revision"
	"github.com/direktiv/direktiv/pkg/flow/ent/route"
	"github.com/direktiv/direktiv/pkg/flow/ent/varref"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// WorkflowCreate is the builder for creating a Workflow entity.
type WorkflowCreate struct {
	config
	mutation *WorkflowMutation
	hooks    []Hook
}

// SetLive sets the "live" field.
func (wc *WorkflowCreate) SetLive(b bool) *WorkflowCreate {
	wc.mutation.SetLive(b)
	return wc
}

// SetNillableLive sets the "live" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableLive(b *bool) *WorkflowCreate {
	if b != nil {
		wc.SetLive(*b)
	}
	return wc
}

// SetLogToEvents sets the "logToEvents" field.
func (wc *WorkflowCreate) SetLogToEvents(s string) *WorkflowCreate {
	wc.mutation.SetLogToEvents(s)
	return wc
}

// SetNillableLogToEvents sets the "logToEvents" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableLogToEvents(s *string) *WorkflowCreate {
	if s != nil {
		wc.SetLogToEvents(*s)
	}
	return wc
}

// SetReadOnly sets the "readOnly" field.
func (wc *WorkflowCreate) SetReadOnly(b bool) *WorkflowCreate {
	wc.mutation.SetReadOnly(b)
	return wc
}

// SetNillableReadOnly sets the "readOnly" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableReadOnly(b *bool) *WorkflowCreate {
	if b != nil {
		wc.SetReadOnly(*b)
	}
	return wc
}

// SetUpdatedAt sets the "updated_at" field.
func (wc *WorkflowCreate) SetUpdatedAt(t time.Time) *WorkflowCreate {
	wc.mutation.SetUpdatedAt(t)
	return wc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableUpdatedAt(t *time.Time) *WorkflowCreate {
	if t != nil {
		wc.SetUpdatedAt(*t)
	}
	return wc
}

// SetID sets the "id" field.
func (wc *WorkflowCreate) SetID(u uuid.UUID) *WorkflowCreate {
	wc.mutation.SetID(u)
	return wc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (wc *WorkflowCreate) SetNillableID(u *uuid.UUID) *WorkflowCreate {
	if u != nil {
		wc.SetID(*u)
	}
	return wc
}

// SetInodeID sets the "inode" edge to the Inode entity by ID.
func (wc *WorkflowCreate) SetInodeID(id uuid.UUID) *WorkflowCreate {
	wc.mutation.SetInodeID(id)
	return wc
}

// SetNillableInodeID sets the "inode" edge to the Inode entity by ID if the given value is not nil.
func (wc *WorkflowCreate) SetNillableInodeID(id *uuid.UUID) *WorkflowCreate {
	if id != nil {
		wc = wc.SetInodeID(*id)
	}
	return wc
}

// SetInode sets the "inode" edge to the Inode entity.
func (wc *WorkflowCreate) SetInode(i *Inode) *WorkflowCreate {
	return wc.SetInodeID(i.ID)
}

// SetNamespaceID sets the "namespace" edge to the Namespace entity by ID.
func (wc *WorkflowCreate) SetNamespaceID(id uuid.UUID) *WorkflowCreate {
	wc.mutation.SetNamespaceID(id)
	return wc
}

// SetNamespace sets the "namespace" edge to the Namespace entity.
func (wc *WorkflowCreate) SetNamespace(n *Namespace) *WorkflowCreate {
	return wc.SetNamespaceID(n.ID)
}

// AddRevisionIDs adds the "revisions" edge to the Revision entity by IDs.
func (wc *WorkflowCreate) AddRevisionIDs(ids ...uuid.UUID) *WorkflowCreate {
	wc.mutation.AddRevisionIDs(ids...)
	return wc
}

// AddRevisions adds the "revisions" edges to the Revision entity.
func (wc *WorkflowCreate) AddRevisions(r ...*Revision) *WorkflowCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return wc.AddRevisionIDs(ids...)
}

// AddRefIDs adds the "refs" edge to the Ref entity by IDs.
func (wc *WorkflowCreate) AddRefIDs(ids ...uuid.UUID) *WorkflowCreate {
	wc.mutation.AddRefIDs(ids...)
	return wc
}

// AddRefs adds the "refs" edges to the Ref entity.
func (wc *WorkflowCreate) AddRefs(r ...*Ref) *WorkflowCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return wc.AddRefIDs(ids...)
}

// AddInstanceIDs adds the "instances" edge to the Instance entity by IDs.
func (wc *WorkflowCreate) AddInstanceIDs(ids ...uuid.UUID) *WorkflowCreate {
	wc.mutation.AddInstanceIDs(ids...)
	return wc
}

// AddInstances adds the "instances" edges to the Instance entity.
func (wc *WorkflowCreate) AddInstances(i ...*Instance) *WorkflowCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return wc.AddInstanceIDs(ids...)
}

// AddRouteIDs adds the "routes" edge to the Route entity by IDs.
func (wc *WorkflowCreate) AddRouteIDs(ids ...uuid.UUID) *WorkflowCreate {
	wc.mutation.AddRouteIDs(ids...)
	return wc
}

// AddRoutes adds the "routes" edges to the Route entity.
func (wc *WorkflowCreate) AddRoutes(r ...*Route) *WorkflowCreate {
	ids := make([]uuid.UUID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return wc.AddRouteIDs(ids...)
}

// AddLogIDs adds the "logs" edge to the LogMsg entity by IDs.
func (wc *WorkflowCreate) AddLogIDs(ids ...uuid.UUID) *WorkflowCreate {
	wc.mutation.AddLogIDs(ids...)
	return wc
}

// AddLogs adds the "logs" edges to the LogMsg entity.
func (wc *WorkflowCreate) AddLogs(l ...*LogMsg) *WorkflowCreate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return wc.AddLogIDs(ids...)
}

// AddVarIDs adds the "vars" edge to the VarRef entity by IDs.
func (wc *WorkflowCreate) AddVarIDs(ids ...uuid.UUID) *WorkflowCreate {
	wc.mutation.AddVarIDs(ids...)
	return wc
}

// AddVars adds the "vars" edges to the VarRef entity.
func (wc *WorkflowCreate) AddVars(v ...*VarRef) *WorkflowCreate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return wc.AddVarIDs(ids...)
}

// AddWfeventIDs adds the "wfevents" edge to the Events entity by IDs.
func (wc *WorkflowCreate) AddWfeventIDs(ids ...uuid.UUID) *WorkflowCreate {
	wc.mutation.AddWfeventIDs(ids...)
	return wc
}

// AddWfevents adds the "wfevents" edges to the Events entity.
func (wc *WorkflowCreate) AddWfevents(e ...*Events) *WorkflowCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return wc.AddWfeventIDs(ids...)
}

// Mutation returns the WorkflowMutation object of the builder.
func (wc *WorkflowCreate) Mutation() *WorkflowMutation {
	return wc.mutation
}

// Save creates the Workflow in the database.
func (wc *WorkflowCreate) Save(ctx context.Context) (*Workflow, error) {
	var (
		err  error
		node *Workflow
	)
	wc.defaults()
	if len(wc.hooks) == 0 {
		if err = wc.check(); err != nil {
			return nil, err
		}
		node, err = wc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WorkflowMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wc.check(); err != nil {
				return nil, err
			}
			wc.mutation = mutation
			if node, err = wc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(wc.hooks) - 1; i >= 0; i-- {
			if wc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, wc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Workflow)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from WorkflowMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WorkflowCreate) SaveX(ctx context.Context) *Workflow {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WorkflowCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WorkflowCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wc *WorkflowCreate) defaults() {
	if _, ok := wc.mutation.Live(); !ok {
		v := workflow.DefaultLive
		wc.mutation.SetLive(v)
	}
	if _, ok := wc.mutation.UpdatedAt(); !ok {
		v := workflow.DefaultUpdatedAt()
		wc.mutation.SetUpdatedAt(v)
	}
	if _, ok := wc.mutation.ID(); !ok {
		v := workflow.DefaultID()
		wc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WorkflowCreate) check() error {
	if _, ok := wc.mutation.Live(); !ok {
		return &ValidationError{Name: "live", err: errors.New(`ent: missing required field "Workflow.live"`)}
	}
	if _, ok := wc.mutation.NamespaceID(); !ok {
		return &ValidationError{Name: "namespace", err: errors.New(`ent: missing required edge "Workflow.namespace"`)}
	}
	return nil
}

func (wc *WorkflowCreate) sqlSave(ctx context.Context) (*Workflow, error) {
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
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

func (wc *WorkflowCreate) createSpec() (*Workflow, *sqlgraph.CreateSpec) {
	var (
		_node = &Workflow{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: workflow.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: workflow.FieldID,
			},
		}
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := wc.mutation.Live(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: workflow.FieldLive,
		})
		_node.Live = value
	}
	if value, ok := wc.mutation.LogToEvents(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: workflow.FieldLogToEvents,
		})
		_node.LogToEvents = value
	}
	if value, ok := wc.mutation.ReadOnly(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: workflow.FieldReadOnly,
		})
		_node.ReadOnly = value
	}
	if value, ok := wc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: workflow.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := wc.mutation.InodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   workflow.InodeTable,
			Columns: []string{workflow.InodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: inode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.inode_workflow = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.NamespaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   workflow.NamespaceTable,
			Columns: []string{workflow.NamespaceColumn},
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
		_node.namespace_workflows = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.RevisionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.RevisionsTable,
			Columns: []string{workflow.RevisionsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.RefsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.RefsTable,
			Columns: []string{workflow.RefsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: ref.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.InstancesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.InstancesTable,
			Columns: []string{workflow.InstancesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.RoutesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.RoutesTable,
			Columns: []string{workflow.RoutesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: route.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := wc.mutation.LogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.LogsTable,
			Columns: []string{workflow.LogsColumn},
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
	if nodes := wc.mutation.VarsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.VarsTable,
			Columns: []string{workflow.VarsColumn},
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
	if nodes := wc.mutation.WfeventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   workflow.WfeventsTable,
			Columns: []string{workflow.WfeventsColumn},
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

// WorkflowCreateBulk is the builder for creating many Workflow entities in bulk.
type WorkflowCreateBulk struct {
	config
	builders []*WorkflowCreate
}

// Save creates the Workflow entities in the database.
func (wcb *WorkflowCreateBulk) Save(ctx context.Context) ([]*Workflow, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Workflow, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WorkflowMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WorkflowCreateBulk) SaveX(ctx context.Context) []*Workflow {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WorkflowCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WorkflowCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}
