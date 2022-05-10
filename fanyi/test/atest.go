package main

import (
	"flag"
	"fmt"
)

func main() {
	methodPtr := flag.String("method", "default", "method of sample") //return pointer
	valuePtr := flag.Int("value", -1, "value of sample")
	flag.Parse()
	fmt.Println(*methodPtr, *valuePtr)
	flag.PrintDefaults()
}
