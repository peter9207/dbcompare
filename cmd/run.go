/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	executor "github.com/peter9207/dbcompare/executor"
	queries "github.com/peter9207/dbcompare/queries"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run <read> <write> <db>",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 3 {
			cmd.Help()
			return
		}

		read, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		write, err := strconv.Atoi(args[1])
		if err != nil {
			panic(err)
		}

		dbURL := args[2]

		runner, err := queries.NewRunner(dbURL)
		if err != nil {
			panic(err)
		}

		err = runner.Setup()
		if err != nil {
			panic(err)
		}

		exec := executor.NewTimedExecutor(5, runner)
		res, _ := exec.Run(int64(read), int64(write))
		fmt.Printf("in 5 seconds Read: %d, Write: %d", res.ReadCount, res.WriteCount)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

}
