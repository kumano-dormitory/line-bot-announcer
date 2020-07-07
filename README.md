# line-bot-announcer

全寮向けに周知したい内容を各 BL のグループラインに周知してくれる Bot

## LINE Bot の動作

1. Slack/Discordから通知するメッセージがサーバーに送られてくる
2. 送られてきたメッセージに"〇〇から周知です"を付け加えて、参加しているライングループに送信する

## 必要なこと

- 参加したグループの ID の取得/脱退時に削除
- グループ ID をデータベースに保存
- データベースとのやりとり
- クイックレスポンスメッセージの送信

## 確認すること

- グループに参加・脱退したことをイベントとして受け取る方法  
  -> [ここに書いてある](https://developers.line.biz/ja/reference/messaging-api/#join-event)
- ボタンの表示について  
  -> [ここがすごく参考になる](https://blog.kazu634.com/labs/golang/2019-02-23-line-sdk-go/)
- メンションを検知する方法  
  -> [参考](https://www.nowsprinting.com/entry/2017/10/01/005607)
- GAE でデータベースを使う方法
- 受け取ったメッセージを別のグループラインに送信する方法  
  -> [ああ、神よ](https://developers.line.biz/ja/reference/messaging-api/#send-push-message)

## 困ったこと

- LINE BOTのWebhook URLに"line"を含む文字列を使ったらダメ出しされた  
  [LINE Developersアカウントでプロバイダーやチャネルが作れないときの対策法](https://qiita.com/hidehiro98/items/4265f42de8e39cb241b6)
- ただ鸚鵡返しをして欲しいだけなのに変なメッセージがついてきた  
  [LINE Botで「メッセージありがとうございます 申し訳ありませんが...」を返信させなくする方法](https://www.virtual-surfer.com/entry/2018/07/22/190000)

## 参考文献

- [公式ドキュメント](https://godoc.org/github.com/line/line-bot-sdk-go/linebot)

- [Golang と Google App Engine を使って LINEBot を作ってみる](https://qiita.com/moja0316/items/a726ef746476fe470a66)
- [Go で書いた API サーバーを Google App Engine へデプロイするまでの最低限の手順](https://qiita.com/croquette0212/items/1e9df0f25f69b97d06e2)
- [LINE Bot を使って別グループに代理発言させる方法](https://arukayies.com/gas/line_bot/speak-on-behalf)
- [Golang で Line API をいじくるよ](https://blog.kazu634.com/labs/golang/2019-02-23-line-sdk-go/)
