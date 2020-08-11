package main

import (
	"fmt"
	"strconv"
)

func main()  {
	s := "56"
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
}