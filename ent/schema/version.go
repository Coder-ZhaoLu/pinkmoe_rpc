package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

// Version holds the schema definition for the Version entity.
type Version struct {
	ent.Schema
}

func (Version) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("Title | 标题"),
		field.Text("content").Comment("Content | 内容"),
		field.Uint64("service_uuid").Comment("ServiceUuid | 服务UUID"),
		field.String("url").Comment("Url | 服务链接"),
	}
}

func (Version) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}

func (Version) Edges() []ent.Edge {
	return nil
}
