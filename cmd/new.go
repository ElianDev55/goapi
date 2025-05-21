/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
		libs.Info("Creating main directory")
		errD2 := libs.CreateDirectories(newListDirs)
		libs.Info("Creating directories")
		errF  := libs.CreateFiles(newListFiles)
		libs.Info("Creating files type .go")

		if errD1 != nil || errD2 != nil  || errF != nil {
			libs.Error("Error creating directories or files, pls check the path")
			return
		}

		errSeed := newcommands.Seed(args[0])
		
		if errSeed != nil {
			libs.Error("Error creating seed file, pls check the path")
			return
		}

		errDbConfig := newcommands.DbConfig(args[0])
		if errDbConfig != nil {
			libs.Error("Error creating db config file, pls check the path")
			return
		}

		errExampleEnv := newcommands.ExampleEnv(args[0])
		if errExampleEnv != nil {
			libs.Error("Error creating example env file, pls check the path")
			return
		}

		errGoMod := newcommands.GoMod(args[0])
		if errGoMod != nil {
			libs.Error("Error creating go mod file, pls check the path")
			return
		}

		errGoSum := newcommands.GoSum(args[0])
		if errGoSum != nil {
			libs.Error("Error creating go sum file, pls check the path")
			return
		}

		libs.Info("Finishing creating project whit the name: " + args[0])
		
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





