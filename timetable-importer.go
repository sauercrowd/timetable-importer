package main

import (
	"flag"
	"fmt"

	"github.com/sauercrowd/timetable-importer/pkg/parser"
)

func main() {
	filePath := flag.String("file", "times.yml", "yaml file to import")
	flag.Parse()
	tt, err := parser.Parse(*filePath)
	if err != nil {
		fmt.Println("Could not parse file:", err)
		return
	}
	fmt.Println("Successfully parsed", tt.Name)

}
