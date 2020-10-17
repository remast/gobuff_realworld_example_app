package actions

import (
	"gobuff_realworld_example_app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"
)

//UsersRegister renders the users form
func UsersRegister(c buffalo.Context) error {
	u := models.User{}
	c.Set("user", u)
	return c.Render(200, r.HTML("users/register.html"))
}

//UsersProfile renders the user profile
func UsersProfile(c buffalo.Context) error {
	email := c.Param("user_email")

	u := []models.User{}
	tx := c.Value("tx").(*pop.Connection)
	tx.Where("email = ?", email).All(&u)

	// user not found so redirect to home
	if len(u) == 0 {
		return c.Redirect(302, "/")
	}

	user := u[0]
	c.Set("profile_user", user)

	a := []models.Article{}
	tx.Where("user_id = ?", user.ID).Order("created_at desc").Limit(10).Eager().All(&a)

	c.Set("articles", a)

	return c.Render(200, r.HTML("users/profile.html"))
}

// UsersCreate registers a new user with the application.
func UsersCreate(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := u.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("user", u)
		c.Set("errors", verrs)
		return c.Render(200, r.HTML("users/register.html"))
	}

	c.Session().Set("current_user_id", u.ID)
	c.Flash().Add("success", "Welcome to Buffalo!")

	return c.Redirect(302, "/")
}

// SetCurrentUser attempts to find a user based on the current_user_id
// in the session. If one is found it is set on the context.
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				return errors.WithStack(err)
			}
			c.Set("current_user", u)
		}
		return next(c)
	}
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Session().Set("redirectURL", c.Request().URL.String())

			err := c.Session().Save()
			if err != nil {
				return errors.WithStack(err)
			}

			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Redirect(302, "/auth/new")
		}
		return next(c)
	}
}
