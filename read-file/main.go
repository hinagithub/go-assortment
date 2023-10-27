package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// ユーザ入力されたファイル名を取得
	fmt.Print("input? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	fmt.Println("input is", text)
	if text == "" {
		text = "favorites.txt"
	}

	// ファイルを開く
	fp, err := os.Open(text)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// 内容をスライスに格納して表示
	var lines []string
	fileScanner := bufio.NewScanner(fp)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	fmt.Println(lines)
}
