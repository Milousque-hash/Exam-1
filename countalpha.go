package main

import "fmt"

func CountAlpha(texte string) int {
	total := 0
	for _, r := range texte {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			total++
		}
	}
	return total
}

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}
