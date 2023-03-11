package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Comment("Name | 名称"),
		field.String("slug").Comment("Slug | 标识"),
		field.String("icon").Comment("Icon | 图标"),
		field.String("desc").Comment("Desc | 描述"),
		field.Uint32("type").Default(1).Comment("Type | 类型 1: 微服务 2: 组件库"),
	}
}

func (Category) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
		mixins.SortMixin{},
		mixins.StatusMixin{},
	}
}

func (Category) Edges() []ent.Edge {
	return nil
}
