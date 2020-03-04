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
		result := "i don't care what this is"
		mMsgClient := mocks.MessageClientInterface{}
		mLogger := mocks.LoggerInterface{}
		mLogger.On("Infoln", mock.AnythingOfType("string")).Return()
		mMsgClient.On("SendMessage", "some-channel", "Hello World!").Return(result, nil)
		mLogger.On("Infoln", "----- Message Sent, result: ", result).Return()

		app := &App{
			MsgClient: &mMsgClient,
			Logger:    &mLogger, // You can also generate mock struct for LoggerInterface
		}
		err := app.Start()
		mMsgClient.AssertCalled(t, "SendMessage", "some-channel", "Hello World!")
		mLogger.AssertCalled(t, "Infoln", "----- App Start!")
		mLogger.AssertCalled(t, "Infoln", "----- Try to Send Message")
		mLogger.AssertCalled(t, "Infoln", "----- Message Sent, result: ", result)

		assert.Nil(t, err)
	})
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}
