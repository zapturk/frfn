/*
Copyright © 2024 Scott McKeefer scottmckeefer@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Extension string
)

// prependCmd represents the prepend command
var prependCmd = &cobra.Command{
	Use:   "prepend",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(prependCmd)

	// Here you will define your flags and configuration settings.
	prependCmd.Flags().StringVarP(&NewText, "NewText", "n", "", "The new text you want to have in the file name")
	prependCmd.Flags().StringVarP(&Extension, "Extension", "e", "", "The file extension that you want to change")

	if err := prependCmd.MarkFlagRequired("NewText"); err != nil {
		fmt.Println(err)
	}

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prependCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prependCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
