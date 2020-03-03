package main

import (
	"github.com/nlopes/slack"
	"github.com/sirupsen/logrus"
)

// App sends slack message
type App struct {
	ExtSrv ExternalServiceInterface
	Logger LoggerInterface
}

// ExternalServiceInterface defines the message sending interface
type ExternalServiceInterface interface {
	SendMessage(channel string, options ...slack.MsgOption) (string, string, string, error)
}

// LoggerInterface defines the message logging interface
type LoggerInterface interface {
	Infoln(...interface{})
}

// Start logs some messages, and send message through slack
func (c *App) Start() error {

	c.Logger.Infoln("----- App Start!")
	c.Logger.Infoln("----- Connect external service")

	x, y, z, err := c.ExtSrv.SendMessage("Hello World!")
	if err != nil {
		c.Logger.Infoln("----- Fail to connect external service", err)
		return err
	}
	c.Logger.Infoln("----- Message Sent, returns: ", x, y, z)
	return nil
}

func main() {
	key := "some-api-key"

	srv := slack.New(key)
	app := &App{
		ExtSrv: srv,
		Logger: logrus.New(),
	}

	app.Start()
}
