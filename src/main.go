package main

import (
	colorconverter "clh/colorConverter"
	"clh/forms"
	"clh/listastree"
	"clh/timetable"
	"log"
)

func main() {
	utype, err := forms.InitialForm()
	if err != nil {
		log.Fatalf("[ERROR] trying to run initial form with error:\n%v", err.Error())
	}

	switch utype {
	case 0:
		listastree.RunTree()
	case 1:
		colorconverter.Init()
	case 2:
		timetable.GetTime()
	default:
		log.Println("No util know like that")
	}
}
