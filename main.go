package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	numero, _ := strconv.ParseInt(os.Args[1], 10, 0)
	fmt.Print(Extenso(numero))
}
