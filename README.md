# go-notion-api-client

## 環境構築
環境変数 `GO_NOTION_ACCESS_TOKEN` に、Notionのアクセストークンを設定する。

## 実行
以下のいずれかのコマンドを実行する。
```
go run main.go client.go
```
or
```
go build && ./go-notion-api-client
```
## 構成
- ディレクトリは分けず、main.goやgo.modと同じ階層に各機能ごとのファイルを作成していく
- ファイル名はsnake caseが一般的


## 参考
- https://qiita.com/maiyama18/items/fd382b1a85a7b5071840