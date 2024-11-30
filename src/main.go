package main

import "github.com/charmbracelet/huh"

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().Title("note 1"),
		),
	)

	form.Run()
}
