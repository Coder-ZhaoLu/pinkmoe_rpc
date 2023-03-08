package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

// Sitemeta holds the schema definition for the Sitemeta entity.
type Sitemeta struct {
	ent.Schema
}

func (Sitemeta) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Comment("Key | 键"),
		field.String("value").Comment("Value | 值"),
	}
}

func (Sitemeta) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

func (Sitemeta) Edges() []ent.Edge {
	return nil
}
