# Go_GW: Go言語 学習ワークショップ

## 概要

このリポジトリは、Go言語の学習過程で作成したサンプルコードをまとめたものです。特に、**簡単なバッチ処理の開発**を通して、Goの**並列処理機能 (GoroutineとChannel)** および**ネットワークライブラリ (`net/http`)** の使い方を学ぶことを目的としています。

macOS環境での動作確認を基本としています。

## このプロジェクトで学べること

* Go言語の基本的な構文と標準ライブラリの使い方
    * ファイルシステム操作 (`os`, `path/filepath`)
    * データ入出力 (`io`)
    * HTTP通信 (`net/http`)
    * JSON処理 (`encoding/json`)
    * 同期処理 (`sync`)
* 簡単なバッチ処理の実装方法
* Goroutine (`go`) を使った処理の並列化
* `sync.WaitGroup` を使ったGoroutineの完了待ち
* Channel (`chan`) を使ったGoroutine間の安全なデータ通信
* `defer` を使ったリソース管理（ファイルクローズ、レスポンスボディクローズなど）
* 基本的なエラーハンドリング (`if err != nil`)

## 含まれるサンプルコード

1.  `list_files.go`
    * **内容:** カレントディレクトリ内のファイル一覧、または特定の拡張子を持つファイル一覧を表示します。
    * **学習テーマ:** 基本的なファイル操作、`os`, `path/filepath` パッケージ。
2.  `parallel_copy.go`
    * **内容:** カレントディレクトリ内の `.txt` ファイルを `backup` ディレクトリに**並列**でコピーします。(`backup` ディレクトリは自動生成されます)
    * **学習テーマ:** Goroutine, `sync.WaitGroup`, ファイルI/O, `defer`。
3.  `Workspace_todo.go`
    * **内容:** [JSONPlaceholder](https://jsonplaceholder.typicode.com/) APIから指定したIDのTODOデータを1件取得し、表示します。
    * **学習テーマ:** `net/http` による基本的なAPIリクエスト、JSONデコード (`encoding/json`)、`defer resp.Body.Close()` の重要性。
4.  `parallel_fetch.go`
    * **内容:** 複数のTODO IDを指定し、それぞれのTODOデータをAPIから**並列**で取得します。結果はChannelを使って集約し、表示します。
    * **学習テーマ:** GoroutineとChannelを組み合わせた並列ネットワーク処理、`range` によるChannelからの受信、`close()`。

## 実行方法

### 前提条件

* Go言語がインストールされていること。
    * インストールガイド: [https://go.dev/doc/install](https://go.dev/doc/install)
    * macOS (Homebrew): `brew install go`

### 各サンプルの実行

ターミナルで各 `.go` ファイルのあるディレクトリに移動し、`go run` コマンドで実行します。

```bash
# ファイル一覧表示
go run list_files.go

# 並列ファイルコピー
# (事前にカレントディレクトリに test1.txt, test2.txt などのファイルを作成してください)
go run parallel_copy.go

# 単一TODO取得 (ID: 1)
go run fetch_todo.go

# 複数TODOを並列取得 (ID: 1, 5, 10, 15, 20)
go run parallel_fetch.go