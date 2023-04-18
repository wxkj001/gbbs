package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("group_id"),              //用户组uuid
		field.Float("integral").Default(0), //积分
		field.String("password").Sensitive(),
		field.String("nick_name"),
		field.String("email"),
		field.String("phone"),
		field.String("avatar").Default("/1"),
		field.Enum("status").Values("on", "off").Default("on"),         //用户状态
		field.Int64("reg_time").Default(time.Now().Unix()).Comment(""), //注册时间
	}

}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("topics", Topic.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("topics", Topic.Type),
		//edge.From("topics", Topic.Type).Ref("users").Unique().Required().Field("user_uuid"),
	}
}
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
		index.Fields("nick_name"),
		index.Fields("email"),
		index.Fields("phone"),
	}
}
