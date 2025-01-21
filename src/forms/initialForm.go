package forms

import (
	"clh/consts"
	"fmt"

	"github.com/charmbracelet/huh"
)

func InitialForm() (int, error) {
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
	return utype, err
}

func DebugForm() (int, error) {
	var utype int
	initialForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().Title("Select the option you want to execute").Options(
				huh.NewOption(fmt.Sprintf("%s Test file type icons", consts.IconMap[".c"]), 0),
			).Value(&utype),
		),
	)

	err := initialForm.Run()
	return utype, err
}
