package actions

import (
	"gobuff_realworld_example_app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"
)

// HomeHandler is a default handler to serve the home page.
func HomeHandler(c buffalo.Context) error {
	a := []models.Article{}
	tx := c.Value("tx").(*pop.Connection)

	q := tx.PaginateFromParams(c.Params())
	err := q.Order("created_at desc").Eager("User").Eager("ArticleFavorites").Eager("ArticleTags").Eager("ArticleTags.Tag").All(&a)
	if err != nil {
		return errors.WithStack(err)
	}
	c.Set("paginator", q.Paginator)
	c.Set("articles", a)

	tags, err := models.LoadPopularArticleTags(tx, 20)
	if err != nil {
		return errors.WithStack(err)
	}
	c.Set("tags", tags)

	c.Set("source_page", c.Request().URL)

	return c.Render(200, r.HTML("index.html"))
}
