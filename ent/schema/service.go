package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-core/pkg/ent/schema/mixins"
)

// Service holds the schema definition for the Service entity.
type Service struct {
	ent.Schema
}

func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Comment("Title | 标题"),
		field.String("desc").Comment("Desc | 描述"),
		field.String("content").Comment("Content | 内容"),
		field.Uint64("category_id").Comment("CategoryId | 分类ID"),
		field.String("author_uuid").Comment("AuthorUuid | 作者ID"),
		field.String("cover").Comment("Cover | 封面"),
		field.Uint32("type").Default(1).Comment("Type | 类型 1: 微服务 2: 组件库 3: 文章"),
		field.Uint32("price").Comment("Price | 价格"),
		field.Uint64("view").Default(0).Comment("View | 查看数"),
	}
}

func (Service) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UUIDMixin{},
		mixins.StatusMixin{},
	}
}

func (Service) Edges() []ent.Edge {
	return nil
}
