package models

import (
	"encoding/json"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Follow is used by pop to map your follows database table to your go code.
type Follow struct {
	ID       uuid.UUID `json:"id" db:"id"`
	UserID   uuid.UUID `db:"user_id"`
	FollowID uuid.UUID `db:"follow_id"`
}

// String is not required by pop and may be deleted
func (f Follow) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Follows is not required by pop and may be deleted
type Follows []Follow

// String is not required by pop and may be deleted
func (f Follows) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Create a follow relation
func (f *Follow) Create(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndCreate(f)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (f *Follow) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (f *Follow) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (f *Follow) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
