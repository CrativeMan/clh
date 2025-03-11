package forms

import (
	"clh/consts"
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
)

func InitialForm() int {
	var utype int
	initialForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().Title("Select the option you want to execute").Options(
				huh.NewOption(fmt.Sprintf("%s Tree", consts.IconMap[".c"]), 0),
				huh.NewOption(fmt.Sprintf("%s Color Converter", consts.IconMap[".css"]), 1),
				huh.NewOption(fmt.Sprintf("%s School time", consts.IconMap["book"]), 2),
			).Value(&utype),
		),
	)

	err := initialForm.Run()

	if err != nil {
		log.Fatalf("[ERROR] trying to run initial form with error:\n%v", err.Error())
	}

	return utype
}

func DebugForm() int {
	var utype int
	initialForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().Title("Select the option you want to execute").Options(
				huh.NewOption(fmt.Sprintf("%s Test file type icons", consts.IconMap[".c"]), 0),
			).Value(&utype),
		),
	)

	err := initialForm.Run()

	if err != nil {
		log.Fatalf("[ERROR] trying to run debug form with error:\n%v", err.Error())
	}
	return utype
}
