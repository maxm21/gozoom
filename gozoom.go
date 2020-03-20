package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
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
	Name,
	AccountID string
}

type MessageHead struct {
	Text    string `json:"text,omitempty"`
	SubHead *Text  `json:"sub_head,omitempty"`
}

type Text struct {
	Text string `json:"text,omitempty"`
}

type MessageBody struct {
	Type              string                `json:"type,omitempty"`
	SidebarColor      string                `json:"sidebar_color,omitempty"`
	Sections          []*MessageBodySection `json:"sections,omitempty"`
	Footer            string                `json:"footer,omitempty"`
	AttResourceURL string                `json:"resource_url,omitempty"`
	AttImgURL      string                `json:"img_url,omitempty"`
	AttInformation *AttData     `json:"information,omitempty"`
}

type AttData struct {
	Title       *Text `json:"title,omitempty"`
	Description *Text `json:"description,omitempty"`
}

type MessageBodySection struct {
	Type  string             `json:"type,omitempty"`
	Text  string             `json:"text,omitempty"`
	Items []*MessageBodyItem `json:"items,omitempty"`
}

type MessageBodyItem struct {
	Key      string `json:"key,omitempty"`
	Value    string `json:"value,omitempty"`
	Editable bool   `json:"editable,omitempty"`
}

type Message struct {
	Head *MessageHead   `json:"head,omitempty"`
	Body []*MessageBody `json:"body,omitempty"`
}

func (b *Bot) authorize() error {
	auth := b64.StdEncoding.EncodeToString([]byte(b.ClientID + ":" + b.ClientSecret)) // encode the clientid:clientsecret in base64

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

	b.Bearer = response.AccessToken
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

func (u *User) sendComplexMsg(m *Message, b *Bot) (string, error) {
	client := &http.Client{}

	var payload struct {
		RobotJid          string   `json:"robot_jid"`
		ToJid             string   `json:"to_jid"`
		AccountID         string   `json:"account_id"`
		IsMarkdownSupport bool     `json:"is_markdown_support"`
		Message           *Message `json:"content"`
	}

	payload.RobotJid = b.BotJID
	payload.ToJid = u.ToJid
	payload.AccountID = u.AccountID
	payload.IsMarkdownSupport = true
	payload.Message = m
	//working on making this more concise please please excuse the lines above :)

	pload, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "https://api.zoom.us/v2/im/chat/messages", strings.NewReader(string(pload)))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer " + b.Bearer)

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

func main() {
	bot := &Bot{
		AuthCode:     "XXXX",
		ClientID:     "XXXX",
		ClientSecret: "XXXX",
		RedirectURL:  "XXXX",
		BotJID:       "XXXX"}
	user := &User{
		ToJid:     "XXXX",
		Name:      "XXXX",
		AccountID: "XXXX",
	}

	err := bot.authorize()
	if err == nil {
		m := &Message{ //this is an example of how to make a complex message (embed, if you will)
			Head: &MessageHead{
				Text: "title can be _italic_ ~or not~",
				SubHead: &Text{
					"SUBHEADER, <#" + user.ToJid + "|" + user.Name + ">",
				},
			},
			Body: []*MessageBody{
				&MessageBody{ //First body
					Type:         "section",
					SidebarColor: "#F56416",
					Sections: []*MessageBodySection{
						&MessageBodySection{
							Type: "message",
							Text: "*first message*",
						},
						&MessageBodySection{
							Type: "message",
							Text: "second message <img:https://images-na.ssl-images-amazon.com/images/I/51Mt-I6%2BEQL._AC_SX466_.jpg>",
						},
					},
					Footer: "FOOTER TEST",
				},
				&MessageBody{ //Attachment demo
					Type:        "attachments",
					AttResourceURL: "https://golang.org",
					AttImgURL: "https://images-na.ssl-images-amazon.com/images/I/51Mt-I6%2BEQL._AC_SX466_.jpg",
					AttInformation: &AttData{
						Title: &Text{"Gopher"},
						Description: &Text{"Golang is so cool!"},
					},
				},
			},
		}

		mID, err := user.sendComplexMsg(m, bot)
		// mID, err := user.sendSimpleMsg("hi", bot)

		if err != nil {
			fmt.Println("Error sending message:", err)
		} else {
			fmt.Println("Message sent successfully with id", mID)
		}

	} else {
		fmt.Println(err, "Failed to get bearer token. Invalid credentials?\nProgram closing.")
		os.Exit(1)
	}
}
