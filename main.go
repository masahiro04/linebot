package main

import (
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	bot, err := linebot.New(
		os.Getenv("LINE_SECRET_KEY"),
		os.Getenv("LINE_ACCESS_TOKEN"),
	)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	})

	r.POST("/callback", func(c *gin.Context) {
		events, err := bot.ParseRequest(c.Request)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				log.Print(err)
			}
			return
		}

		//// "可愛い" 単語を含む場合、返信される
		//var replyText string
		//replyText = "可愛い"
		//
		//// チャットの回答
		//var response string
		//response = "ありがとう！！"
		//
		//// "おはよう" 単語を含む場合、返信される
		//var replySticker string
		//replySticker = "おはよう"
		//
		//// スタンプで回答が来る
		//responseSticker := linebot.NewStickerMessage("11537", "52002757")
		//
		//// "猫" 単語を含む場合、返信される
		//var replyImage string
		//replyImage = "猫"
		//
		//// 猫の画像が表示される
		//responseImage := linebot.NewImageMessage(
		//	"https://i.gyazo.com/2db8f85c496dd8f21a91eccc62ceee05.jpg",
		//	"https://i.gyazo.com/2db8f85c496dd8f21a91eccc62ceee05.jpg",
		//)
		//
		//// "ディズニー" 単語を含む場合、返信される
		//var replyLocation string
		//replyLocation = "ディズニー"
		//
		//// ディズニーが地図表示される
		//responseLocation := linebot.NewLocationMessage("東京ディズニーランド", "千葉県浦安市舞浜", 35.632896, 139.880394)

		for _, event := range events {
			// イベントがメッセージの受信だった場合
			if event.Type == linebot.EventTypeMessage {
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

	err = r.Run()
	if err != nil {
		return
	}
}
