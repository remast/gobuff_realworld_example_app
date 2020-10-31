package models

import (
	"encoding/json"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

// Tag is used by pop to map your tags database table to your go code.
type Tag struct {
	ID   uuid.UUID `json:"id" db:"id"`
	Name string    `json:"name" db:"name"`
}

// String is not required by pop and may be deleted
func (t Tag) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Tags is not required by pop and may be deleted
type Tags []Tag

// String is not required by pop and may be deleted
func (t Tags) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// LoadOrCreateTags loads given tags or creates them if not present yet
func LoadOrCreateTags(tx *pop.Connection, tagNames []string) ([]Tag, error) {
	tags := []Tag{}
	err := tx.Where("name in (?)", tagNames).All(&tags)
	if err != nil {
		return nil, err
	}

	for _, tagName := range tagNames {

		tagPresent := false
		for _, tag := range tags {
			if tag.Name == tagName {
				tagPresent = true
				break
			}
		}

		if tagPresent {
			continue
		}

		// create Tag
		t := &Tag{
			Name: tagName,
		}

		_, err := t.Create(tx)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		tags = append(tags, *t)
	}

	return tags, err
}

// Create a tag
func (t *Tag) Create(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndCreate(t)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Tag) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Name, Name: "Name"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Tag) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Tag) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
