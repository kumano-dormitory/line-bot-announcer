# line-bot-announcer

全寮向けに周知したい内容を各BLのグループラインに周知してくれるBot

## LINE Botの動作

1. SCラインでメンションされる
2. 「周知しますか？」というクイックレスポンスメッセージを出す
3. メンションされた内容をSCライン以外のグループで繰り返す

## 必要なこと

- 参加したグループのIDの取得/脱退時に削除
- グループIDをデータベースに保存
- データベースとのやりとり
- クイックレスポンスメッセージの送信

## 確認すること

- グループに参加したことをイベントとして受け取る方法
- 脱退をイベントとして受け取る方法  
-> [ここに書いてある](https://developers.line.biz/ja/reference/messaging-api/#join-event)
- ボタンの表示について  
-> [ここがすごく参考になる](https://blog.kazu634.com/labs/golang/2019-02-23-line-sdk-go/)
- GAEでデータベースを使う方法

## 参考文献

- [公式ドキュメント](https://godoc.org/github.com/line/line-bot-sdk-go/linebot)

- [GolangとGoogle App Engineを使ってLINEBotを作ってみる](https://qiita.com/moja0316/items/a726ef746476fe470a66)
- [Goで書いたAPIサーバーをGoogle App Engineへデプロイするまでの最低限の手順](https://qiita.com/croquette0212/items/1e9df0f25f69b97d06e2)
- [LINE Botを使って別グループに代理発言させる方法](https://arukayies.com/gas/line_bot/speak-on-behalf)
- [GolangでLine APIをいじくるよ](https://blog.kazu634.com/labs/golang/2019-02-23-line-sdk-go/)
