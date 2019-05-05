package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/appleboy/drone-facebook/template"
)

type (
	// GitHub information.
	GitHub struct {
		Workflow  string
		Workspace string
		Action    string
		EventName string
		EventPath string
	}

	// Repo information
	Repo struct {
		FullName  string
		Namespace string
		Name      string
	}

	// Build information
	Build struct {
		Tag      string
		Event    string
		Number   int
		Commit   string
		RefSpec  string
		Branch   string
		Author   string
		Avatar   string
		Message  string
		Email    string
		Status   string
		Link     string
		Started  float64
		Finished float64
	}

	// Config for the plugin.
	Config struct {
		Url     string
		UserID  string
		Token   string
		Message []string
		Drone   bool
		GitHub  bool
	}

	// Payload struct
	Payload struct {
		Channel string `json:"channel"`
		Text    string `json:"text"`
		Avatar  string `json:"avatar"`
	}

	// Plugin values.
	Plugin struct {
		GitHub  GitHub
		Repo    Repo
		Build   Build
		Config  Config
		Payload Payload
	}
)

func trimElement(keys []string) []string {
	var newKeys []string

	for _, value := range keys {
		value = strings.Trim(value, " ")
		if len(value) == 0 {
			continue
		}
		newKeys = append(newKeys, value)
	}

	return newKeys
}

// Exec executes the plugin.
func (p *Plugin) Exec() error {
	if p.Config.Url == "" || p.Config.UserID == "" || p.Config.Token == "" {
		return errors.New("missing Rocket.Chat config")
	}
	var message []string
	if len(p.Config.Message) > 0 {
		message = p.Config.Message
	} else {
		message = p.Message()
	}

	message = trimElement(message)

	for _, value := range message {
		txt, err := template.RenderTrim(value, p)
		if err != nil {
			return err
		}

		err = p.SendMessage(txt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p Plugin) Message() []string {
	if p.Config.GitHub {
		return []string{fmt.Sprintf("%s/%s triggered by %s (%s)",
			p.Repo.FullName,
			p.GitHub.Workflow,
			p.Repo.Namespace,
			p.GitHub.EventName,
		)}
	}

	return []string{fmt.Sprintf("[%s] <%s> (%s)『%s』by %s",
		p.Build.Status,
		p.Build.Link,
		p.Build.Branch,
		p.Build.Message,
		p.Build.Author,
	)}
}

func (p *Plugin) SendMessage(msg string) error {
	URL := fmt.Sprintf("%s/api/v1/chat.postMessage", p.Config.Url)
	b := new(bytes.Buffer)
	var payload = p.Payload
	payload.Text = msg
	if err := json.NewEncoder(b).Encode(payload); err != nil {
		return err
	}
	log.Println(URL, b)
	client := &http.Client{}
	req, err := http.NewRequest("POST", URL, b)
	if err != nil {
		return err
	}
	req.Header.Add("X-User-Id", p.Config.UserID)
	req.Header.Add("X-Auth-Token", p.Config.Token)
	resp, err := client.Do(req)
	log.Println(resp)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	return nil
}
