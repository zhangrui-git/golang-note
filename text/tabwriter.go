package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main()  {
	tabw := tabwriter.NewWriter(os.Stdout, 0, 10, 10, '\t', 0)
	fmt.Fprintln(tabw, "a\tb\tc.")
	fmt.Fprintln(tabw, "1\t2\t3.")
	fmt.Fprintln(tabw, "aaa\tbbb\tccc.")
	fmt.Fprintln(tabw, "10\t20\t30.")
	tabw.Flush()

	tabw2 := new(tabwriter.Writer)
	tabw2.Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(tabw2, "a\tb\tc.")
	fmt.Fprintln(tabw2, "1\t2\t3.")
	fmt.Fprintln(tabw2, "aaa\tbbb\tccc.")
	fmt.Fprintln(tabw2, "10\t20\t30.")
	tabw2.Flush()
}