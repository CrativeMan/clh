package forms

import (
	"github.com/charmbracelet/huh"
)

func ConvTypeForm() (int, error) {
	var convType int
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().Title("Select the option you want to execute").Options(
				huh.NewOption("RGB to Hex", 0),
				huh.NewOption("Hex to RGB", 1),
				huh.NewOption("RGB to 1-0 RGB", 2),
			).Value(&convType),
		),
	)

	err := form.Run()

	return convType, err
}

func InputForm() (string, error) {
	var input string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Enter input").Value(&input),
		),
	)

	err := form.Run()
	return input, err
}
