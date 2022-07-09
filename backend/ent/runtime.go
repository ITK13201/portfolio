// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/ITK13201/portfolio/backend/ent/abouttopic"
	"github.com/ITK13201/portfolio/backend/ent/image"
	"github.com/ITK13201/portfolio/backend/ent/schema"
	"github.com/ITK13201/portfolio/backend/ent/user"
	"github.com/ITK13201/portfolio/backend/ent/work"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	abouttopicFields := schema.AboutTopic{}.Fields()
	_ = abouttopicFields
	// abouttopicDescTitle is the schema descriptor for title field.
	abouttopicDescTitle := abouttopicFields[1].Descriptor()
	// abouttopic.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	abouttopic.TitleValidator = abouttopicDescTitle.Validators[0].(func(string) error)
	// abouttopicDescDescriptionJp is the schema descriptor for description_jp field.
	abouttopicDescDescriptionJp := abouttopicFields[2].Descriptor()
	// abouttopic.DescriptionJpValidator is a validator for the "description_jp" field. It is called by the builders before save.
	abouttopic.DescriptionJpValidator = abouttopicDescDescriptionJp.Validators[0].(func(string) error)
	// abouttopicDescDescriptionEn is the schema descriptor for description_en field.
	abouttopicDescDescriptionEn := abouttopicFields[3].Descriptor()
	// abouttopic.DescriptionEnValidator is a validator for the "description_en" field. It is called by the builders before save.
	abouttopic.DescriptionEnValidator = abouttopicDescDescriptionEn.Validators[0].(func(string) error)
	// abouttopicDescCreatedAt is the schema descriptor for created_at field.
	abouttopicDescCreatedAt := abouttopicFields[6].Descriptor()
	// abouttopic.DefaultCreatedAt holds the default value on creation for the created_at field.
	abouttopic.DefaultCreatedAt = abouttopicDescCreatedAt.Default.(func() time.Time)
	// abouttopicDescUpdatedAt is the schema descriptor for updated_at field.
	abouttopicDescUpdatedAt := abouttopicFields[7].Descriptor()
	// abouttopic.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	abouttopic.DefaultUpdatedAt = abouttopicDescUpdatedAt.Default.(func() time.Time)
	// abouttopic.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	abouttopic.UpdateDefaultUpdatedAt = abouttopicDescUpdatedAt.UpdateDefault.(func() time.Time)
	imageFields := schema.Image{}.Fields()
	_ = imageFields
	// imageDescSlug is the schema descriptor for slug field.
	imageDescSlug := imageFields[1].Descriptor()
	// image.SlugValidator is a validator for the "slug" field. It is called by the builders before save.
	image.SlugValidator = imageDescSlug.Validators[0].(func(string) error)
	// imageDescPath is the schema descriptor for path field.
	imageDescPath := imageFields[2].Descriptor()
	// image.PathValidator is a validator for the "path" field. It is called by the builders before save.
	image.PathValidator = imageDescPath.Validators[0].(func(string) error)
	// imageDescCreatedAt is the schema descriptor for created_at field.
	imageDescCreatedAt := imageFields[3].Descriptor()
	// image.DefaultCreatedAt holds the default value on creation for the created_at field.
	image.DefaultCreatedAt = imageDescCreatedAt.Default.(func() time.Time)
	// imageDescUpdatedAt is the schema descriptor for updated_at field.
	imageDescUpdatedAt := imageFields[4].Descriptor()
	// image.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	image.DefaultUpdatedAt = imageDescUpdatedAt.Default.(func() time.Time)
	// image.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	image.UpdateDefaultUpdatedAt = imageDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescHashedPassword is the schema descriptor for hashed_password field.
	userDescHashedPassword := userFields[2].Descriptor()
	// user.HashedPasswordValidator is a validator for the "hashed_password" field. It is called by the builders before save.
	user.HashedPasswordValidator = userDescHashedPassword.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[3].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	workFields := schema.Work{}.Fields()
	_ = workFields
	// workDescTitle is the schema descriptor for title field.
	workDescTitle := workFields[1].Descriptor()
	// work.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	work.TitleValidator = workDescTitle.Validators[0].(func(string) error)
	// workDescDescriptionJp is the schema descriptor for description_jp field.
	workDescDescriptionJp := workFields[2].Descriptor()
	// work.DescriptionJpValidator is a validator for the "description_jp" field. It is called by the builders before save.
	work.DescriptionJpValidator = workDescDescriptionJp.Validators[0].(func(string) error)
	// workDescDescriptionEn is the schema descriptor for description_en field.
	workDescDescriptionEn := workFields[3].Descriptor()
	// work.DescriptionEnValidator is a validator for the "description_en" field. It is called by the builders before save.
	work.DescriptionEnValidator = workDescDescriptionEn.Validators[0].(func(string) error)
	// workDescLink is the schema descriptor for link field.
	workDescLink := workFields[5].Descriptor()
	// work.LinkValidator is a validator for the "link" field. It is called by the builders before save.
	work.LinkValidator = workDescLink.Validators[0].(func(string) error)
	// workDescCreatedAt is the schema descriptor for created_at field.
	workDescCreatedAt := workFields[7].Descriptor()
	// work.DefaultCreatedAt holds the default value on creation for the created_at field.
	work.DefaultCreatedAt = workDescCreatedAt.Default.(func() time.Time)
	// workDescUpdatedAt is the schema descriptor for updated_at field.
	workDescUpdatedAt := workFields[8].Descriptor()
	// work.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	work.DefaultUpdatedAt = workDescUpdatedAt.Default.(func() time.Time)
	// work.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	work.UpdateDefaultUpdatedAt = workDescUpdatedAt.UpdateDefault.(func() time.Time)
}