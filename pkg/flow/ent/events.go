// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/direktiv/direktiv/pkg/flow/ent/events"
	"github.com/direktiv/direktiv/pkg/flow/ent/instance"
	"github.com/direktiv/direktiv/pkg/flow/ent/namespace"
	"github.com/direktiv/direktiv/pkg/flow/ent/workflow"
	"github.com/google/uuid"
)

// Events is the model entity for the Events schema.
type Events struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id"`
	// Events holds the value of the "events" field.
	Events []map[string]interface{} `json:"events,omitempty"`
	// Correlations holds the value of the "correlations" field.
	Correlations []string `json:"correlations,omitempty"`
	// Signature holds the value of the "signature" field.
	Signature []byte `json:"signature,omitempty"`
	// Count holds the value of the "count" field.
	Count int `json:"count,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EventsQuery when eager-loading is set.
	Edges                        EventsEdges `json:"edges"`
	instance_eventlisteners      *uuid.UUID
	namespace_namespacelisteners *uuid.UUID
	workflow_wfevents            *uuid.UUID
}

// EventsEdges holds the relations/edges for other nodes in the graph.
type EventsEdges struct {
	// Workflow holds the value of the workflow edge.
	Workflow *Workflow `json:"workflow,omitempty"`
	// Wfeventswait holds the value of the wfeventswait edge.
	Wfeventswait []*EventsWait `json:"wfeventswait,omitempty"`
	// Instance holds the value of the instance edge.
	Instance *Instance `json:"instance,omitempty"`
	// Namespace holds the value of the namespace edge.
	Namespace *Namespace `json:"namespace,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// WorkflowOrErr returns the Workflow value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EventsEdges) WorkflowOrErr() (*Workflow, error) {
	if e.loadedTypes[0] {
		if e.Workflow == nil {
			// The edge workflow was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: workflow.Label}
		}
		return e.Workflow, nil
	}
	return nil, &NotLoadedError{edge: "workflow"}
}

// WfeventswaitOrErr returns the Wfeventswait value or an error if the edge
// was not loaded in eager-loading.
func (e EventsEdges) WfeventswaitOrErr() ([]*EventsWait, error) {
	if e.loadedTypes[1] {
		return e.Wfeventswait, nil
	}
	return nil, &NotLoadedError{edge: "wfeventswait"}
}

// InstanceOrErr returns the Instance value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EventsEdges) InstanceOrErr() (*Instance, error) {
	if e.loadedTypes[2] {
		if e.Instance == nil {
			// The edge instance was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: instance.Label}
		}
		return e.Instance, nil
	}
	return nil, &NotLoadedError{edge: "instance"}
}

// NamespaceOrErr returns the Namespace value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EventsEdges) NamespaceOrErr() (*Namespace, error) {
	if e.loadedTypes[3] {
		if e.Namespace == nil {
			// The edge namespace was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: namespace.Label}
		}
		return e.Namespace, nil
	}
	return nil, &NotLoadedError{edge: "namespace"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Events) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case events.FieldEvents, events.FieldCorrelations, events.FieldSignature:
			values[i] = new([]byte)
		case events.FieldCount:
			values[i] = new(sql.NullInt64)
		case events.FieldCreatedAt, events.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case events.FieldID:
			values[i] = new(uuid.UUID)
		case events.ForeignKeys[0]: // instance_eventlisteners
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case events.ForeignKeys[1]: // namespace_namespacelisteners
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case events.ForeignKeys[2]: // workflow_wfevents
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Events", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Events fields.
func (e *Events) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case events.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				e.ID = *value
			}
		case events.FieldEvents:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field events", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &e.Events); err != nil {
					return fmt.Errorf("unmarshal field events: %w", err)
				}
			}
		case events.FieldCorrelations:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field correlations", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &e.Correlations); err != nil {
					return fmt.Errorf("unmarshal field correlations: %w", err)
				}
			}
		case events.FieldSignature:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field signature", values[i])
			} else if value != nil {
				e.Signature = *value
			}
		case events.FieldCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field count", values[i])
			} else if value.Valid {
				e.Count = int(value.Int64)
			}
		case events.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case events.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				e.UpdatedAt = value.Time
			}
		case events.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field instance_eventlisteners", values[i])
			} else if value.Valid {
				e.instance_eventlisteners = new(uuid.UUID)
				*e.instance_eventlisteners = *value.S.(*uuid.UUID)
			}
		case events.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field namespace_namespacelisteners", values[i])
			} else if value.Valid {
				e.namespace_namespacelisteners = new(uuid.UUID)
				*e.namespace_namespacelisteners = *value.S.(*uuid.UUID)
			}
		case events.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_wfevents", values[i])
			} else if value.Valid {
				e.workflow_wfevents = new(uuid.UUID)
				*e.workflow_wfevents = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryWorkflow queries the "workflow" edge of the Events entity.
func (e *Events) QueryWorkflow() *WorkflowQuery {
	return (&EventsClient{config: e.config}).QueryWorkflow(e)
}

// QueryWfeventswait queries the "wfeventswait" edge of the Events entity.
func (e *Events) QueryWfeventswait() *EventsWaitQuery {
	return (&EventsClient{config: e.config}).QueryWfeventswait(e)
}

// QueryInstance queries the "instance" edge of the Events entity.
func (e *Events) QueryInstance() *InstanceQuery {
	return (&EventsClient{config: e.config}).QueryInstance(e)
}

// QueryNamespace queries the "namespace" edge of the Events entity.
func (e *Events) QueryNamespace() *NamespaceQuery {
	return (&EventsClient{config: e.config}).QueryNamespace(e)
}

// Update returns a builder for updating this Events.
// Note that you need to call Events.Unwrap() before calling this method if this Events
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Events) Update() *EventsUpdateOne {
	return (&EventsClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Events entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Events) Unwrap() *Events {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Events is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Events) String() string {
	var builder strings.Builder
	builder.WriteString("Events(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("events=")
	builder.WriteString(fmt.Sprintf("%v", e.Events))
	builder.WriteString(", ")
	builder.WriteString("correlations=")
	builder.WriteString(fmt.Sprintf("%v", e.Correlations))
	builder.WriteString(", ")
	builder.WriteString("signature=")
	builder.WriteString(fmt.Sprintf("%v", e.Signature))
	builder.WriteString(", ")
	builder.WriteString("count=")
	builder.WriteString(fmt.Sprintf("%v", e.Count))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(e.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// EventsSlice is a parsable slice of Events.
type EventsSlice []*Events

func (e EventsSlice) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
