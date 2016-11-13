package main

import (
	"fmt"
	// "log"
	// "runtime"
	// "strings"
	// "time"
	// "bytes"
	// "unicode/utf8"
	"strconv"
)

func main() {
	a := "23.123"
	res, _ := strconv.ParseFloat(a, 32)
	fmt.Println(res + 1)
	res = Float(a)
	fmt.Println(res)
}
