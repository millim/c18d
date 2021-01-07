package main

import (
	"flag"
	"fmt"
	"github.com/millim/c18d/lib"
	"os"
)

var Usage = func() {
	fmt.Fprintf(os.Stderr, "就一个： ./c18d num\r\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	flag.Parse()
	num := flag.Args()[0]
	lib.Run(num)
}
