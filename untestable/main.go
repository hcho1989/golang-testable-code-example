package main

import (
	"github.com/nlopes/slack"
	"github.com/sirupsen/logrus"
)

// App sends slack message
type App struct{}

var logger = logrus.New()         // Unmockable!!
var srv = slack.New("some-token") // Unmockable!!

// Start logs some messages, and send message through slack
func (c *App) Start() error {
	logger.Infoln("----- App Start!")
	logger.Infoln("----- Connect external service")

	x, y, z, err := srv.SendMessage("Hello World!")
	if err != nil {
		logger.Infoln("----- Fail to connect external service", err)
		return err
	}
	logger.Infoln("----- Message Sent, returns: ", x, y, z)
	return nil
}

func main() {
	app := &App{}
	app.Start()
}
