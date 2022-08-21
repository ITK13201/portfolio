package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Language holds the schema definition for the Language entity.
type Language struct {
	ent.Schema
}

// Annotations of the Language.
func (Language) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "languages"},
	}
}

// Fields of the Language.
func (Language) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable(),
		field.String("name").NotEmpty().Unique().StructTag("binding:\"required\""),
		field.Int64("image_id").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Language.
func (Language) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("image", Image.Type).Field("image_id").Unique(),
	}
}
