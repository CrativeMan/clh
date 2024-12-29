package forms

import (
	"strconv"

	"github.com/charmbracelet/huh"
)

func TreeForm() (int, error) {
	var maxLevelS string
	initialForm := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Max Level").Placeholder("0").Value(&maxLevelS),
		),
	)

	err := initialForm.Run()

	maxLevel, _ := strconv.Atoi(maxLevelS)
	return maxLevel, err
}
