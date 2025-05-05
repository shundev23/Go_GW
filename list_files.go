package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// カレントディレクトリを取得
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("カレントディレクトリの取得に失敗しました: %v", err)
	}
	fmt.Printf("%s 内のファイル一覧：\n", dir)

	// ディレクトリ内のエントリ読み込み
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("ディレクトリの読み込みに失敗しました: %v", err)
	}

	for _, file := range files {
		// ディレクトリはスキップ（オプション）
		if file.IsDir() {
			continue
		}
		fmt.Println(file.Name())
	}

	fmt.Println("\n---特定の拡張子のファイルのみ---")
	// filepath.Globを使って特定の拡張子のファイルのみを取得
	targetPattern := "*.go"
	matches, err := filepath.Glob(filepath.Join(dir, targetPattern))
	if err != nil {
		log.Printf("Globでのファイル検索に失敗しました: %v", err)
	} else {
		for _, match := range matches {
			fmt.Println(filepath.Base(match)) // ファイル名だけを表示する
		}
	}

}
