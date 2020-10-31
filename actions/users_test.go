package actions

import (
	"gobuff_realworld_example_app/models"
)

func (as *ActionSuite) Test_Users_Register() {
	res := as.HTML("/users/register").Get()
	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_Users_Create() {
	// Arrange
	count, err := as.DB.Count("users")
	as.NoError(err)
	as.Equal(0, count)

	u := &models.User{
		Name:     "Mark Example",
		Email:    "mark@example.com",
		Password: "password",
	}

	// Act
	res := as.HTML("/users/register").Post(u)

	// Assert
	as.Equal(302, res.Code)

	count, err = as.DB.Count("users")
	as.NoError(err)
	as.Equal(1, count)
}
