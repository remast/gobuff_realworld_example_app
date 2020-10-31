package actions

import (
	"database/sql"
	"net/http"
	"strings"

	"gobuff_realworld_example_app/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// AuthLoginHandler loads the signin page
func AuthLoginHandler(c buffalo.Context) error {
	c.Set("user", models.User{})
	return c.Render(200, r.HTML("auth/login.html"))
}

// AuthCreateHandler attempts to log the user in with an existing account.
func AuthCreateHandler(c buffalo.Context) error {
	u := &models.User{}
	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)

	// find a user by email
	err := tx.Where("email = ?", strings.ToLower(strings.TrimSpace(u.Email))).First(u)

	// helper function to handle bad attempts
	bad := func() error {
		verrs := validate.NewErrors()
		verrs.Add("email", "invalid email/password")

		c.Set("errors", verrs)
		c.Set("user", u)

		return c.Render(http.StatusUnauthorized, r.HTML("auth/login.html"))
	}

	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			// couldn't find an user with the supplied email address.
			return bad()
		}
		return errors.WithStack(err)
	}

	// confirm that the given password matches the hashed password from the db
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
	if err != nil {
		return bad()
	}
	c.Session().Set("current_user_id", u.ID)
	c.Flash().Add("success", "Welcome Back to Buffalo!")

	redirectURL := "/"
	if redir, ok := c.Session().Get("redirectURL").(string); ok && redir != "" {
		redirectURL = redir
	}

	return c.Redirect(302, redirectURL)
}

// AuthLogoutHandler clears the session and logs a user out
func AuthLogoutHandler(c buffalo.Context) error {
	c.Session().Clear()
	c.Flash().Add("success", "You have been logged out!")
	return c.Redirect(302, "/")
}
