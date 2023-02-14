package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	switch len(os.Args) {
	case 2:
		//fmt.Println("Usage: go run . [STRING] [BANNER]")
		//return
	//}
		font, _ := os.ReadFile("font/standard.txt")
		banner := strings.Split(string(font[1:]), "\n\n") //reading right banner from txt files
		printBanner(os.Args[1], banner)
	case 3:
		font, err := os.ReadFile("font/" + os.Args[2] + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	banner := strings.Split(string(font[1:]), "\n\n")
	printBanner(os.Args[1], banner)
default:
	fmt.Println("Usage: go run. [STRING] [BANNER]")
	fmt.Println("Ex: go run . something standard")
	}
}

func printBanner(text string, banner []string) {
	var argument []string = strings.Split(text, "\\n")
	if len(argument) == 1 && argument[0] == "" {
		fmt.Println("Error: Empty argument")
		return
	}
	for a := range argument {
		var qwerty [8]string
		if argument[a] == "" {
			fmt.Println()
		} else {
			for _, letter := range argument[a] {
				for i := range banner {
					if letter == rune(i+32) {
						line := strings.Split(banner[i], "\n")
						for l := 0; l < 8; l++ { // write every line sepparately
							qwerty[l] += line[l]
						}
					}
				}
			}

			if text != "" { // print line by line
				for i := range qwerty {
					fmt.Println(qwerty[i])
				}
			}
		}
	}
}
