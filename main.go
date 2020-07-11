package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/linebot/httphandler"
)

const (
	collectionName = "line-groups"
	onlyBlockGruop = 0
	allGroups      = 1
)

// Group データベース上でのライングループの表現
type Group struct {
	ID       string `firestore:"id"`        // ライングループID
	IsCenter bool   `firestore:"is_center"` //　全寮ラインかどうかのフラグ
}

var (
	ctx    context.Context
	bot    *linebot.Client
	client *firestore.Client
)

func init() {
	ctx = context.Background()
}

func main() {
	// HTTP Handlerの初期化
	handler, err := httphandler.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	bot, err = handler.NewClient()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Firestoreのclientを初期化
	client, err = firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer client.Close()

	// LINEから受け取ったイベントを処理するハンドラを設定する
	handler.HandleEvents(lineEventsHandler)

	// ハンドラを設定
	http.Handle("/callback", handler)
	http.HandleFunc("/messages", reciever)

	// ポート番号を設定
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// HTTPサーバの起動
	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatal(err)
		return
	}
}

// LINEからイベントを受けとって処理する
func lineEventsHandler(events []*linebot.Event, r *http.Request) {
	for _, event := range events {
		groupID := event.Source.GroupID

		switch event.Type {
		case linebot.EventTypeJoin:
			// 入室時に参加したグループのIDを登録する
			if err := registerGroupID(groupID); err != nil {
				log.Fatal(err)
				return
			}

		case linebot.EventTypeLeave:
			// 退出時に脱退したグループIDを削除する
			if err := deleteGropID(groupID); err != nil {
				log.Fatal(err)
				return
			}

		case linebot.EventTypeMessage:
			if message, ok := event.Message.(*linebot.TextMessage); ok && message.Text == "leave" {
				// "leave"というメッセージがきた場合にはグループから退出する
				if err := leaveGroup(event.Source); err != nil {
					log.Fatal(err)
					return
				}
			}
		}
	}
}

// 他所からメッセージを受け取ってラインに流す
func reciever(w http.ResponseWriter, req *http.Request) {
	r := struct {
		Mode    uint8  `json:"mode"`
		Message string `json:"message"`
		From    string `json:"from"`
	}{}

	// POST以外は弾く
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Only Post is allowed")
	}

	bytesBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintln(w, err)
		return
	}

	if err := json.Unmarshal(bytesBody, &r); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusNotAcceptable)
		fmt.Fprintln(w, err)
		return
	}

	reply := linebot.NewTextMessage(r.Message)

	groups, err := getAllGroup()
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, group := range groups {
		// ブロックラインに送るモードの場合はスキップ
		if r.Mode == onlyBlockGruop && group.IsCenter {
			continue
		}

		if _, err := bot.PushMessage(group.ID, reply).Do(); err != nil {
			log.Fatal(err)
			w.WriteHeader(http.StatusNotAcceptable)
			fmt.Fprintln(w, err)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Success")
}

// グループ・トークルームから退出させる
func leaveGroup(eventSource *linebot.EventSource) (err error) {
	switch eventSource.Type {
	case linebot.EventSourceTypeGroup:
		if _, err = bot.LeaveGroup(eventSource.GroupID).Do(); err != nil {
			return
		}

		if err = deleteGropID(eventSource.GroupID); err != nil {
			return
		}

	case linebot.EventSourceTypeRoom:
		_, err = bot.LeaveRoom(eventSource.RoomID).Do()
		return
	}

	return
}
