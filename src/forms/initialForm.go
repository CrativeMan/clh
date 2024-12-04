package forms

import (
	"clh/icons"
	"fmt"

	"github.com/charmbracelet/huh"
)

func InitialForm() (int, error) {
	var utype int
	initialForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().Title("Select the option you want to execute").Options(
				huh.NewOption(fmt.Sprintf("%s Tree", icons.FileTypeIcons[".c"]), 0),
			).Value(&utype),
		),
	)

	err := initialForm.Run()
	return utype, err
}
