package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_uuid", uuid.UUID{}).Default(uuid.New),
		field.UUID("group_uuid", uuid.UUID{}), //用户组uuid
		field.Float("integral").Default(0),    //积分
		field.String("password"),
		field.String("nick_name"),
		field.String("email"),
		field.String("phone"),
		field.String("avatar").Default("/1"),
		field.Enum("status").Values("on", "off").Default("on"), //用户状态
		field.Int64("reg_time").Default(time.Now().Unix()),     //注册时间
	}

}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
