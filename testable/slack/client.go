package slack

import "github.com/nlopes/slack"

type SlackClient struct {
	Key    string
	client *slack.Client
}

func (s *SlackClient) init() {
	s.client = slack.New(s.Key)
}

func (s *SlackClient) SendMessage(channel, message string) (interface{}, error) {
	x, y, z, err := s.client.SendMessage(channel, slack.MsgOptionText(message, false))
	return SlackMessageResult{X: x, Y: y, Z: z}, err
}
