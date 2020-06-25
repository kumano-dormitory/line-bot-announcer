# line-bot-announcer

全寮向けに周知したい内容を各 BL のグループラインに周知してくれる Bot

## LINE Bot の動作

1. SC ラインでメンションされる
2. 「周知しますか？」というクイックレスポンスメッセージを出す
3. メンションされた内容を SC ライン以外のグループで繰り返す

## 必要なこと

- 参加したグループの ID の取得/脱退時に削除
- グループ ID をデータベースに保存
- データベースとのやりとり
- クイックレスポンスメッセージの送信

## 確認すること

- グループに参加したことをイベントとして受け取る方法
- 脱退をイベントとして受け取る方法  
  -> [ここに書いてある](https://developers.line.biz/ja/reference/messaging-api/#join-event)
- ボタンの表示について  
  -> [ここがすごく参考になる](https://blog.kazu634.com/labs/golang/2019-02-23-line-sdk-go/)
- GAE でデータベースを使う方法
- メンションを検知する方法
  -> [参考](https://www.nowsprinting.com/entry/2017/10/01/005607)

## 参考文献

- [公式ドキュメント](https://godoc.org/github.com/line/line-bot-sdk-go/linebot)

- [Golang と Google App Engine を使って LINEBot を作ってみる](https://qiita.com/moja0316/items/a726ef746476fe470a66)
- [Go で書いた API サーバーを Google App Engine へデプロイするまでの最低限の手順](https://qiita.com/croquette0212/items/1e9df0f25f69b97d06e2)
- [LINE Bot を使って別グループに代理発言させる方法](https://arukayies.com/gas/line_bot/speak-on-behalf)
- [Golang で Line API をいじくるよ](https://blog.kazu634.com/labs/golang/2019-02-23-line-sdk-go/)
