package main

import (
	"fmt"
	"strings"
)

func main() {
	module := "some module"
	author := "amartya"

	fmt.Println(converter(module, author))
}

func converter(module, author string) (s1, s2 string) {
	module = strings.Title(module)
	author = strings.ToUpper(author)

	return module, author
}
