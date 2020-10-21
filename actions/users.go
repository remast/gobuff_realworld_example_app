package actions

import (
	"gobuff_realworld_example_app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gofrs/uuid"
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
	c.Set("source_page", c.Request().URL)
	c.Set("profile_user", user)

	a := []models.Article{}

	q := tx.PaginateFromParams(c.Params())
	q.Where("user_id = ?", user.ID).Order("created_at desc").Limit(10).Eager("User").Eager("ArticleFavorites").All(&a)

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

// UsersFollow creates a follow relation
func UsersFollow(c buffalo.Context) error {
	userID := c.Value("current_user").(*models.User).ID
	followID := uuid.FromStringOrNil(c.Request().Form.Get("FollowID"))

	follow := &models.Follow{}
	tx := c.Value("tx").(*pop.Connection)
	found, err := tx.Where("user_id = ? and follow_id = ?", userID, followID).Exists(follow)
	if err != nil {
		return errors.WithStack(err)
	}

	if found {
		follow = &models.Follow{}
		tx.Where("user_id = ? and follow_id = ?", userID, followID).First(follow)
		err = tx.Destroy(follow)
		if err != nil {
			return errors.WithStack(err)
		}
	} else {
		follow = &models.Follow{
			UserID:   userID,
			FollowID: followID,
		}

		_, err := follow.Create(tx)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	sourcePage := c.Request().Form.Get("SourcePage")
	return c.Redirect(302, sourcePage)
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
			return c.Redirect(302, "/auth/login")
		}
		return next(c)
	}
}
