package slack

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type Message struct {
	Channel  string `json:"channel"`
	Username string `json:"username,omitempty"`
	Text     string `json:"text"`
	Icon     string `json:"icon_emoji,omitempty"`
}

func (m *Message) Send(slackApi string) error {
	if len(m.Channel) == 0 {
		return errors.New("Must specify a slack channel!")
	}

	if len(m.Text) == 0 {
		return errors.New("Must specifty text for slack message!")
	}

	//added for those who forget or don't want to use the hashtag prefix
	if !strings.HasPrefix(m.Channel, "#") {
		m.Channel = "#" + m.Channel
	}

	js, err := json.Marshal(m)
	if err != nil {
		return err
	}

	resp, err := http.PostForm(slackApi, url.Values{
		"payload": {string(js)},
	})
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
