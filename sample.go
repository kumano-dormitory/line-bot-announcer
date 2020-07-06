package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/line/line-bot-sdk-go/linebot"
)

const (
	userStatusAvailable    string = "available"
	userStatusNotAvailable string = "not_available"
)

type user struct {
	ID         string `gorm:"primary_key"`
	IDType     string
	Timestamp  time.Time `gorm:"not null;type:datetime"`
	ReplyToken string
	Status     string
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "Holiday88"
	PROTOCOL := "tcp(192.168.10.200:3307)"
	DBNAME := "LineBot"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func sampleMain() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
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
			switch event.Type {
			case linebot.EventTypeFollow:
				// db instance
				db := gormConnect()
				defer db.Close()

				userData := user{ID: event.Source.UserID,
					IDType:     string(event.Source.Type),
					Timestamp:  event.Timestamp,
					ReplyToken: event.ReplyToken,
					Status:     userStatusAvailable,
				}
				err := db.Where(user{ID: event.Source.UserID}).Assign(&userData).FirstOrCreate(&user{})
				log.Println(err)

			case linebot.EventTypeUnfollow:
				log.Println("Unfollow Event: " + event.Source.UserID)
				log.Println(event)

				// db instance
				db := gormConnect()
				defer db.Close()

				userData := user{ID: event.Source.UserID,
					IDType:     string(event.Source.Type),
					Timestamp:  event.Timestamp,
					ReplyToken: event.ReplyToken,
					Status:     userStatusNotAvailable,
				}

				err := db.Where(user{ID: event.Source.UserID}).Assign(&userData).FirstOrCreate(&user{})
				log.Println(err)

			case linebot.EventTypeMessage:
				log.Println(event)
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					switch message.Text {
					case "text":
						resp := linebot.NewTextMessage(message.Text)

						_, err := bot.ReplyMessage(event.ReplyToken, resp).Do()
						if err != nil {
							log.Print(err)
						}
					case "sticker":
						resp := linebot.NewStickerMessage("3", "230")

						_, err = bot.ReplyMessage(event.ReplyToken, resp).Do()
						if err != nil {
							log.Print(err)
						}
					case "location":
						resp := linebot.NewLocationMessage("現在地", "宮城県多賀城市", 38.297807, 141.031)

						_, err = bot.ReplyMessage(event.ReplyToken, resp).Do()
						if err != nil {
							log.Print(err)
						}
					case "image":
						resp := linebot.NewImageMessage("https://farm5.staticflickr.com/4849/45718165635_328355a940_m.jpg", "https://farm5.staticflickr.com/4849/45718165635_328355a940_m.jpg")

						_, err = bot.ReplyMessage(event.ReplyToken, resp).Do()
						if err != nil {
							log.Print(err)
						}
					case "buttontemplate":
						resp := linebot.NewTemplateMessage(
							"this is a buttons template",
							linebot.NewButtonsTemplate(
								"https://farm5.staticflickr.com/4849/45718165635_328355a940_m.jpg",
								"Menu",
								"Please select",
								linebot.NewPostbackAction("Buy", "action=buy&itemid=123", "", "displayText"),
								linebot.NewPostbackAction("Buy", "action=buy&itemid=123", "text", ""),
								linebot.NewURIAction("View detail", "http://example.com/page/123"),
							),
						)

						_, err = bot.ReplyMessage(event.ReplyToken, resp).Do()
						if err != nil {
							log.Print(err)
						}
					case "datetimepicker":
						resp := linebot.NewTemplateMessage(
							"this is a buttons template",
							linebot.NewButtonsTemplate(
								"https://farm5.staticflickr.com/4849/45718165635_328355a940_m.jpg",
								"Menu",
								"Please select a date,  time or datetime",
								linebot.NewDatetimePickerAction("Date", "action=sel&only=date", "date", "2017-09-01", "2017-09-03", ""),
								linebot.NewDatetimePickerAction("Time", "action=sel&only=time", "time", "", "23:59", "00:00"),
								linebot.NewDatetimePickerAction("DateTime", "action=sel", "datetime", "2017-09-01T12:00", "", ""),
							),
						)

						_, err = bot.ReplyMessage(event.ReplyToken, resp).Do()
						if err != nil {
							log.Print(err)
						}
					case "confirm":
						resp := linebot.NewTemplateMessage(
							"this is a confirm template",
							linebot.NewConfirmTemplate(
								"Are you sure?",
								linebot.NewMessageAction("Yes", "yes"),
								linebot.NewMessageAction("No", "no"),
							),
						)

						_, err = bot.ReplyMessage(event.ReplyToken, resp).Do()
						if err != nil {
							log.Print(err)
						}
					case "carousel":
						resp := linebot.NewTemplateMessage(
							"this is a carousel template with imageAspectRatio,  imageSize and imageBackgroundColor",
							linebot.NewCarouselTemplate(
								linebot.NewCarouselColumn(
									"https://farm5.staticflickr.com/4849/45718165635_328355a940_m.jpg",
									"this is menu",
									"description",
									linebot.NewPostbackAction("Buy", "action=buy&itemid=111", "", ""),
									linebot.NewPostbackAction("Add to cart", "action=add&itemid=111", "", ""),
									linebot.NewURIAction("View detail", "http://example.com/page/111"),
								).WithImageOptions("#FFFFFF"),
								linebot.NewCarouselColumn(
									"https://farm5.staticflickr.com/4849/45718165635_328355a940_m.jpg",
									"this is menu",
									"description",
									linebot.NewPostbackAction("Buy", "action=buy&itemid=111", "", ""),
									linebot.NewPostbackAction("Add to cart", "action=add&itemid=111", "", ""),
									linebot.NewURIAction("View detail", "http://example.com/page/111"),
								).WithImageOptions("#FFFFFF"),
							).WithImageOptions("rectangle", "cover"),
						)
						_, err = bot.ReplyMessage(event.ReplyToken, resp).Do()
						if err != nil {
							log.Print(err)
						}
					case "flex":
						resp := linebot.NewFlexMessage(
							"this is a flex message",
							&linebot.BubbleContainer{
								Type: linebot.FlexContainerTypeBubble,
								Body: &linebot.BoxComponent{
									Type:   linebot.FlexComponentTypeBox,
									Layout: linebot.FlexBoxLayoutTypeVertical,
									Contents: []linebot.FlexComponent{
										&linebot.TextComponent{
											Type: linebot.FlexComponentTypeText,
											Text: "hello",
										},
										&linebot.TextComponent{
											Type: linebot.FlexComponentTypeText,
											Text: "world",
										},
									},
								},
							},
						)

						_, err = bot.ReplyMessage(event.ReplyToken, resp).Do()
						if err != nil {
							log.Print(err)
						}
					case "quickresponse":
						resp := linebot.NewTextMessage(
							"Select your favorite food category or send me your location!",
						).WithQuickReplies(
							linebot.NewQuickReplyItems(
								linebot.NewQuickReplyButton("https://example.com/sushi.png", linebot.NewMessageAction("Sushi", "Sushi")),
								linebot.NewQuickReplyButton("https://example.com/tempura.png", linebot.NewMessageAction("Tempura", "Tempura")),
								linebot.NewQuickReplyButton("", linebot.NewLocationAction("Send location")),
							),
						)

						_, err = bot.ReplyMessage(event.ReplyToken, resp).Do()
						if err != nil {
							log.Print(err)
						}
					}
				}
			}
		}
	})

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
