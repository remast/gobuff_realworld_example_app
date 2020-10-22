package actions

import (
	"gobuff_realworld_example_app/models"
	"net/http"
)

func (as *ActionSuite) readUser() *models.User {
	as.LoadFixture("basics")

	u := &models.User{}
	as.DB.Where("email = ?", "sarah@sample.de").First(u)

	return u
}

func (as *ActionSuite) Test_Auth_Signin() {
	res := as.HTML("/auth/login").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), `<button class="btn btn-success">Sign In!</button>`)
}

func (as *ActionSuite) Test_Auth_Register() {
	res := as.HTML("/users/register").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "Sign In")
}

func (as *ActionSuite) Test_Auth_SignUp() {
	tcases := []struct {
		Name     string
		Email    string
		Password string
		Status   int

		Identifier string
	}{
		{"Mia Mice", "mia@mice.com", "cat", http.StatusFound, "Valid"},
		{"My Maik", "my.maik@test.com", "", http.StatusOK, "Password Invalid"},
	}

	for _, tcase := range tcases {
		as.Run(tcase.Identifier, func() {
			res := as.HTML("/users/register").Post(&models.User{
				Name:     tcase.Name,
				Email:    tcase.Email,
				Password: tcase.Password,
			})

			as.Equal(tcase.Status, res.Code)
		})
	}
}
