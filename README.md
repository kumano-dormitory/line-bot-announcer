# line-bot-announcer

全寮向けに周知したい内容を各BLのグループラインに周知してくれるBot

## LINE Botの動作

1. SCラインでメンションされる
2. メンションされた内容をSCライン以外のグループで繰り返す

メンションされた時に一応、ボタンを出して各グループラインで繰り返すかどうかの確認をとる？

## 必要なこと

LINE Botが、自分の所属している全てのグループのgroup IDを取得する必要がある。そこからSCラインを除いた全てのグループにメッセージを送信する。

## 参考文献

- [公式ドキュメント](https://godoc.org/github.com/line/line-bot-sdk-go/linebot)

- [GolangとGoogle App Engineを使ってLINEBotを作ってみる](https://qiita.com/moja0316/items/a726ef746476fe470a66)
- [Goで書いたAPIサーバーをGoogle App Engineへデプロイするまでの最低限の手順](https://qiita.com/croquette0212/items/1e9df0f25f69b97d06e2)
- [LINE Botを使って別グループに代理発言させる方法](https://arukayies.com/gas/line_bot/speak-on-behalf)
