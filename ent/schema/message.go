package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/lucsky/cuid"
	"time"
)

type Message struct {
	ent.Schema
}

func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(cuid.New),
		field.String("text").NotEmpty(),
		field.Time("created").Default(time.Now),
	}
}

func (Message) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("thread", Thread.Type).
			Ref("messages").
			Unique().
			Required(),
	}
}
