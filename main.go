package main

import (
	"os"
	"fmt"
	"strings"
	"time"
)

const URL_MENSAPLAN = "https://www.uni-ulm.de/mensaplan/data/mensaplan.json"
const URL_MENSAPLAN_STATIC = "https://www.uni-ulm.de/mensaplan/data/mensaplan_static.json"

var place string
var day int

var supportedPlaces map[string]string
var weekDays map[string]int

func init() {
	supportedPlaces = map[string]string{
		"mensa":      "Mensa University",
		"bistro":     "Bistro",
		"burgerbar":  "Burgerbar Southside",
		"cafeteriab": "CafeteriaB",
		"west":       "Cafeteria West",
		"westside":   "West Side Diner",
		"hochschule": "Mensa Hochschule",
	}

	weekDays = map[string]int{
		"mon": 0,
		"tue": 1,
		"wed": 2,
		"thu": 3,
		"fri": 4,
	}
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("fix usage pls thx")
		os.Exit(1)
	}

	if !isInMap(supportedPlaces, os.Args[1]) {
		fmt.Println("invalid place")
		os.Exit(1)
	}
	place = os.Args[1]

	if len(os.Args) >= 3 {
		var ok bool
		day, ok = weekDays[strings.ToLower(os.Args[2])[:3]]
		if !ok {
			fmt.Println("weekday not recognized")
		}
	} else {
		// time.Weekday is specified as sunday = 0, we use monday = 0
		day = int(time.Now().Weekday()) - 1
	}

	fmt.Printf("Mensaplan for \"%s\" on %s:\n", supportedPlaces[place], time.Weekday(day+1))

	// TODO print plan

}

func isInMap(m map[string]string, k string) bool {
	if _, ok := m[k]; ok {
		return true
	}
	return false
}
