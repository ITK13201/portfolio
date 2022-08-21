// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/ITK13201/portfolio/backend/ent/image"
	"github.com/ITK13201/portfolio/backend/ent/work"
)

// Work is the model entity for the Work schema.
type Work struct {
	config `binding:"-" json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty" binding:"required"`
	// DescriptionJp holds the value of the "description_jp" field.
	DescriptionJp string `json:"description_jp,omitempty" binding:"required"`
	// DescriptionEn holds the value of the "description_en" field.
	DescriptionEn string `json:"description_en,omitempty" binding:"required"`
	// LanguageID holds the value of the "language_id" field.
	LanguageID int64 `json:"language_id,omitempty"`
	// Link holds the value of the "link" field.
	Link string `json:"link,omitempty" binding:"required"`
	// Priority holds the value of the "priority" field.
	Priority int `json:"priority,omitempty" binding:"required"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkQuery when eager-loading is set.
	Edges WorkEdges `json:"edges"`
}

// WorkEdges holds the relations/edges for other nodes in the graph.
type WorkEdges struct {
	// Language holds the value of the language edge.
	Language *Image `json:"language,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// LanguageOrErr returns the Language value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkEdges) LanguageOrErr() (*Image, error) {
	if e.loadedTypes[0] {
		if e.Language == nil {
			// The edge language was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: image.Label}
		}
		return e.Language, nil
	}
	return nil, &NotLoadedError{edge: "language"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Work) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case work.FieldID, work.FieldLanguageID, work.FieldPriority:
			values[i] = new(sql.NullInt64)
		case work.FieldTitle, work.FieldDescriptionJp, work.FieldDescriptionEn, work.FieldLink:
			values[i] = new(sql.NullString)
		case work.FieldCreatedAt, work.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Work", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Work fields.
func (w *Work) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case work.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int64(value.Int64)
		case work.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				w.Title = value.String
			}
		case work.FieldDescriptionJp:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description_jp", values[i])
			} else if value.Valid {
				w.DescriptionJp = value.String
			}
		case work.FieldDescriptionEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description_en", values[i])
			} else if value.Valid {
				w.DescriptionEn = value.String
			}
		case work.FieldLanguageID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field language_id", values[i])
			} else if value.Valid {
				w.LanguageID = value.Int64
			}
		case work.FieldLink:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field link", values[i])
			} else if value.Valid {
				w.Link = value.String
			}
		case work.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				w.Priority = int(value.Int64)
			}
		case work.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				w.CreatedAt = value.Time
			}
		case work.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				w.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// QueryLanguage queries the "language" edge of the Work entity.
func (w *Work) QueryLanguage() *ImageQuery {
	return (&WorkClient{config: w.config}).QueryLanguage(w)
}

// Update returns a builder for updating this Work.
// Note that you need to call Work.Unwrap() before calling this method if this Work
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Work) Update() *WorkUpdateOne {
	return (&WorkClient{config: w.config}).UpdateOne(w)
}

// Unwrap unwraps the Work entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Work) Unwrap() *Work {
	tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Work is not a transactional entity")
	}
	w.config.driver = tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Work) String() string {
	var builder strings.Builder
	builder.WriteString("Work(")
	builder.WriteString(fmt.Sprintf("id=%v", w.ID))
	builder.WriteString(", title=")
	builder.WriteString(w.Title)
	builder.WriteString(", description_jp=")
	builder.WriteString(w.DescriptionJp)
	builder.WriteString(", description_en=")
	builder.WriteString(w.DescriptionEn)
	builder.WriteString(", language_id=")
	builder.WriteString(fmt.Sprintf("%v", w.LanguageID))
	builder.WriteString(", link=")
	builder.WriteString(w.Link)
	builder.WriteString(", priority=")
	builder.WriteString(fmt.Sprintf("%v", w.Priority))
	builder.WriteString(", created_at=")
	builder.WriteString(w.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(w.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Works is a parsable slice of Work.
type Works []*Work

func (w Works) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}
