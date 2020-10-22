package actions

import (
	"fmt"
	"gobuff_realworld_example_app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

// ArticlesRead renders the article
func ArticlesRead(c buffalo.Context) error {
	slug := c.Param("slug")

	a := []models.Article{}
	tx := c.Value("tx").(*pop.Connection)
	tx.Where("slug = ?", slug).Eager("ArticleFavorites").All(&a)

	// article not found so redirect to home
	if len(a) == 0 {
		return c.Redirect(302, "/")
	}

	article := &a[0]

	c.Set("source_page", c.Request().URL)
	c.Set("article", article)
	c.Set("comment", &models.Comment{})

	comments := []models.Comment{}
	tx.Where("article_id = ?", article.ID).Order("created_at desc").Limit(20).Eager().All(&comments)
	c.Set("comments", comments)

	author := &models.User{}
	tx.Eager("Followers").Find(author, article.UserID)
	c.Set("author", author)

	return c.Render(200, r.HTML("articles/read.html"))
}

// ArticlesComment renders the article
func ArticlesComment(c buffalo.Context) error {
	u := c.Value("current_user").(*models.User)
	slug := c.Param("slug")

	comment := &models.Comment{}

	if err := c.Bind(comment); err != nil {
		return errors.WithStack(err)
	}

	if comment.Body == "" {
		return c.Redirect(302, fmt.Sprintf("/articles/%v", slug))
	}

	a := []models.Article{}
	tx := c.Value("tx").(*pop.Connection)
	tx.Where("slug = ?", slug).All(&a)

	// article not found so redirect to home
	if len(a) == 0 {
		return c.Redirect(302, "/")
	}

	article := a[0]

	comment.UserID = u.ID
	comment.ArticleID = article.ID

	_, err := comment.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	return c.Redirect(302, fmt.Sprintf("/articles/%v", slug))
}

// ArticlesDelete deletes an article
func ArticlesDelete(c buffalo.Context) error {
	u := c.Value("current_user").(*models.User)
	slug := c.Param("slug")

	a := []models.Article{}

	tx := c.Value("tx").(*pop.Connection)
	tx.Where("slug = ? and user_id = ?", slug, u.ID).Eager().All(&a)

	if len(a) > 0 {
		tx.Destroy(&a)
		c.Flash().Add("success", "Article deleted")
	}

	return c.Redirect(302, "/")
}

// ArticlesNew renders the article form
func ArticlesNew(c buffalo.Context) error {
	a := models.Article{}
	c.Set("article", a)
	return c.Render(200, r.HTML("articles/new.html"))
}

// ArticlesStar stars an article
func ArticlesStar(c buffalo.Context) error {
	userID := c.Value("current_user").(*models.User).ID
	articleID := uuid.FromStringOrNil(c.Request().Form.Get("ArticleID"))

	articleFavorite := &models.ArticleFavorite{}
	tx := c.Value("tx").(*pop.Connection)
	found, err := tx.Where("user_id = ? and article_id = ?", userID, articleID).Exists(articleFavorite)
	if err != nil {
		return errors.WithStack(err)
	}

	if found {
		articleFavorite = &models.ArticleFavorite{}
		tx.Where("user_id = ? and article_id = ?", userID, articleID).First(articleFavorite)
		err = tx.Destroy(articleFavorite)
		if err != nil {
			return errors.WithStack(err)
		}
	} else {
		articleFavorite = &models.ArticleFavorite{
			UserID:    userID,
			ArticleID: articleID,
		}

		_, err := articleFavorite.Create(tx)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	sourcePage := c.Request().Form.Get("SourcePage")
	return c.Redirect(302, sourcePage)
}

// ArticlesEdit renders the edit article form
func ArticlesEdit(c buffalo.Context) error {
	u := c.Value("current_user").(*models.User)
	slug := c.Param("slug")

	a := []models.Article{}

	tx := c.Value("tx").(*pop.Connection)
	tx.Where("slug = ? and user_id = ?", slug, u.ID).Eager().All(&a)

	if len(a) == 0 {
		return c.Redirect(302, "/")
	}

	c.Set("article", a[0])

	return c.Render(200, r.HTML("articles/edit.html"))
}

// ArticlesCreate creates a new article
func ArticlesCreate(c buffalo.Context) error {
	u := c.Value("current_user").(*models.User)

	a := &models.Article{}
	a.UserID = u.ID

	if err := c.Bind(a); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := a.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("article", a)
		c.Set("errors", verrs)
		return c.Render(200, r.HTML("articles/new.html"))
	}

	c.Flash().Add("success", "Article created")

	return c.Redirect(302, fmt.Sprintf("/articles/%v", a.Slug))
}

// ArticlesUpdate updates an article
func ArticlesUpdate(c buffalo.Context) error {
	u := c.Value("current_user").(*models.User)
	slug := c.Param("slug")

	tx := c.Value("tx").(*pop.Connection)
	article := &models.Article{}
	article.UserID = u.ID

	if err := c.Bind(article); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := article.Update(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("article", article)
		c.Set("errors", verrs)
		return c.Redirect(302, fmt.Sprintf("/articles/%v/edit", slug))
	}

	c.Flash().Add("success", "Article updated")

	return c.Redirect(302, fmt.Sprintf("/articles/%v", article.Slug))
}
