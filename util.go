package main

import (
	"errors"
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

// コレクションline-groupに登録している全てのグループの情報を取得する
func fetchGroups() ([]Group, error) {
	// 大量にデータを取得しすぎて遅くならないようにデータの取得個数を100個までに制限
	docs, err := client.
		Collection(collectionName).
		Limit(100).
		Documents(ctx).
		GetAll()
	if err != nil {
		return nil, err
	}

	groups := make([]Group, len(docs))
	for i, doc := range docs {
		var g Group
		if err := doc.DataTo(&g); err != nil {
			return nil, err
		}

		groups[i] = g
	}

	return groups, nil
}

// データベースにグループIDを登録する
func registerGroupID(groupID string) (err error) {
	_, _, err = client.
		Collection(collectionName).
		Add(ctx, Group{
			ID:   groupID,
			Type: groupOther,
		})

	return
}

// データベースからグループIDを削除する
func deleteGropID(groupID string) error {
	// GroupIDからドキュメントを一つだけ取得
	docs, err := client.
		Collection(collectionName).
		Where("id", "==", groupID).
		Limit(1).
		Documents(ctx).
		GetAll()
	if err != nil {
		return err
	}

	if len(docs) == 0 {
		return errors.New("Group ID is not found")
	}

	_, err = docs[0].Ref.Delete(ctx)
	return err
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

// メッセージを送信する
func sendMessage(message string, secret bool) error {
	reply := linebot.NewTextMessage(message)

	groups, err := fetchGroups()
	if err != nil {
		return err
	}

	for _, group := range groups {
		// SCラインや全寮ライン・ブロックライン以外のグループにはメッセージを送信しない
		if group.Type == groupSC || group.Type == groupOther {
			continue
		}

		// ブロックラインに送るモードの場合は全寮ラインをスキップ
		if secret && group.Type == groupZR {
			continue
		}

		if _, err := bot.PushMessage(group.ID, reply).Do(); err != nil {
			log.Println(err)
		}
	}

	return nil
}
