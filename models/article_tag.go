package models

import (
	"encoding/json"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// ArticleTag is used by pop to map your article_tags database table to your go code.
type ArticleTag struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Tag       Tag       `belongs_to:"tag"`
	TagID     uuid.UUID `db:"tag_id"`
	Article   Article   `belongs_to:"article"`
	ArticleID uuid.UUID `db:"article_id"`
}

// String is not required by pop and may be deleted
func (a ArticleTag) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// ArticleTags is not required by pop and may be deleted
type ArticleTags []ArticleTag

// String is not required by pop and may be deleted
func (a ArticleTags) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Create an article tag
func (a *ArticleTag) Create(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndCreate(a)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *ArticleTag) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *ArticleTag) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *ArticleTag) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
