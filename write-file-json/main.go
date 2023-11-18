package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

type Favorite struct {
	ID        int
	Title     string
	CreatedAt time.Time
}

func main() {
	// ファイル名を取得
	fmt.Println("Enter file name")
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filename := scanner.Text()

	fmt.Println("filename is ", filename)
	if filename == "" {
		filename = "favorites.json"
	}

	// タスクリストを作成
	favorites := []Favorite{}
	var userInput string
	fmt.Printf("Enter task name (or press [Enter-Key] to finish):" + "\n")
	for {
		fmt.Print("> ")
		scanner.Scan()
		userInput = scanner.Text()

		if userInput == "" {
			break
		}

		task := Favorite{
			ID:        len(favorites) + 1,
			Title:     userInput,
			CreatedAt: time.Now(),
		}
		favorites = append(favorites, task)
	}

	// リストをJSONにエンコード
	jsonData, err := json.MarshalIndent(favorites, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	// ファイルに書き込む
	err = ioutil.WriteFile(filename, jsonData, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully saved!")

}
