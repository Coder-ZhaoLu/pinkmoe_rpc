// Code generated by ent, DO NOT EDIT.

package comment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Coder-ZhaoLu/pinkmoe_rpc/ent/predicate"
	uuid "github.com/gofrs/uuid/v5"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldStatus, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldContent, v))
}

// ServiceUUID applies equality check predicate on the "service_uuid" field. It's identical to ServiceUUIDEQ.
func ServiceUUID(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldServiceUUID, v))
}

// UserUUID applies equality check predicate on the "user_uuid" field. It's identical to UserUUIDEQ.
func UserUUID(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUserUUID, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v uint32) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldType, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Comment {
	return predicate.Comment(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Comment {
	return predicate.Comment(sql.FieldNotNull(FieldStatus))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldContent, v))
}

// ServiceUUIDEQ applies the EQ predicate on the "service_uuid" field.
func ServiceUUIDEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldServiceUUID, v))
}

// ServiceUUIDNEQ applies the NEQ predicate on the "service_uuid" field.
func ServiceUUIDNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldServiceUUID, v))
}

// ServiceUUIDIn applies the In predicate on the "service_uuid" field.
func ServiceUUIDIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldServiceUUID, vs...))
}

// ServiceUUIDNotIn applies the NotIn predicate on the "service_uuid" field.
func ServiceUUIDNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldServiceUUID, vs...))
}

// ServiceUUIDGT applies the GT predicate on the "service_uuid" field.
func ServiceUUIDGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldServiceUUID, v))
}

// ServiceUUIDGTE applies the GTE predicate on the "service_uuid" field.
func ServiceUUIDGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldServiceUUID, v))
}

// ServiceUUIDLT applies the LT predicate on the "service_uuid" field.
func ServiceUUIDLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldServiceUUID, v))
}

// ServiceUUIDLTE applies the LTE predicate on the "service_uuid" field.
func ServiceUUIDLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldServiceUUID, v))
}

// ServiceUUIDContains applies the Contains predicate on the "service_uuid" field.
func ServiceUUIDContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldServiceUUID, v))
}

// ServiceUUIDHasPrefix applies the HasPrefix predicate on the "service_uuid" field.
func ServiceUUIDHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldServiceUUID, v))
}

// ServiceUUIDHasSuffix applies the HasSuffix predicate on the "service_uuid" field.
func ServiceUUIDHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldServiceUUID, v))
}

// ServiceUUIDEqualFold applies the EqualFold predicate on the "service_uuid" field.
func ServiceUUIDEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldServiceUUID, v))
}

// ServiceUUIDContainsFold applies the ContainsFold predicate on the "service_uuid" field.
func ServiceUUIDContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldServiceUUID, v))
}

// UserUUIDEQ applies the EQ predicate on the "user_uuid" field.
func UserUUIDEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUserUUID, v))
}

// UserUUIDNEQ applies the NEQ predicate on the "user_uuid" field.
func UserUUIDNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldUserUUID, v))
}

// UserUUIDIn applies the In predicate on the "user_uuid" field.
func UserUUIDIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldUserUUID, vs...))
}

// UserUUIDNotIn applies the NotIn predicate on the "user_uuid" field.
func UserUUIDNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldUserUUID, vs...))
}

// UserUUIDGT applies the GT predicate on the "user_uuid" field.
func UserUUIDGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldUserUUID, v))
}

// UserUUIDGTE applies the GTE predicate on the "user_uuid" field.
func UserUUIDGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldUserUUID, v))
}

// UserUUIDLT applies the LT predicate on the "user_uuid" field.
func UserUUIDLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldUserUUID, v))
}

// UserUUIDLTE applies the LTE predicate on the "user_uuid" field.
func UserUUIDLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldUserUUID, v))
}

// UserUUIDContains applies the Contains predicate on the "user_uuid" field.
func UserUUIDContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldUserUUID, v))
}

// UserUUIDHasPrefix applies the HasPrefix predicate on the "user_uuid" field.
func UserUUIDHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldUserUUID, v))
}

// UserUUIDHasSuffix applies the HasSuffix predicate on the "user_uuid" field.
func UserUUIDHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldUserUUID, v))
}

// UserUUIDEqualFold applies the EqualFold predicate on the "user_uuid" field.
func UserUUIDEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldUserUUID, v))
}

// UserUUIDContainsFold applies the ContainsFold predicate on the "user_uuid" field.
func UserUUIDContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldUserUUID, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v uint32) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v uint32) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...uint32) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...uint32) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v uint32) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v uint32) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v uint32) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v uint32) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldType, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
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
func Not(p predicate.Comment) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		p(s.Not())
	})
}
