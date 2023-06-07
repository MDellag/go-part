package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func showOptions() {
	fmt.Println(" 1 - list users ")
	fmt.Println(" 2 - add user ")
	fmt.Println(" 3 - modify user ")
	fmt.Println(" 4 - remove user ")
	fmt.Println(" 0 - show options ")
	fmt.Println(" 9 - exit ")
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("ABM Shell")
	fmt.Println("---------------------")

	showOptions()
	for {

		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)
		if strings.Compare("0", text) == 0 {
			showOptions()
		}
		if strings.Compare("9", text) == 0 {
			break
		}

	}
}
