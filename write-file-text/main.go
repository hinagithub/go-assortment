package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// ファイル名を取得
	fmt.Println("Enter file name")
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filename := scanner.Text()

	fmt.Println("filename is ", filename)
	if filename == "" {
		filename = "favorites.txt"
	}

	// 内容を取得
	var userInput string
	var combinedInput string
	fmt.Printf("Enter text (or press [Enter-Key] to finish): " + "\n")
	for {
		fmt.Print("> ")
		scanner.Scan()
		userInput = scanner.Text()

		if userInput == "" {
			break
		}

		if userInput == "" {
			if combinedInput != "" {
				combinedInput += "\n"
			}
		} else {
			combinedInput += userInput + "\n"
		}
	}
	content := []byte(combinedInput)

	// ファイルに書き込む
	err := ioutil.WriteFile(filename, content, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully saved!")
}
