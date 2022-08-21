package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Work holds the schema definition for the Work entity.
type Work struct {
	ent.Schema
}

// Annotations of the Work.
func (Work) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "works"},
	}
}

// Fields of the Work.
func (Work) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable(),
		field.String("title").NotEmpty().Unique().StructTag("binding:\"required\""),
		field.String("description_jp").NotEmpty().StructTag("binding:\"required\""),
		field.String("description_en").NotEmpty().StructTag("binding:\"required\""),
		field.Int64("language_id").Optional(),
		field.String("link").NotEmpty().StructTag("binding:\"required\""),
		field.Int("priority").StructTag("binding:\"required\""),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Work.
func (Work) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("language", Image.Type).Field("language_id").Unique(),
	}
}
