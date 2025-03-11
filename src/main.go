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

func switchType(utype int) {
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

func switchDebugType(utype int) {
	switch utype {
	case 0:
		consts.TestIconMap()
	default:
		log.Println("No util know like that")
	}
}

func main() {
	debugPtr := flag.Bool("debug", false, "Run debug form")
	typePtr := flag.String("t", "", "Type of form to run")
	helpPtr := flag.Bool("help", false, "Show help")
	flag.Parse()

	if *debugPtr {
		utype := forms.DebugForm()
		switchDebugType(utype)
	} else if *helpPtr {
		forms.HelpForm()
	} else if *typePtr != "" {
		switchType(consts.GetUtilTypeByString(*typePtr))
	} else {
		utype := forms.InitialForm()
		switchType(utype)
	}
}
