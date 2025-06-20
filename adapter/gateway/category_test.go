package gateway_test

import (
	"go-api-arch-clean/adapter/gateway"
	"go-api-arch-clean/entity"
	"go-api-arch-clean/pkg/tester"
	"testing"

	"github.com/stretchr/testify/suite"
)

type CategoryRepositorySuite struct {
	tester.DBSQLiteSuite
	repository gateway.CategoryRepository
}

func TestCategoryRepositorySuite(t *testing.T) {
	suite.Run(t, new(CategoryRepositorySuite))
}

func (suite *CategoryRepositorySuite) SetupSuite() {
	suite.DBSQLiteSuite.SetupSuite()
	suite.repository = gateway.NewCategoryRepository(suite.DBSQLiteSuite.DB)
}

func (suite *CategoryRepositorySuite) TestCategory() {
	paramCategory, err := entity.NewCategory("sports")
	suite.Assert().Nil(err)
	category, err := suite.repository.GetOrCreate(paramCategory)
	suite.Assert().Nil(err)
	suite.Assert().Equal(1, category.ID)
	suite.Assert().Equal("sports", string(category.Name))

	category, err = suite.repository.GetOrCreate(paramCategory)
	suite.Assert().Nil(err)
	suite.Assert().Equal(1, category.ID)
	suite.Assert().Equal("sports", string(category.Name))
}
