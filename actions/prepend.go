package actions

import (
	"fmt"
	"os"
	"strings"
)

func PrependFileName(text string, force bool) error {
	makeChange := force
	input := ""

	dir, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return err
	}

	if !makeChange {
		for _, entry := range dir {
			oldFileName := entry.Name()
			newFileName := text + oldFileName
			fmt.Println(oldFileName + " will change to " + newFileName)
		}

		fmt.Print("Would you like to make this change? [Y/n]: ")
		fmt.Scanln(&input)

		if input == "" || strings.ToLower(input) == "y" || strings.ToLower(input) == "yes" {
			makeChange = true
		} else {
			fmt.Println("No change was made")
		}
	}

	if makeChange {
		for _, entry := range dir {
			oldFileName := entry.Name()
			newFileName := text + oldFileName
			oldPath := "./" + oldFileName
			newPath := "./" + newFileName

			// try to rename the file
			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println("Failed to rename file" + oldFileName)
				return err
			}

			fmt.Println(oldFileName + " was changed to " + newFileName)
		}
	}
	return nil
}
