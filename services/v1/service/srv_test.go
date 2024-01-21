package service_test

import (
	"be-cadidate-votes/adapter/repositories"
	"be-cadidate-votes/appconfig"
	"be-cadidate-votes/services/v1/service"
	"log"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"

	mockRepo "be-cadidate-votes/mocks/repositories"
)

type TestServiceSuite struct {
	suite.Suite
	ctrl *gomock.Controller

	mockTbRepo         *mockRepo.MockTbRepo
	mockThirdPartyRepo *mockRepo.MockThirdPartyRepo
	mockCfg            *appconfig.AppConfig

	srv service.Service
}

func TestService(t *testing.T) {
	suite.Run(t, new(TestServiceSuite))
}

func (t *TestServiceSuite) BeforeTest(suiteName string, testName string) {
	log.Println(`[TESTNAME]:  `, testName)
	t.ctrl = gomock.NewController(t.T())

	t.mockTbRepo = mockRepo.NewMockTbRepo(t.ctrl)
	t.mockThirdPartyRepo = mockRepo.NewMockThirdPartyRepo(t.ctrl)
	t.mockCfg = &appconfig.AppConfig{}

	t.srv = service.NewService(t.mockCfg, &repositories.Repo{
		TbRepo:         t.mockTbRepo,
		ThirdPartyRepo: t.mockThirdPartyRepo,
	})
}

func (t *TestServiceSuite) AfterTest(suiteName string, testName string) {
	t.ctrl.Finish()
}
