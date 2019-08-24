package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"
	"github.com/ski7777/gomultiwa/internal/gomultiwa"
)

func main() {
	parser := argparse.NewParser("GoMultiWA", "Awesome tool with a missing description")
	configpath := parser.String("c", "config", &argparse.Options{Required: false, Help: "Path to gomultiwa.json", Default: "gomultiwa.json"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	gmw, err := gomultiwa.NewGoMultiWA(*configpath)
	if err != nil {
		log.Fatal(err)
	}
	gmw.Start()
	<-make(chan int, 1)
}
