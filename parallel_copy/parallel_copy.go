package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func copyFile(srcPath string, destDir string, wg *sync.WaitGroup) {
	defer wg.Done() //このGoroutineが終了したことを通知

	// ファイル名を取得
	fileName := filepath.Base(srcPath)
	destPath := filepath.Join(destDir, fileName)

	// コピー元ファイルを開く
	srcFile, err := os.Open(srcPath)
	if err != nil {
		log.Printf("ファイルを開けませんでした %s: %v\n", srcPath, err)
	}
	defer srcFile.Close()

	// コピー先ファイルを作成（存在する場合は上書き）
	destFile, err := os.Create(destPath)
	if err != nil {
		log.Printf("ファイルを作成できませんでした %s: %v\n", destPath, err)
	}
	defer destFile.Close()

	// ファイル内容をコピーする
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		log.Printf("ファイルをコピーできませんでした %s -> %s: %v\n", srcPath, destPath, err)
		return
	}

	fmt.Printf("コピー完了: %s -> %s\n", srcPath, destPath)
}

func main() {
	sourceDir, _ := os.Getwd() //エラーハンドリングは省略
	destDir := filepath.Join(sourceDir, "backup")

	// backup ディレクトリを作成する（存在しない場合）
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err := os.Mkdir(destDir, 0755)
		if err != nil {
			log.Fatalf("バックアップディレクトリの作成に失敗しました: %v", err)
		}
		fmt.Printf("ディレクトリを作成しました: %s\n", destDir)
	}
	// コピー対象のファイル（.txt）を探す
	pattern := filepath.Join(sourceDir, "*.txt")
	filesToCopy, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatalf("ファイルの探索に失敗しました: %v", err)
	}

	if len(filesToCopy) == 0 {
		fmt.Println(".txtファイルが見つかりませんでした。")
		return
	}

	fmt.Printf("%d 個の .txtファイルを %s にコピーしています...\n", len(filesToCopy), destDir)

	// WaitGroupを使ってGoRoutineの完了を待つ
	var wg sync.WaitGroup

	for _, srcPath := range filesToCopy {
		wg.Add(1)                          // Goroutineを起動する前にカウンタを増やす増やす
		go copyFile(srcPath, destDir, &wg) // Goroutineとしてコピー処理を開始
	}

	// 全てのGoroutineが完了するまで待つ
	wg.Wait()

	fmt.Println("全てのコピーが完了しました")
}
