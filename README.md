# go-notion-api-client

## 環境構築
以下を実行する。
```
go mod tidy
```
また、.envに `GO_NOTION_ACCESS_TOKEN` の値を記入する。
## 実行
以下のいずれかのコマンドを実行する。
```
# xxxには実行したいファイル名を入れる
make go_run-xxx.go
```
or
```
make go_build
```
## 構成
- ディレクトリは分けず、main.goやgo.modと同じ階層に各機能ごとのファイルを作成していく
- ファイル名はsnake caseが一般的

とりあえず[これ](https://github.com/layerXcom/freee-go)を真似していく

## 参考
- https://qiita.com/maiyama18/items/fd382b1a85a7b5071840