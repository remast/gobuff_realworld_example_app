package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
	"github.com/gosimple/slug"
	"github.com/pkg/errors"
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
	ArticleTags      []ArticleTag      `has_many:"tags" fk_id:"article_id"`
	Tags             string            `json:"-" db:"-"`
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

// ParseTags parses the string of tags
func (a *Article) ParseTags() []string {
	tagsString := strings.ToLower(strings.Replace(a.Tags, "#", "", 0))

	// Try with separator ','
	tagsToNormalize := strings.Split(tagsString, ",")
	if len(tagsToNormalize) == 1 {
		// Try with separator ' '
		tagsToNormalize = strings.Split(tagsString, " ")
	}

	tagsUnique := map[string]bool{}
	tagsNormalized := []string{}
	for _, tagToNormalize := range tagsToNormalize {
		_, ok := tagsUnique[tagToNormalize] // check for existence
		if ok {
			continue
		}

		tagsUnique[tagToNormalize] = true // add tag
		tagsNormalized = append(tagsNormalized, strings.TrimSpace(tagToNormalize))
	}
	return tagsNormalized
}

// Create an article with slug
func (a *Article) Create(tx *pop.Connection) (*validate.Errors, error) {
	a.Slug = slug.Make(a.Title)
	verrs, err := tx.ValidateAndCreate(a)
	if err != nil {
		return verrs, errors.WithStack(err)
	}
	if verrs.HasAny() {
		return verrs, err
	}

	err = a.updateTags(tx)
	return verrs, err
}

func (a *Article) updateTags(tx *pop.Connection) error {
	tags, err := LoadOrCreateTags(tx, a.ParseTags())
	if err != nil {
		return errors.WithStack(err)
	}

	articleTags := []ArticleTag{}
	for _, tag := range tags {
		articleTag := &ArticleTag{
			ArticleID: a.ID,
			TagID:     tag.ID,
		}
		articleTags = append(articleTags, *articleTag)
	}

	// 1. Delete all tags of this article
	q := tx.RawQuery("delete from article_tags where article_id = ?", a.ID)
	err = q.Exec()
	if err != nil {
		return errors.WithStack(err)
	}

	// 2. Insert all tags of this article
	for _, articleTag := range articleTags {
		_, err = articleTag.Create(tx)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

// Destroy an article
func (a *Article) Destroy(tx *pop.Connection) error {
	// Delete all tags of this article
	q := tx.RawQuery("delete from article_tags where article_id = ?", a.ID)
	err := q.Exec()
	if err != nil {
		return errors.WithStack(err)
	}

	// Delete all favorites of this article
	q = tx.RawQuery("delete from article_favorites where article_id = ?", a.ID)
	err = q.Exec()
	if err != nil {
		return errors.WithStack(err)
	}

	return tx.Destroy(a)
}

// Update an article with slug
func (a *Article) Update(tx *pop.Connection) (*validate.Errors, error) {
	a.Slug = slug.Make(a.Title)
	err := a.updateTags(tx)
	if err != nil {
		return nil, errors.WithStack(err)
	}
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
