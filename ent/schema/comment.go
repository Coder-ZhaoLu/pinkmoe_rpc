package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").Comment("Content | 评论内容"),
		field.String("service_uuid").Comment("Service Uuid | 服务ID"),
		field.String("user_uuid").Comment("User Uuid | 用户ID"),
		field.Uint32("type").Default(1).Comment("Type | 类型 1: 服务 2: 组件"),
	}
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}

func (Comment) Edges() []ent.Edge {
	return nil
}
