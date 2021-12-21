package main

import "fmt"

func main() {
	fmt.Println(Reverse("Hello, World!"))
}

func Reverse(s string) string {
	r := []rune(s)
	length := len(r)

	for i, j := 0, length-1; i < length/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}
