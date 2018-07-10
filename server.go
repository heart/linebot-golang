package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
)

type Configuration struct {
	Port              int
	Connection_String string
}

func main() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/webhook", func(w http.ResponseWriter, req *http.Request) {

		//อ่าน Raw Request 
		rawBody, err := ioutil.ReadAll(req.Body)
	    if err != nil {
	        panic(err)
	    }

	    rawbodyString := string(rawBody)
	    fmt.Println(rawbodyString)
		//-----

		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}

		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					/*if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}*/

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(rawbodyString)).Do(); err != nil {
						log.Print(err)
					}

				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}