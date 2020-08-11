package main

//import "go/scanner"

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main()  {
	str := `
Great minds have purpose, others have wishes.
// I want, i do.
杰出的人有着目标，其他人只有愿望。
`
	var scan scanner.Scanner
	scan.Init(strings.NewReader(str))
	for t := scan.Scan(); t != scanner.EOF; t = scan.Scan() {
		fmt.Println(scan.TokenText())
	}
}