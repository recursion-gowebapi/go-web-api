# go-web-api

## Overview

Native Slang API は、日本人英語学習者向けの英語スラング辞書 API です。

スラングの意味・ニュアンス・使用場面・感情カテゴリ・例文などを取得できるほか、検索やフィルタリング、ランダム取得、新規追加などの機能を提供します。

## Demo / Screenshot
![](https://github.com/user-attachments/assets/dd07a678-dbac-4fe2-bc80-b72c8d05c9bf)

## Features

スラング一覧取得
スラング詳細取得
キーワード検索
使用場面（scene）による検索
感情カテゴリ（emotion）による検索
ランダムスラング取得
関連スラング取得
カテゴリ一覧取得
新規スラング登録

## Tech Stack

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?logo=go&logoColor=white)
![HTML](https://img.shields.io/badge/HTML-5-E34F26?logo=html5&logoColor=white)
![CSS](https://img.shields.io/badge/CSS-3-1572B6?logo=css3&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-ES2023-F7DF1E?logo=javascript&logoColor=white)
  
## Setup

**必要環境**
- Go 1.25以上

**手順**
1. リポジトリをクローン
```
git clone https://github.com/recursion-gowebapi/go-web-api.git
cd go-web-api
```
2. サーバーを起動
```
go run .
```
3. ブラウザで開く <br/>
http://localhost:8080

## API Documentation

API 仕様は OpenAPI (Swagger) 形式で管理しています。

docs/swagger.yaml

Swagger Editor に読み込むことで、エンドポイントやリクエスト・レスポンス仕様を確認できます。

## Design / Implementation Notes

- Go 標準ライブラリ（net/http）のみを利用して API を実装
- データストアとして JSON ファイルを採用
- 共通モデルを利用し、レスポンス構造を統一
- handler / model / store を分離し責務を明確化
- 使用場面や感情カテゴリを利用した関連スラング検索機能を実装
- API 動作確認用のシンプルな Web UI を実装
