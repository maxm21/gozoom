<img align="right" src="https://i.imgur.com/WqW85j1.jpg" height=343px length=303px>

# gozoom

index:
1) [what is gozoom?](#about)
2) [goal](#goal)
3) [how to set up a zoom bot?](#set-up)
4) [sources and resources](#sources-and-resources)
5) [features (W.I.P.)](#features)

# about
*This is a go repository for interacting with the Zoom API (Chatbots specifically).*

Due to the COVID-19 (novel coronavirus) pandemic, the majority of my classes will be taking place on Zoom, a free app for meetings.

As an avid enthusiast of golang, I wanted to saturate this new platform with my love of programming and combine my two favorite worlds: school and code. Digging into the documentation, I realized binding Go to Zoom was actually *pretty  straightforward*.

# goal
The goal of this project is to add some (if not all) bindings for Zoom's Chatbot API.

Past this project, my end goal is to use this repository to make an integral solver bot on Zoom for my Calculus class (to be the best TA ever!). It will utilize [integral-calculator.com](https://integral-calculator.com), a very popular website that solves integrals with steps. Hopefully the class can learn from it!

# set-up
Coming soon!

# sources-and-resources

1) [Zoom's tutorial on creating a chatbot](https://marketplace.zoom.us/docs/guides/chatbots/build-a-chatbot)
2) [Using Postman to test Zoom's chatbots](https://marketplace.zoom.us/docs/guides/tools-resources/postman/using-postman-to-test-zoom-chatbots)
3) [Zoom's documentation](https://marketplace.zoom.us/docs/guides)

# features

### Embeds
<img align="center" src="https://i.imgur.com/vfjGQ1b.png" height=343px length=303px>
Code for this test embed:

```go
&Message{
	Head: &MessageHead{
		Text: "_GoZoom Test Embed_",
		SubHead: &Text{
			"> \"Who even uses subheaders anyways?\"",
		},
	},
	Body: []*MessageBody{
		&MessageBody{
		    Type:         "section",
			SidebarColor: "#F12BE4",
			Sections: []*MessageBodySection{
				&MessageBodySection{
					Type: "message",
					Text: "*GoZoom*: <img:https://images-na.ssl-images-amazon.com/images/I/51Mt-I6%2BEQL._AC_SX466_.jpg> Edition",
				},
				&MessageBodySection{
					Type: "message",
					Text: "A go repository for interacting with the Zoom API (Chatbots specifically)",
				},
			},
			Footer: "If you didn't know, today is a " + time.Now().Format("Monday 🙄"),
		},
		&MessageBody{
			Type:           "attachments",
			AttResourceURL: "https://github.com/maxthegopher/gozoom",
			AttImgURL:      "https://i.imgur.com/WqW85j1.jpg",
			AttInformation: &AttData{
				Title: &Text{"GoZoom Github!"},
			},
		},
	},
}
```
_So easy_!
