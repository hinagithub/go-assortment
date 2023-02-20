package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func main(){

	// ユーザの手を取得
	flag.Parse()
	player_hand := flag.Arg(0)

	// g|c|p チェック
	if player_hand == "" || !(player_hand == "g" || player_hand=="c" || player_hand=="p"){
		fmt.Println("set arg 'g' for Goo | 'c' for Choki | 'p' for Paa ")
		return
	} 

	fmt.Printf("you choose '%v' .", player_hand)

	// cpuの手を生成
	cpu_hand := generateRamdomInt()

	switch player_hand {
	case "g": gResult(cpu_hand)
	case "c": cResult(cpu_hand)
	case "p": pResult(cpu_hand)
	}
	return
}

func generateRamdomInt () string {
	rand.Seed(time.Now().Unix())
	result := rand.Intn(3)
	fmt.Println(result)
	switch result{
		case 0: return "g"
		case 1:return "c"
		case 2: return "p"
	default: return "g"
	}
}

func gResult(cpu_hand string) {
	if cpu_hand=="g"{
		fmt.Println(sameResult("Goo"))
	} else if cpu_hand=="c" {
		fmt.Println(winResult("Goo", "Choki"))
	} else {
		fmt.Println(looseResult("Goo","Paa"))
	}
}
func cResult(cpu_hand string) {
	if cpu_hand=="g"{
		fmt.Println(looseResult("Choki","Goo"))
	} else if cpu_hand=="c" {
		fmt.Println(sameResult("Choki"))
	} else {
		fmt.Println(winResult("Choki","Paa"))
	}
}
func pResult(cpu_hand string) {
	if cpu_hand=="g"{
		fmt.Println(winResult("Paa","Goo"))
	} else if cpu_hand=="c" {
		fmt.Println(looseResult("Paa", "Choki"))
	} else {
		fmt.Println(sameResult("Paa"))
	}
}

func sameResult(hand string)string {
	return ("same! hand: " + hand )
}

func looseResult(player_hand string, cpu_hand string)string {
	return "you loose! hand: " + player_hand + ", cpu_hand: " +cpu_hand
}

func winResult(player_hand string, cpu_hand string)string {
	return "you win! hand: " + player_hand + ", cpu_hand: " +cpu_hand
}



 