package actions

import (
	"gobuff_realworld_example_app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	a := []models.Article{}
	tx := c.Value("tx").(*pop.Connection)

	q := tx.PaginateFromParams(c.Params())
	q.Order("created_at desc").Eager("User").Eager("ArticleFavorites").All(&a)

	c.Set("paginator", q.Paginator)

	c.Logger().Error(q.Paginator.String())

	// article not found so redirect to home
	if len(a) == 0 {
		return c.Redirect(302, "/")
	}

	c.Set("source_page", c.Request().URL)
	c.Set("articles", a)

	return c.Render(200, r.HTML("index.html"))
}
