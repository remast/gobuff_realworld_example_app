package models

import (
	"encoding/json"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// ArticleFavorite is used by pop to map your article_favorites database table to your go code.
type ArticleFavorite struct {
	ID        uuid.UUID `json:"id" db:"id"`
	User      User      `belongs_to:"user"`
	UserID    uuid.UUID `db:"user_id"`
	Article   Article   `belongs_to:"article"`
	ArticleID uuid.UUID `db:"article_id"`
}

// String is not required by pop and may be deleted
func (a ArticleFavorite) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// ArticleFavorites is not required by pop and may be deleted
type ArticleFavorites []ArticleFavorite

// String is not required by pop and may be deleted
func (a ArticleFavorites) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Create an article with slug
func (a *ArticleFavorite) Create(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndCreate(a)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *ArticleFavorite) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *ArticleFavorite) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *ArticleFavorite) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
