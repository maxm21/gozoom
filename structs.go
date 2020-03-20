package main

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

