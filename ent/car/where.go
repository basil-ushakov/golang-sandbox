// Code generated by ent, DO NOT EDIT.

package car

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/NICKNAME-wengreen/BigDemo/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Car {
	return predicate.Car(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Car {
	return predicate.Car(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Car {
	return predicate.Car(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Car {
	return predicate.Car(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Car {
	return predicate.Car(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Car {
	return predicate.Car(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Car {
	return predicate.Car(sql.FieldLTE(FieldID, id))
}

// Model applies equality check predicate on the "model" field. It's identical to ModelEQ.
func Model(v string) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldModel, v))
}

// RegisteredAt applies equality check predicate on the "registered_at" field. It's identical to RegisteredAtEQ.
func RegisteredAt(v time.Time) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldRegisteredAt, v))
}

// ModelEQ applies the EQ predicate on the "model" field.
func ModelEQ(v string) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldModel, v))
}

// ModelNEQ applies the NEQ predicate on the "model" field.
func ModelNEQ(v string) predicate.Car {
	return predicate.Car(sql.FieldNEQ(FieldModel, v))
}

// ModelIn applies the In predicate on the "model" field.
func ModelIn(vs ...string) predicate.Car {
	return predicate.Car(sql.FieldIn(FieldModel, vs...))
}

// ModelNotIn applies the NotIn predicate on the "model" field.
func ModelNotIn(vs ...string) predicate.Car {
	return predicate.Car(sql.FieldNotIn(FieldModel, vs...))
}

// ModelGT applies the GT predicate on the "model" field.
func ModelGT(v string) predicate.Car {
	return predicate.Car(sql.FieldGT(FieldModel, v))
}

// ModelGTE applies the GTE predicate on the "model" field.
func ModelGTE(v string) predicate.Car {
	return predicate.Car(sql.FieldGTE(FieldModel, v))
}

// ModelLT applies the LT predicate on the "model" field.
func ModelLT(v string) predicate.Car {
	return predicate.Car(sql.FieldLT(FieldModel, v))
}

// ModelLTE applies the LTE predicate on the "model" field.
func ModelLTE(v string) predicate.Car {
	return predicate.Car(sql.FieldLTE(FieldModel, v))
}

// ModelContains applies the Contains predicate on the "model" field.
func ModelContains(v string) predicate.Car {
	return predicate.Car(sql.FieldContains(FieldModel, v))
}

// ModelHasPrefix applies the HasPrefix predicate on the "model" field.
func ModelHasPrefix(v string) predicate.Car {
	return predicate.Car(sql.FieldHasPrefix(FieldModel, v))
}

// ModelHasSuffix applies the HasSuffix predicate on the "model" field.
func ModelHasSuffix(v string) predicate.Car {
	return predicate.Car(sql.FieldHasSuffix(FieldModel, v))
}

// ModelEqualFold applies the EqualFold predicate on the "model" field.
func ModelEqualFold(v string) predicate.Car {
	return predicate.Car(sql.FieldEqualFold(FieldModel, v))
}

// ModelContainsFold applies the ContainsFold predicate on the "model" field.
func ModelContainsFold(v string) predicate.Car {
	return predicate.Car(sql.FieldContainsFold(FieldModel, v))
}

// RegisteredAtEQ applies the EQ predicate on the "registered_at" field.
func RegisteredAtEQ(v time.Time) predicate.Car {
	return predicate.Car(sql.FieldEQ(FieldRegisteredAt, v))
}

// RegisteredAtNEQ applies the NEQ predicate on the "registered_at" field.
func RegisteredAtNEQ(v time.Time) predicate.Car {
	return predicate.Car(sql.FieldNEQ(FieldRegisteredAt, v))
}

// RegisteredAtIn applies the In predicate on the "registered_at" field.
func RegisteredAtIn(vs ...time.Time) predicate.Car {
	return predicate.Car(sql.FieldIn(FieldRegisteredAt, vs...))
}

// RegisteredAtNotIn applies the NotIn predicate on the "registered_at" field.
func RegisteredAtNotIn(vs ...time.Time) predicate.Car {
	return predicate.Car(sql.FieldNotIn(FieldRegisteredAt, vs...))
}

// RegisteredAtGT applies the GT predicate on the "registered_at" field.
func RegisteredAtGT(v time.Time) predicate.Car {
	return predicate.Car(sql.FieldGT(FieldRegisteredAt, v))
}

// RegisteredAtGTE applies the GTE predicate on the "registered_at" field.
func RegisteredAtGTE(v time.Time) predicate.Car {
	return predicate.Car(sql.FieldGTE(FieldRegisteredAt, v))
}

// RegisteredAtLT applies the LT predicate on the "registered_at" field.
func RegisteredAtLT(v time.Time) predicate.Car {
	return predicate.Car(sql.FieldLT(FieldRegisteredAt, v))
}

// RegisteredAtLTE applies the LTE predicate on the "registered_at" field.
func RegisteredAtLTE(v time.Time) predicate.Car {
	return predicate.Car(sql.FieldLTE(FieldRegisteredAt, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
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
func Not(p predicate.Car) predicate.Car {
	return predicate.Car(func(s *sql.Selector) {
		p(s.Not())
	})
}
