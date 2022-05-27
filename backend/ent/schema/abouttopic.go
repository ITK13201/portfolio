package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// AboutTopic holds the schema definition for the AboutTopic entity.
type AboutTopic struct {
	ent.Schema
}

// Annotations of the AboutTopic.
func (AboutTopic) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "about_topics"},
	}
}

// Fields of the AboutTopic.
func (AboutTopic) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique().Immutable(),
		field.String("title").NotEmpty().Unique().StructTag("binding:\"required\""),
		field.String("description_jp").NotEmpty().StructTag("binding:\"required\""),
		field.String("description_en").NotEmpty().StructTag("binding:\"required\""),
		field.Int64("image_id").Optional(),
		field.Int("priority").StructTag("binding:\"required\""),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the AboutTopic.
func (AboutTopic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("image", Image.Type).Field("image_id").Unique(),
	}
}
