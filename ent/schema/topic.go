package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Topic holds the schema definition for the Topic entity.
type Topic struct {
	ent.Schema
}

// Fields of the Topic.
func (Topic) Fields() []ent.Field {
	return []ent.Field{
		//field.UUID("user_uuid", uuid.UUID{}),
		field.UUID("category_uuid", uuid.UUID{}).Optional(), //类别
		field.Strings("tags").Optional(),                    //标签
		field.String("title"),                               //标题
		field.String("content"),                             //内容
		field.Time("create_time"),                           //创建日期
		field.Time("edit_time"),                             //编辑日期
		field.Time("delete_time").Optional(),                //删除日期
	}
}

// Edges of the Topic.
func (Topic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).Ref("topic").Field("user_uuid"),
	}
}
func (Topic) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_uuid"),
		index.Fields("category_uuid"),
		index.Fields("title"),
	}
}
