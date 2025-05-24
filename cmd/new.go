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
		

		var logger = libs.LoggerP{
			Filename: "cmd/new.go",
		}
	
		namesDirs := env.GenerateNamesDirs()
		nameFiles := env.GenerateNameFiles()

		newListDirs := libs.AddPathTo(args[0], namesDirs)
		newListFiles := libs.AddPathTo(args[0], nameFiles)

		errD1 := libs.CreateDirectories(args)
		logger.InfoLogger("CREATING MAIN DIRECTORY")
		errD2 := libs.CreateDirectories(newListDirs)
		logger.InfoLogger("CREATING DIRECTORIES")
		errF  := libs.CreateFiles(newListFiles)
		logger.InfoLogger("CREATING FILES TYPE .GO")

		if errD1 != nil || errD2 != nil  || errF != nil {
			logger.ErrorLogger("ERROR CREATING DIRECTORIES OR FILES, PLS CHECK THE PATH", errD1)
			return
		}

		errSeed := newcommands.Seed(args[0])
		
		if errSeed != nil {
			logger.ErrorLogger("ERROR CREATING SEED FILE, PLS CHECK THE PATH", errSeed)
			return
		}

		errDbConfig := newcommands.DbConfig(args[0])
		if errDbConfig != nil {
			logger.ErrorLogger("ERROR CREATING DB CONFIG FILE, PLS CHECK THE PATH", errDbConfig)
			return
		}

		errExampleEnv := newcommands.ExampleEnv(args[0])
		if errExampleEnv != nil {
			logger.ErrorLogger("ERROR CREATING EXAMPLE ENV FILE, PLS CHECK THE PATH", errExampleEnv)
			return
		}

		errGoMod := newcommands.GoMod(args[0])
		if errGoMod != nil {
			logger.ErrorLogger("ERROR CREATING GO MOD FILE, PLS CHECK THE PATH", errGoMod)
			return
		}

		errGoSum := newcommands.GoSum(args[0])
		if errGoSum != nil {
			logger.ErrorLogger("ERROR CREATING GO SUM FILE, PLS CHECK THE PATH", errGoSum)
			return
		}

		errPkgBd := newcommands.PkgDb(args[0])
		if errPkgBd != nil {
			logger.ErrorLogger("ERROR CREATING GO SUM FILE, PLS CHECK THE PATH", errPkgBd)
			return
		}

		errMainFile := newcommands.Main(args[0])
		if errMainFile != nil {
			logger.ErrorLogger("ERROR CREATING GO SUM FILE, PLS CHECK THE PATH", errPkgBd)
		}

		logger.InfoLogger("FINISHING CREATING PROJECT WHIT THE NAME: " + args[0])
		
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





