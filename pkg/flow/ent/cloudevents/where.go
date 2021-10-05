// Code generated by entc, DO NOT EDIT.

package cloudevents

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/vorteil/direktiv/pkg/flow/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EventId applies equality check predicate on the "eventId" field. It's identical to EventIdEQ.
func EventId(v string) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventId), v))
	})
}

// Fire applies equality check predicate on the "fire" field. It's identical to FireEQ.
func Fire(v time.Time) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFire), v))
	})
}

// Created applies equality check predicate on the "created" field. It's identical to CreatedEQ.
func Created(v time.Time) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// Processed applies equality check predicate on the "processed" field. It's identical to ProcessedEQ.
func Processed(v bool) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProcessed), v))
	})
}

// EventIdEQ applies the EQ predicate on the "eventId" field.
func EventIdEQ(v string) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEventId), v))
	})
}

// EventIdNEQ applies the NEQ predicate on the "eventId" field.
func EventIdNEQ(v string) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEventId), v))
	})
}

// EventIdIn applies the In predicate on the "eventId" field.
func EventIdIn(vs ...string) predicate.CloudEvents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CloudEvents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEventId), v...))
	})
}

// EventIdNotIn applies the NotIn predicate on the "eventId" field.
func EventIdNotIn(vs ...string) predicate.CloudEvents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CloudEvents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEventId), v...))
	})
}

// EventIdGT applies the GT predicate on the "eventId" field.
func EventIdGT(v string) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEventId), v))
	})
}

// EventIdGTE applies the GTE predicate on the "eventId" field.
func EventIdGTE(v string) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEventId), v))
	})
}

// EventIdLT applies the LT predicate on the "eventId" field.
func EventIdLT(v string) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEventId), v))
	})
}

// EventIdLTE applies the LTE predicate on the "eventId" field.
func EventIdLTE(v string) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEventId), v))
	})
}

// EventIdContains applies the Contains predicate on the "eventId" field.
func EventIdContains(v string) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEventId), v))
	})
}

// EventIdHasPrefix applies the HasPrefix predicate on the "eventId" field.
func EventIdHasPrefix(v string) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEventId), v))
	})
}

// EventIdHasSuffix applies the HasSuffix predicate on the "eventId" field.
func EventIdHasSuffix(v string) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEventId), v))
	})
}

// EventIdEqualFold applies the EqualFold predicate on the "eventId" field.
func EventIdEqualFold(v string) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEventId), v))
	})
}

// EventIdContainsFold applies the ContainsFold predicate on the "eventId" field.
func EventIdContainsFold(v string) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEventId), v))
	})
}

// FireEQ applies the EQ predicate on the "fire" field.
func FireEQ(v time.Time) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFire), v))
	})
}

// FireNEQ applies the NEQ predicate on the "fire" field.
func FireNEQ(v time.Time) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFire), v))
	})
}

// FireIn applies the In predicate on the "fire" field.
func FireIn(vs ...time.Time) predicate.CloudEvents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CloudEvents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFire), v...))
	})
}

// FireNotIn applies the NotIn predicate on the "fire" field.
func FireNotIn(vs ...time.Time) predicate.CloudEvents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CloudEvents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFire), v...))
	})
}

// FireGT applies the GT predicate on the "fire" field.
func FireGT(v time.Time) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFire), v))
	})
}

// FireGTE applies the GTE predicate on the "fire" field.
func FireGTE(v time.Time) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFire), v))
	})
}

// FireLT applies the LT predicate on the "fire" field.
func FireLT(v time.Time) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFire), v))
	})
}

// FireLTE applies the LTE predicate on the "fire" field.
func FireLTE(v time.Time) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFire), v))
	})
}

// CreatedEQ applies the EQ predicate on the "created" field.
func CreatedEQ(v time.Time) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreated), v))
	})
}

// CreatedNEQ applies the NEQ predicate on the "created" field.
func CreatedNEQ(v time.Time) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreated), v))
	})
}

// CreatedIn applies the In predicate on the "created" field.
func CreatedIn(vs ...time.Time) predicate.CloudEvents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CloudEvents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreated), v...))
	})
}

// CreatedNotIn applies the NotIn predicate on the "created" field.
func CreatedNotIn(vs ...time.Time) predicate.CloudEvents {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CloudEvents(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreated), v...))
	})
}

// CreatedGT applies the GT predicate on the "created" field.
func CreatedGT(v time.Time) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreated), v))
	})
}

// CreatedGTE applies the GTE predicate on the "created" field.
func CreatedGTE(v time.Time) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreated), v))
	})
}

// CreatedLT applies the LT predicate on the "created" field.
func CreatedLT(v time.Time) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreated), v))
	})
}

// CreatedLTE applies the LTE predicate on the "created" field.
func CreatedLTE(v time.Time) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreated), v))
	})
}

// ProcessedEQ applies the EQ predicate on the "processed" field.
func ProcessedEQ(v bool) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProcessed), v))
	})
}

// ProcessedNEQ applies the NEQ predicate on the "processed" field.
func ProcessedNEQ(v bool) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProcessed), v))
	})
}

// HasNamespace applies the HasEdge predicate on the "namespace" edge.
func HasNamespace() predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNamespaceWith applies the HasEdge predicate on the "namespace" edge with a given conditions (other predicates).
func HasNamespaceWith(preds ...predicate.Namespace) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NamespaceInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NamespaceTable, NamespaceColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CloudEvents) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CloudEvents) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CloudEvents) predicate.CloudEvents {
	return predicate.CloudEvents(func(s *sql.Selector) {
		p(s.Not())
	})
}