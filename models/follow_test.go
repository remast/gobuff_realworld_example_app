package models

func (ms *ModelSuite) Test_Follow() {
	// Arrange
	ms.LoadFixture("basics")

	u1 := &User{}
	ms.DB.Where("email = ?", "sarah@sample.de").First(u1)

	u2 := &User{}
	ms.DB.Where("email = ?", "max@sample.de").First(u2)

	follow := &Follow{
		UserID:   u2.ID,
		FollowID: u1.ID,
	}

	// Act
	verrs, err := follow.Create(ms.DB)

	// Assert
	ms.NoError(err)
	ms.False(verrs.HasAny())

	u1WithFollowers := &User{}
	ms.DB.Where("email = ?", "sarah@sample.de").Eager("Followers").First(u1WithFollowers)

	ms.Equal(1, len(u1WithFollowers.Followers))
}
