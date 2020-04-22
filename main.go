package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	xj "github.com/basgys/goxml2json"
)

func onExit(err error) {
	log.Output(2, err.Error())
	os.Exit(1)
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	flag.Parse()
	if flag.NFlag() > 1 {
		flag.Usage()
		os.Exit(0)
	}

	file, err := os.Open(flag.Args()[0])
	if err != nil {
		onExit(err)
	}
	defer file.Close()

	json, err := xj.Convert(file)
	if err != nil {
		onExit(err)
	}

	fmt.Println(json.String())
}
