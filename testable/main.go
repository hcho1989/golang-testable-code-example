package main

import (
	"github.com/hcho1989/golang-testable-code-example/testable/slack"
	"github.com/sirupsen/logrus"
)

// App is now decoupled from the message client and logger
type App struct {
	MsgClient MessageClientInterface
	Logger    LoggerInterface
}

// MessageClientInterface defines the message sending interface
type MessageClientInterface interface {
	SendMessage(channel, message string) (interface{}, error)
}

// LoggerInterface defines the message logging interface
type LoggerInterface interface {
	Infoln(...interface{})
}

// Start logs some messages, and send a message to a channel
func (c *App) Start() error {

	c.Logger.Infoln("----- App Start!")
	c.Logger.Infoln("----- Try to Send Message")

	result, err := c.MsgClient.SendMessage("some-channel", "Hello World!")
	if err != nil {
		c.Logger.Infoln("----- Fail to send message", err)
		return err
	}
	c.Logger.Infoln("----- Message Sent, result: ", result)
	return nil
}

func main() {
	sender := slack.SlackClient{Key: "some-api-key"}
	logger := logrus.New()
	app := &App{
		MsgClient: &sender,
		Logger:    logger,
	}

	app.Start()
}
