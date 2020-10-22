package models

func (ms *ModelSuite) Test_Article() {
	// Arrange
	ms.LoadFixture("basics")

	u := &User{}
	ms.DB.Where("email = ?", "sarah@sample.de").First(u)

	countBefore, _ := ms.DB.Count(&Article{})

	article := &Article{
		UserID:      u.ID,
		Title:       "Title",
		Body:        "Body",
		Description: "Description",
	}

	// Act
	verrs, err := article.Create(ms.DB)

	// Assert
	ms.NoError(err)
	ms.False(verrs.HasAny())

	countAfter, _ := ms.DB.Count(&Article{})
	ms.Equal(countBefore+1, countAfter)
}
