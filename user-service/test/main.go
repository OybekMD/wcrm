package main

import (
	"fmt"
	"strings"
)

func main() {
	dw := "2023-03-08|W"
	dws := strings.Split(dw, "|")
	fmt.Println(dws[0])

}
