package main

import (
	"fmt"
	"os"
	"text/template"
)

func main()  {
	type Student struct {
		Name string
		Age uint8
		Grade uint8
		Class uint8
	}

	temp := "name:{{ .Name }} age:{{ .Age }} grade:{{ .Grade }} class:{{ .Class }}"
	student := Student{"Tom cat", 9, 3, 2}
	t, err := template.New("studentInfo").Parse(temp)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = t.Execute(os.Stdout, student)
	if err != nil {
		fmt.Println(err)
		return
	}
}