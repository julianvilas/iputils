package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/julianvilas/iputils"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		usage()
		os.Exit(1)
	}

	ip := args[0]
	networks := args[1:]

	ok, network, err := iputils.ContainsIP(ip, networks...)
	if err != nil {
		panic(err)
	}

	if ok {
		fmt.Println(network)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: iputils [flags] ip network1 [network2 ...]")
	flag.PrintDefaults()
}
