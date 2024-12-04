package listastree

import (
	"clh/icons"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func walk(path string, dashes int, spaces int, level int) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for i, e := range entries {
		if e.Name()[0] == '.' {
			continue
		}

		printSpacing(spaces, level, i, entries, dashes)

		if e.IsDir() {
			printFile(e)
			walk(path+"/"+e.Name(), dashes, spaces, level+1)
			continue
		} else {
			printFile(e)
		}
	}
}

func printSpacing(spaces int, level int, i int, entries []fs.DirEntry, dashes int) {
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
}

func printFile(e fs.DirEntry) {
	if e.IsDir() {
		color.Yellow(fmt.Sprintf("%s %s", icons.FileTypeIcons["directory"], e.Name()))
	} else {
		icon := icons.FileTypeIcons[filepath.Ext(e.Name())]
		if len(icon) == 0 {
			icon = icons.FileTypeIcons["other"]
		}
		color.Cyan(fmt.Sprintf("%s %s", icon, e.Name()))
	}
}

func RunTree() {
	walk(".", 3, 4, 0)
}
