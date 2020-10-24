package models

func (ms *ModelSuite) Test_Tag_Create() {
	// Arrange
	countBefore, err := ms.DB.Count("tags")
	ms.NoError(err)

	t := &Tag{
		Name: "mytag",
	}

	// Act
	verrs, err := t.Create(ms.DB)

	// Assert
	ms.False(verrs.HasAny())
	ms.NoError(err)

	count, err := ms.DB.Count("tags")
	ms.NoError(err)
	ms.Equal(countBefore+1, count)
}
