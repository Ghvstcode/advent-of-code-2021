package main

import (
	"github.com/Ghvstcode/advent-of-code/day1/i"
	"github.com/Ghvstcode/advent-of-code/day1/ii"
	"github.com/spf13/cobra"
)


var Cmd = &cobra.Command{
	Use:           "Advent",
	SilenceUsage:  true,
	SilenceErrors: true,
	Short:         "Run all the solved advent exercises",
	RunE: func(cmd *cobra.Command, args []string) error{
		//Cringe! Cringe! Cringe!
		err := i.SaveTheElves()
		err = ii.SaveTheThreeElves()
		return err
	},
}


func main() {
	Cmd.Execute()
}
