package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Bot struct {
	Bearer,
	AuthCode,
	ClientID,
	ClientSecret,
	RedirectURL,
	BotJID string
}

type User struct {
	ToJid,
	AccountID string
}

type Message struct {
	Head struct {
		Text    string `json:"text,omitempty"`
		SubHead struct {
			Text string `json:"text,omitempty"`
		} `json:"sub_head,omitempty"`
	} `json:"head,omitempty"`
	Body []struct {
		Type     string `json:"type,omitempty"`
		Sections []struct {
			Type  string `json:"type,omitempty"`
			Text  string `json:"text,omitempty"`
			Items []struct {
				Key      string `json:"key,omitempty"`
				Value    string `json:"value,omitempty"`
				Editable bool   `json:"editable,omitempty"`
			} `json:"items,omitempty"`
		} `json:"sections,omitempty"`
		Footer string `json:"footer,omitempty"`
	} `json:"body,omitempty"`
}

func authorize(bot *Bot) error {
	auth := b64.StdEncoding.EncodeToString([]byte(bot.ClientID + ":" + bot.ClientSecret)) // encode the clientid:clientsecret in base64

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.zoom.us/oauth/token?grant_type=client_credentials", nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Basic "+auth)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var response struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
		Scope       string `json:"scope"`
	}
	json.Unmarshal(body, &response)

	bot.Bearer = response.AccessToken
	return nil
}

func (u *User) sendSimpleMsg(m string, bot *Bot) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://api.zoom.us/v2/im/chat/messages", strings.NewReader("{\n    \"robot_jid\": \""+bot.BotJID+"\",\n    \"to_jid\": \""+u.ToJid+"\",\n    \"account_id\": \""+u.AccountID+"\",\n    \"content\": {\n        \"head\": {\n            \"text\": \""+m+"\"\n        }\n    }\n}"))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+bot.Bearer)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var messageResp struct {
		MessageID string `json:"message_id"`
		RobotJid  string `json:"robot_jid"`
		SentTime  string `json:"sent_time"`
		ToJid     string `json:"to_jid"`
	}
	json.Unmarshal(body, &messageResp)

	return messageResp.MessageID, nil
}

func (u *User) sendComplexMsg(m *Message, bot *Bot) (string, error) {
	/*client := &http.Client{}

	type MessageSend struct {
	RobotJid          string `json:"robot_jid"`
	ToJid             string `json:"to_jid"`
	AccountID         string `json:"account_id"`
	IsMarkdownSupport bool   `json:"is_markdown_support"`
	Message
}

	req, err := http.NewRequest("POST", "https://api.zoom.us/v2/im/chat/messages", strings.NewReader("{\n    \"robot_jid\": \""+bot.BotJID+"\",\n    \"to_jid\": \""+u.ToJid+"\",\n    \"account_id\": \""+u.AccountID+"\",\n    \"content\": {\n        \"head\": {\n            \"text\": \""+m+"\"\n        }\n    }\n}"))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer " + bot.Bearer)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var messageResp struct {
		MessageID string `json:"message_id"`
		RobotJid  string `json:"robot_jid"`
		SentTime  string `json:"sent_time"`
		ToJid     string `json:"to_jid"`
	}
	json.Unmarshal(body, &messageResp)

	return messageResp.MessageID, nil*/
	return "", nil
}

func main() {
	bot := &Bot{
		AuthCode:     "XXXX",
		ClientID:     "XXXX",
		ClientSecret: "XXXX",
		RedirectURL:  "XXXX",
		BotJID:       "XXXX"}
	user := &User{
		ToJid:     "XXXX",
		AccountID: "XXXX",
	}

	err := authorize(bot)
	if err != nil {
		m := &Message{
			Head: ,
			Body: nil,
		}

		fmt.Println(m)

		/*
    mID, err := user.sendSimpleMsg("strring", bot)

		if err != nil {
			fmt.Println("Error sending message:", err)
		} else {
			fmt.Println("Message sent successfully with id", mID)
		}
    */

	} else {
		fmt.Println("Failed to get bearer token. Invalid credentials?\nProgram closing.")
		os.Exit(1)
	}
}
