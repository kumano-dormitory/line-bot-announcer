package main

import (
	"errors"
)

// コレクションline-groupに登録している全てのグループの情報を取得する
func getAllGroup() ([]Group, error) {
	// 大量にデータを取得しすぎて遅くならないようにデータの取得個数を100個までに制限
	docs, err := client.Collection(collectionName).
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
	_, _, err = client.Collection(collectionName).Add(ctx, Group{
		ID:       groupID,
		IsCenter: false,
	})

	return
}

// データベースからグループIDを削除する
func deleteGropID(groupID string) error {
	// GroupIDからドキュメントを一つだけ取得
	docs, err := client.Collection(collectionName).
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
