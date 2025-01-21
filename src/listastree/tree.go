package listastree

import (
	"clh/consts"
	"clh/forms"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func walk(maxDepth int, path string, dashes int, spaces int, level int) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	if maxDepth > 0 && level >= maxDepth {
		return
	}

	for i, e := range entries {
		if e.Name()[0] == '.' {
			continue
		}

		printSpacing(spaces, level, i, entries, dashes)

		if e.IsDir() {
			printFile(e)
			walk(maxDepth, path+"/"+e.Name(), dashes, spaces, level+1)
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
		fmt.Printf("%s%s %s%s\n", consts.IconMap["directory"], consts.YELLOW, e.Name(), consts.DEFAULT)
	} else {
		icon := consts.IconMap[filepath.Ext(e.Name())]
		if len(icon) == 0 {
			icon = consts.IconMap["other"]
		}
		fmt.Printf("%s%s %s%s\n", icon, consts.CYAN, e.Name(), consts.DEFAULT)
	}
}

func RunTree() {
	wd, _ := os.Getwd()
	fmt.Printf("%s%s%s\n", consts.CYAN, wd, consts.DEFAULT)
	maxLevel, _ := forms.TreeForm()
	walk(maxLevel, ".", 3, 4, 0)
}
