package models

func (ms *ModelSuite) Test_ArticleTag_LoadWithTag() {
	// Arrange
	ms.LoadFixture("basics")

	article := &Article{}

	// Act
	err := ms.DB.Eager("ArticleTags").Eager("ArticleTags.Tag").First(article)

	// Assert
	ms.NoError(err)
	ms.Equal(1, len(article.ArticleTags))

	tag := article.ArticleTags[0].Tag
	ms.Equal("beginner", tag.Name)
}

func (ms *ModelSuite) Test_ArticleTag_AddTag() {
	// Arrange
	ms.LoadFixture("basics")

	newTag := &Tag{
		Name: "new_tag",
	}
	_, err := newTag.Create(ms.DB)
	ms.NoError(err)

	article := &Article{}
	err = ms.DB.Eager("ArticleTags").First(article)
	ms.NoError(err)

	ms.Equal(1, len(article.ArticleTags))

	newArticleTag := &ArticleTag{
		ArticleID: article.ID,
		Article:   *article,
		TagID:     newTag.ID,
		Tag:       *newTag,
	}

	// Act
	_, err = newArticleTag.Create(ms.DB)

	// Assert
	ms.NoError(err)
	articleWithNewTag := &Article{}
	err = ms.DB.Eager("ArticleTags").Eager("ArticleTags.Tag").Find(articleWithNewTag, article.ID)
	ms.NoError(err)

	ms.Equal(2, len(articleWithNewTag.ArticleTags))

	ms.Equal("beginner", articleWithNewTag.ArticleTags[0].Tag.Name)
	ms.Equal(newTag.Name, articleWithNewTag.ArticleTags[1].Tag.Name)
}

func (ms *ModelSuite) Test_ArticleFavorite_PopularTags() {
	// Arrange
	ms.LoadFixture("basics")

	// Act
	tags, err := LoadPopularArticleTags(ms.DB, 10)

	// Assert
	ms.NoError(err)
	ms.Equal(1, len(tags))
}
