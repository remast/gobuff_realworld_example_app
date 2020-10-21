package models

func (ms *ModelSuite) Test_ArticleFavorite() {
	// Arrange
	ms.LoadFixture("basics")

	article := &Article{}
	ms.DB.First(article)

	articleFavorite := &ArticleFavorite{
		UserID:    article.UserID,
		ArticleID: article.ID,
	}

	// Act
	verrs, err := articleFavorite.Create(ms.DB)

	// Assert
	ms.NoError(err)
	ms.False(verrs.HasAny())

	articleWithFav := &Article{}
	ms.DB.Eager("ArticleFavorites").Find(articleWithFav, article.ID)
	ms.Equal(1, len(articleWithFav.ArticleFavorites))
}
