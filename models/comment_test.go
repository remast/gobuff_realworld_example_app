package models

func (ms *ModelSuite) Test_Comment() {
	// Arrange
	ms.LoadFixture("basics")

	article := &Article{}
	ms.DB.First(article)

	comment := &Comment{
		UserID:    article.UserID,
		ArticleID: article.ID,
		Body:      "My Comment",
	}

	// Act
	verrs, err := comment.Create(ms.DB)

	// Assert
	ms.NoError(err)
	ms.False(verrs.HasAny())

	comments := []Comment{}
	ms.DB.Where("article_id = ?", article.ID).Order("created_at desc").Limit(20).Eager().All(&comments)

	ms.Equal(1, len(comments))
}
