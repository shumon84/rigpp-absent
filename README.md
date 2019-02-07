# RiG++欠席連絡確認コマンド
欠席連絡チャンネルに流された欠席連絡から、日時と学籍番号を抽出します。

Slackのuser-tokenと欠席連絡チャンネルのchannel IDが必要です。
欠席連絡チャンネルのchannel IDは `CD78CK2Q5` です。
user-tokenの代わりに [ここ](https://api.slack.com/custom-integrations/legacy-tokens) から取得したlegacy-tokenを使用しても同様に動きます。

# ビルド方法
```
$ go build
```

# 実行方法
channel IDとAPIトークンは環境変数として指定する必要があります。
```
$ export CHANNEL_ID="CD78ck2q5"
$ export TOKEN="xoxp-********"
$ ./rigpp-absent
```
一度に指定して実行することもできます。
```
$ CHANNEL_ID="CD78CK2Q5" TOKEN="xoxp-********" ./rigpp-absent
```

# 出力
1行ごとに、日時(YYYY/MM/DD HH:mm:ss)と学籍番号(ハイフン無し)が空白区切りで出力されます。
```
$ ./rigpp-absent
2018/10/26 04:14:22 26001601234
2018/10/26 04:14:12 26001701234
2018/10/26 04:09:54 26001801234
2018/10/26 04:09:28 26001601234
2018/10/26 01:42:05 26001801234
2018/10/26 12:33:35 26001701234
```
