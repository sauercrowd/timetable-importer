package parser

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/sauercrowd/timetable/pkg/festival"
	"gopkg.in/yaml.v2"
)

func Parse(filepath string) (*festival.TimeTable, error) {
	f, err := os.Open(filepath)
	if err != nil {
		log.Println("Could not open file: ", err)
		return nil, err
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("Could not read file: ", err)
		return nil, err
	}
	var ttable festival.TimeTable
	if err := yaml.Unmarshal(content, &ttable); err != nil {
		return nil, err
	}
	return &ttable, err
}
