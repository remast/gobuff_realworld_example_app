package actions

import (
	"gobuff_realworld_example_app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	a := []models.Article{}
	tx := c.Value("tx").(*pop.Connection)
	tx.Order("created_at desc").Eager().Limit(10).All(&a)

	// article not found so redirect to home
	if len(a) == 0 {
		return c.Redirect(302, "/")
	}

	c.Logger().Error(a[0].User)

	c.Set("articles", a)

	return c.Render(200, r.HTML("index.html"))
}
