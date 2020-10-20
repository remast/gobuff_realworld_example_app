package models

func (ms *ModelSuite) Test_Article() {
	ms.LoadFixture("basics")

	u := &User{}
	ms.DB.Where("email = ?", "sarah@sample.de").First(u)

	ms.Failf("BAM", "UID: %v", u)
}
