package actions

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func AppendFileName(text string, force bool) error {
	makeChange := force
	input := ""

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	files, err := os.ReadDir(wd)
	if err != nil {
		return err
	}

	if !makeChange {
		for _, file := range files {
			if !file.IsDir() {
				oldName := file.Name()
				ext := filepath.Ext(oldName)
				newName := fmt.Sprintf("%s%s%s", oldName[:len(oldName)-len(ext)], text, ext)
				fmt.Println(oldName + " will change to " + newName)
			}
		}

		fmt.Print("Would you like to make this change? [Y/n]: ")
		fmt.Scanln(&input)

		if input == "" || strings.ToLower(input) == "y" || strings.ToLower(input) == "yes" {
			makeChange = true
		} else {
			fmt.Println("No change was made")
			return nil
		}
	}

	if makeChange {
		for _, file := range files {
			if !file.IsDir() {
				oldName := file.Name()
				ext := filepath.Ext(oldName)
				newName := fmt.Sprintf("%s%s%s", oldName[:len(oldName)-len(ext)], text, ext)

				err := os.Rename(oldName, newName)
				if err != nil {
					return err
				}
				fmt.Println(oldName + " was changed to " + newName)
			}
		}
	}

	return nil
}
