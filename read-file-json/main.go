package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// JSONファイルの構造体定義
type Pokemon struct {
	Abilities      []Ability `json:"abilities"`
	BaseExperience int       `json:"base_experience"`
	Forms          []Url     `json:"forms"`
	Height         int       `json:"height"`
	Id             int       `json:"id"`
	IsDefault      bool      `json:"is_default"`
	Name           string    `json:"name"`
	Order          int       `json:"order"`
	Stats          []Stat    `json:"stats"`
	Types          []Type    `json:"types"`
	Weight         int       `json:"weigt"`
}

type Ability struct {
	Ability  Url  `json:"ability"`
	IsHidden bool `json:"is_hidden"`
	Slot     int  `json:"slot"`
}

type Url struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Stat struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	Stat     Url `json:"stat"`
}

type Type struct {
	Slot int `json:"slot"`
	Type Url `json:"type"`
}

func main() {
	// ユーザ入力されたファイル名を取得
	fmt.Print("JSONファイル名は？ ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	filename := scanner.Text()
	fmt.Println("input is", filename)
	if filename == "" {
		filename = "pikachu.json"
	}

	// ファイルを開く
	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	// ファイルを開いて内容を構造体に格納
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var pokemon Pokemon
	json.Unmarshal(raw, &pokemon)

	// 表示
	fmt.Println("-----------------------------")
	fmt.Println("ID: ", pokemon.Id)
	fmt.Println("ポケモン名: ", pokemon.Name)
	fmt.Println("タイプ: ")
	for _, t := range pokemon.Types {
		fmt.Println("  " + t.Type.Name)
	}
	fmt.Println("種族値: ")
	for _, s := range pokemon.Stats {
		fmt.Println("  "+s.Stat.Name+": ", s.BaseStat)
	}
	fmt.Println("-----------------------------")
}
