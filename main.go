package main

import (
	day1i "github.com/Ghvstcode/advent-of-code/day1/i"
	day1ii "github.com/Ghvstcode/advent-of-code/day1/ii"
	day2i "github.com/Ghvstcode/advent-of-code/day2/i"
	day2ii "github.com/Ghvstcode/advent-of-code/day2/ii"
	"github.com/spf13/cobra"
)


var Cmd = &cobra.Command{
	Use:           "Advent",
	SilenceUsage:  true,
	SilenceErrors: true,
	Short:         "Run all the solved advent exercises",
	RunE: func(cmd *cobra.Command, args []string) error{
		//Cringe! Cringe! Cringe!
		// Bad code
		err := day1i.SaveTheElves()
		err = day1ii.SaveTheThreeElves()
		err = day2i.Dive()
		err = day2ii.Diveii()
		return err
	},
}


func main() {
	Cmd.Execute()
}
