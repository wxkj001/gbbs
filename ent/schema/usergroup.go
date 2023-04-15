package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// UserGroup holds the schema definition for the UserGroup entity.
type UserGroup struct {
	ent.Schema
}

// Fields of the UserGroup.
func (UserGroup) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("group_uuid", uuid.UUID{}).Default(uuid.New),
		field.String("group_name"),
	}
}

// Edges of the UserGroup.
func (UserGroup) Edges() []ent.Edge {
	return nil
}
