package listastree

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/fatih/color"
)

func walk(path string, dashes int, spaces int, level int) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for i, e := range entries {
		// check if file is hidden
		if e.Name()[0] == '.' {
			continue
		}

		if spaces > 0 {
			for range level {
				fmt.Print("│")
				for range spaces {
					fmt.Print(" ")
				}
			}
		}
		if i == len(entries)-1 && level > 0 {
			fmt.Print("└")
		} else {
			fmt.Print("├")
		}
		for range dashes {
			fmt.Print("─")
		}
		fmt.Print(" ")

		if e.IsDir() {
			printFile(e)
			walk(path+"/"+e.Name(), dashes, spaces, level+1)
			continue
		} else {
			printFile(e)
		}
	}
}

func printFile(e fs.DirEntry) {
	if e.IsDir() {
		color.Yellow(e.Name())
	} else {
		color.Cyan(e.Name())
	}
}

func RunTree() {
	walk(".", 3, 4, 0)
}
