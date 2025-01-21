package main

import (
	colorconverter "clh/colorConverter"
	"clh/consts"
	"clh/forms"
	"clh/listastree"
	"clh/timetable"
	"flag"
	"log"
)

func main() {
	debugPtr := flag.Bool("debug", false, "Run debug form")
	flag.Parse()

	if *debugPtr {
		utype, err := forms.DebugForm()
		if err != nil {
			log.Fatalf("[ERROR] trying to run debug form with error:\n%v", err.Error())
		}

		switch utype {
		case 0:
			consts.TestIconMap()
		default:
			log.Println("No util know like that")
		}
	} else {
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

}
