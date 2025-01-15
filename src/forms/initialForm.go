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
				huh.NewOption(fmt.Sprintf("%s Tree", consts.FileTypeIcons[".c"]), 0),
				huh.NewOption(fmt.Sprintf("%s Color Converter", consts.FileTypeIcons[".css"]), 1),
				huh.NewOption(fmt.Sprintf("%s School time", consts.FileTypeIcons["book"]), 2),
			).Value(&utype),
		),
	)

	err := initialForm.Run()
	return utype, err
}
