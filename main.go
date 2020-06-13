package main

import (
	"extenso/extenso"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numero, _ := strconv.ParseInt(os.Args[1], 10, 0)
	fmt.Print(extenso.Extenso(numero))
}
