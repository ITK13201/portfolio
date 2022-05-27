package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"time"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Annotations of the Image.
func (Image) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "images"},
	}
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable(),
		field.String("slug").NotEmpty().Unique().StructTag("binding:\"required\""),
		field.String("path").NotEmpty().StructTag("binding:\"required\""),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return nil
}
