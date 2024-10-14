/*
Copyright © 2024 Scott McKeefer scottmckeefer@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	OldText     string
	NewText     string
	Force       bool
	PrependText string
	// AppendText  string
)

// replaceCmd represents the replace command
var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "Replace old text of a file name with new text.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		makeChange := Force
		input := ""

		dir, err := os.ReadDir(".")
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		}

		if !Force {
			fmt.Println("New filenames: ")
			for _, entry := range dir {
				var str strings.Builder
				oldFileName := entry.Name()
				if strings.Contains(oldFileName, OldText) {
					newFileName := strings.ReplaceAll(oldFileName, OldText, NewText)
					str.WriteString(PrependText)
					str.WriteString(newFileName)
					newFileName = str.String()
					fmt.Println(newFileName)
				}
			}

			fmt.Print("Would you like to make this change? [Y/n]: ")
			n, err := fmt.Scanln(&input)

			if n != 0 && err != nil {
				fmt.Println("Error reading input:", err)
				return
			}

			if n == 0 || strings.ToLower(input) == "y" || strings.ToLower(input) == "yes" {
				makeChange = true
			} else {
				fmt.Println("No change was made")
			}
		}

		if makeChange {
			for _, entry := range dir {
				var str strings.Builder
				oldFileName := entry.Name()
				if strings.Contains(oldFileName, OldText) {
					newFileName := strings.ReplaceAll(oldFileName, OldText, NewText)
					str.WriteString(PrependText)
					str.WriteString(newFileName)
					newFileName = str.String()
					oldPath := "./" + oldFileName
					newPath := "./" + newFileName
					fmt.Println(oldFileName + " was changed to " + newFileName)

					// try to rename the file
					err := os.Rename(oldPath, newPath)
					if err != nil {
						fmt.Println("Failed to rename file" + oldFileName)
						return
					}
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(replaceCmd)

	// Here you will define your flags and configuration settings.
	replaceCmd.Flags().StringVarP(&OldText, "OldText", "o", "", "The old text you want to repalce in the file name")
	replaceCmd.Flags().StringVarP(&NewText, "NewText", "n", "", "The new text you want to have in the file name")
	replaceCmd.Flags().BoolVarP(&Force, "Force", "f", false, "Skips the dry run step and changes the file names")
	replaceCmd.Flags().StringVarP(&PrependText, "Prepend", "p", "", "Any text you want to add to the front of the file name")
	// replaceCmd.Flags().StringVarP(&AppendText, "Append", "a", "", "Any text you want to add to the end of the file name")

	if err := replaceCmd.MarkFlagRequired("OldText"); err != nil {
		fmt.Println(err)
	}

	if err := replaceCmd.MarkFlagRequired("NewText"); err != nil {
		fmt.Println(err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// replaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// replaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
