package colorconverter

import (
	"clh/forms"
	"clh/util"
	"fmt"
)

func Init() {
	convType, err := forms.ConvTypeForm()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	input, err := forms.InputForm()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	var output string
	switch convType {
	case 0:
		output, err = RGBtoHex(input)
	case 1:
		output, err = HextoRGB(input)
	case 2:
		output, err = RGBtoRGB1_0(input)
	default:
		fmt.Println("Invalid option")
	}

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(output)
		util.CopyToClipboard(output)
	}
}

func RGBtoHex(rgb string) (string, error) {
	//? input is something like rgb(256, 256, 256)
	var r, g, b int

	fmt.Sscanf(rgb, "rgb(%d, %d, %d)", &r, &g, &b)
	for _, c := range []int{r, g, b} {
		if c < 0 || c > 255 {
			fmt.Println("Invalid RGB value")
			err := fmt.Errorf("invalid RGB value")
			return "", err
		}
	}

	return fmt.Sprintf("#%02x%02x%02x", r, g, b), nil
}

func HextoRGB(hex string) (string, error) {
	//? input is something like #ffffff

	if len(hex) != 7 {
		fmt.Println("Invalid hex value")
		err := fmt.Errorf("invalid hex value")
		return "", err
	}

	var r, g, b int
	fmt.Sscanf(hex, "#%02x%02x%02x", &r, &g, &b)

	return fmt.Sprintf("rgb(%d, %d, %d)", r, g, b), nil
}

func RGBtoRGB1_0(rgb string) (string, error) {
	//? input is something like rgb(256, 256, 256)
	var r, g, b int

	fmt.Sscanf(rgb, "rgb(%d, %d, %d)", &r, &g, &b)
	for _, c := range []int{r, g, b} {
		if c < 0 || c > 255 {
			fmt.Println("Invalid RGB value")
			err := fmt.Errorf("invalid RGB value")
			return "", err
		}
	}

	return fmt.Sprintf("{%f, %f, %f}", float64(r)/255, float64(g)/255, float64(b)/255), nil
}
