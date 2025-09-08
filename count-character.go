package main

import "fmt"

func CountChar(texte string, char rune) int {
	total := 0
	for _, c := range texte {
		if c == char {
			total++
		}
	}
	return total
}

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}
