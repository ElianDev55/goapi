/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/ElianDev55/goapi/libs"
	newcommands "github.com/ElianDev55/goapi/new-commands"
	"github.com/ElianDev55/goapi/new-commands/env"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "A new project api",
	Long: `This command creates a new api project on gin with orm grom , this will generate all necessary files and directories.`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		
		namesDirs := env.GenerateNamesDirs()
		nameFiles := env.GenerateNameFiles()

		newListDirs := libs.AddPathTo(args[0], namesDirs)
		newListFiles := libs.AddPathTo(args[0], nameFiles)

		errD1 := libs.CreateDirectories(args)
		errD2 := libs.CreateDirectories(newListDirs)
		errF  := libs.CreateFiles(newListFiles)

		if errD1 != nil || errD2 != nil  || errF != nil {
			fmt.Println("Error creating directories or files")
			return
		}

		errSeed := newcommands.Seed()
		
		if errSeed != nil {
			fmt.Println("Error creating seed")
			return
		}

		fmt.Println("the command new is started", args[0])
		
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}





