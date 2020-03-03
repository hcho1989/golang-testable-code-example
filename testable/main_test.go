package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/hcho1989/golang-testable-code-example/testable/mocks"
)

type MainTestSuite struct {
	suite.Suite
}

func (suite *MainTestSuite) TestApp_Start() {
	t := suite.T()

	t.Run("App.Start shall have called service.Connect and service.SendMessage", func(t *testing.T) {
		mSrv := mocks.ExternalServiceInterface{}
		mLogger := mocks.LoggerInterface{}
		mLogger.On("Infoln", mock.AnythingOfType("string")).Return()
		mSrv.On("SendMessage", "Hello World!").Return("x", "y", "z", nil)
		mLogger.On("Infoln", "----- Message Sent, returns: ", "x", "y", "z").Return()

		app := &App{
			ExtSrv: &mSrv,
			Logger: &mLogger, // You can also generate mock struct for LoggerInterface
		}
		err := app.Start()
		mSrv.AssertCalled(t, "SendMessage", "Hello World!")
		mLogger.AssertCalled(t, "Infoln", "----- App Start!")
		mLogger.AssertCalled(t, "Infoln", "----- Connect external service")
		mLogger.AssertCalled(t, "Infoln", "----- Message Sent, returns: ", "x", "y", "z")

		assert.Nil(t, err)
	})
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}
