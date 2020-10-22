package actions

import "gobuff_realworld_example_app/models"

func (as *ActionSuite) Test_HomeHandler_LoggedIn() {
	// Arrange
	as.LoadFixture("basics")

	u := &models.User{}
	as.DB.Where("email = ?", "sarah@sample.de").First(u)

	as.Session.Set("current_user_id", u.ID)

	// Act
	res := as.HTML("/articles/new").Get()

	// Assert
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sign Out")

	// Arrange
	as.Session.Clear()

	// Act
	res = as.HTML("/articles/new").Get()

	// Assert
	as.Equal(302, res.Code)
}
