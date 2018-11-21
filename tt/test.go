package main

import (
	"fmt"
	"github.com/nsdictionary/go_study/hangul"
)

func fruits() {
	fruits := [3]string{"사과", "바나나", "토마토"}
	for _, fruit := range fruits {
		if hangul.HasConsonantSuffix(fruit) {
			fmt.Printf("%s은 맛있다.\n", fruit)
		} else {
			fmt.Printf("%s는 맛있다.\n", fruit)
		}
	}
}

func sliceCopy() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))
	for i := range src {
		dest[i] = src[i]
	}
	fmt.Println(dest)
}

func main() {
	// fruits()
	sliceCopy()
}
