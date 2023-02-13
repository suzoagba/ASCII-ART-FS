package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	font, err := os.ReadFile("font/" + os.Args[2] + ".txt") //reading right banner from txt files
	if err != nil {
		fmt.Println(err)
	}
	banner := strings.Split(string(font[1:]), "\n\n")
	printBanner(banner)

}

func printBanner(banner []string) {
	var argument []string = strings.Split(string(os.Args[1]), "\\n")
	for a := range argument {
		var qwerty [8]string
		if argument[0] == "" && (len(argument) == 1 || a == len(argument)-1) {
			break
		} else if argument[a] == "" {
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

			if os.Args[1] != "" { // print line by line
				for i := range qwerty {
					fmt.Println(qwerty[i])
				}
			}
		}
	}
}