package main

import (
	"flag"
	"fmt"
	"github.com/millim/c18d"
	"os"
)

var Usage = func() {
	fmt.Fprintf(os.Stderr, "---> ./c18d num\r\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	flag.Parse()
	num := flag.Args()[0]
	c18d.Run(num)
}
