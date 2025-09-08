package main

import "fmt"

func CountChar(str string, char rune) int {
	d := 0
	for _, c := range str {
		if c == char {
			d++
		}
	}
	return d
}

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}
