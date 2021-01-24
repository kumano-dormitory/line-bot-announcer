package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"cloud.google.com/go/firestore"
)

// アクセストークンをデータベースから取得する
func fetchTokens() (tokens []string, err error) {
	r := struct {
		Token string `firestore:"token"`
	}{}

	ctx := context.Background()
	// Firestoreのclientを初期化
	client, err := firestore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	defer client.Close()

	docs, err := client.Collection("line-notify").Limit(10).Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	tokens = make([]string, 0, 10)
	for _, doc := range docs {
		if err := doc.DataTo(&r); err != nil {
			return nil, err
		}
		tokens = append(tokens, r.Token)
	}

	return tokens, nil
}

// メッセージを送信する
func sendMessage(message, from string) error {
	URL := "https://notify-api.line.me/api/notify"

	form := url.Values{
		"message": {message + "文責:" + from},
	}

	req, err := http.NewRequest("POST", URL, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	tokens, err := fetchTokens()
	if err != nil {
		return err
	}

	for _, token := range tokens {
		req.Header.Set("Authorization", "Bearer "+token)
		// notify apiを叩く
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		if res.StatusCode != 200 {
			return fmt.Errorf("notify api status code:%d", res.StatusCode)
		}
	}

	return nil
}
