package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Print("input? ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	fmt.Println("input is", text)
	if text == "" {
		text = "https://www.google.com/"
	}
	exec.Command("open", text).Start()
}
