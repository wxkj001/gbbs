package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Topic holds the schema definition for the Topic entity.
type Topic struct {
	ent.Schema
}

/*func (Topic) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_uuid", "category_uuid"),
	}
}*/

// Fields of the Topic.
func (Topic) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_uid"),
		field.Int("category_id").Optional(),  //类别
		field.Strings("tags").Optional(),     //标签
		field.String("title"),                //标题
		field.String("content"),              //内容
		field.Time("create_time"),            //创建日期
		field.Time("edit_time"),              //编辑日期
		field.Time("delete_time").Optional(), //删除日期
	}
}

// Edges of the Topic.
func (Topic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("users", User.Type).
			Ref("topics").
			Unique().Required().Field("user_uid"),
		//edge.To("users", User.Type),
	}
}
func (Topic) Indexes() []ent.Index {
	return []ent.Index{
		//index.Fields("user_uid").Edges("users"),
		index.Fields("category_id"),
		index.Fields("title"),
	}
}
