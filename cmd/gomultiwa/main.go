package main

import (
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"
	"github.com/ski7777/gomultiwa/internal/gomultiwa"
)

func main() {
	config := new(gomultiwa.Config)
	parser := argparse.NewParser("GoMultiWA", "Awesome tool with a missing description")
	config.ConfigPath = parser.String("c", "config", &argparse.Options{Required: false, Help: "Path to gomultiwa.json", Default: "gomultiwa.json"})
	config.DebugConfigPath = parser.String("", "debug-config", &argparse.Options{Required: false, Help: "Path to gomultiwadebug.json", Default: "gomultiwadebug.json"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}
	log.Println("Starting up...")
	gmw, err := gomultiwa.NewGoMultiWA(*config)
	if err != nil {
		log.Fatal(err)
	}
	gmw.Start()
	log.Println("Ready")
	<-make(chan int, 1)
}
