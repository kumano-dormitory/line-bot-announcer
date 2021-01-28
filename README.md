# line-bot-announcer

全寮向けに周知したい内容を各 BL のグループラインに周知してくれる Bot

## LINE Bot の動作

- セットアップ  
  グループに参加したタイミングでグループIDがデータベースに記録される。
  グループ名は取得できないので、データベースにグループのタイプを手書きする。

- 外部からのメッセージ

1. Slack/Discordから通知するメッセージがサーバーに送られてくる
2. 送られてきたメッセージを参加しているライングループに送信する

## 必要なこと

- 参加したグループの ID の取得/脱退時に削除
- グループ ID をデータベースに保存
- データベースとのやりとり
- Push API

## 確認すること

- グループに参加・脱退したことをイベントとして受け取る方法  
  -> [参考](https://developers.line.biz/ja/reference/messaging-api/#join-event)
- ボタンの表示について  
  -> [参考](https://blog.kazu634.com/labs/golang/2019-02-23-line-sdk-go/)
- メンションを検知する方法  
  -> [参考](https://www.nowsprinting.com/entry/2017/10/01/005607)
- GAE でデータベースを使う方法  
  -> [参考](https://qiita.com/teikoku-penguin/items/b6252cd461b3966d53ac)
- 受け取ったメッセージを別のグループラインに送信する方法  
  -> [参考](https://developers.line.biz/ja/reference/messaging-api/#send-push-message)

## 困ったこと

- LINE BOTのWebhook URLに"line"を含む文字列を使ったらダメ出しされた  
  [LINE Developersアカウントでプロバイダーやチャネルが作れないときの対策法](https://qiita.com/hidehiro98/items/4265f42de8e39cb241b6)
- ただ鸚鵡返しをして欲しいだけなのに変なメッセージがついてきた  
  [LINE Botで「メッセージありがとうございます 申し訳ありませんが...」を返信させなくする方法](https://www.virtual-surfer.com/entry/2018/07/22/190000)
- GAEに立てたはずのサーバーにPOSTしたら502が帰ってきた  
  JSONの形式が間違ってただけだった

## 参考文献

- [公式ドキュメント](https://godoc.org/github.com/line/line-bot-sdk-go/linebot)

- [Golang と Google App Engine を使って LINEBot を作ってみる](https://qiita.com/moja0316/items/a726ef746476fe470a66)
- [Go で書いた API サーバーを Google App Engine へデプロイするまでの最低限の手順](https://qiita.com/croquette0212/items/1e9df0f25f69b97d06e2)
- [Golang で Line API をいじくるよ](https://blog.kazu634.com/labs/golang/2019-02-23-line-sdk-go/)
