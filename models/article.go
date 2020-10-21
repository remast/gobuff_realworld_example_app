package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
	"github.com/gosimple/slug"
)

// Article is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Article struct {
	ID               uuid.UUID         `json:"id" db:"id"`
	Title            string            `json:"title" db:"title"`
	Slug             string            `json:"slug" db:"slug"`
	Description      string            `json:"description" db:"description"`
	Body             string            `json:"body" db:"body"`
	CreatedAt        time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time         `json:"updated_at" db:"updated_at"`
	User             User              `belongs_to:"user"`
	UserID           uuid.UUID         `db:"user_id"`
	ArticleFavorites []ArticleFavorite `has_many:"favorites" fk_id:"article_id"`
}

// String is not required by pop and may be deleted
func (a Article) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Articles is not required by pop and may be deleted
type Articles []Article

// String is not required by pop and may be deleted
func (a Articles) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Create an article with slug
func (a *Article) Create(tx *pop.Connection) (*validate.Errors, error) {
	a.Slug = slug.Make(a.Title)
	return tx.ValidateAndCreate(a)
}

// Update an article with slug
func (a *Article) Update(tx *pop.Connection) (*validate.Errors, error) {
	a.Slug = slug.Make(a.Title)
	return tx.ValidateAndUpdate(a)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *Article) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: a.Title, Name: "Title"},
		&validators.StringIsPresent{Field: a.Description, Name: "Description"},
		&validators.StringIsPresent{Field: a.Body, Name: "Body"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Article) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Article) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
