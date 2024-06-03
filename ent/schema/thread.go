package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/lucsky/cuid"
)

type Thread struct {
	ent.Schema
}

func (Thread) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(cuid.New),
		field.String("name").NotEmpty(),
		field.String("description"),
	}
}

func (Thread) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("messages", Message.Type),
	}
}
