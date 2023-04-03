/*
Copyright © 2023 Ozoniuss
*/
package set

import (
	"time"

	"github.com/Ozoniuss/clisheer/internal/color"
	"github.com/Ozoniuss/clisheer/internal/state"
	"github.com/spf13/cobra"
)

var flagPeriodYear int
var flagPeriodMonth int
var flagPeriodNow bool

// periodCmd represents the period command
var periodCmd = &cobra.Command{
	Use:   "period",
	Short: "Set current period",
	Long: `
Current period consists of month and year, and helps filtering out expenses to
avoid cluttering the terminal with irrelevant information.
`,
	Run: func(cmd *cobra.Command, args []string) {

		var month, year int

		if flagPeriodNow {
			month = int(time.Now().Month())
			year = time.Now().Year()

		} else {

			if flagPeriodMonth < 1 || flagPeriodMonth > 12 {
				color.Println(color.Red, "Error: month must be between 1 and 12")
				return
			}

			if flagPeriodYear < 2000 || flagPeriodYear > 2200 {
				color.Println(color.Red, "Error: year should be between 2000 and 2200")
				return
			}

			month = flagPeriodMonth
			year = flagPeriodYear
		}
		err := state.SetPeriod(year, month)
		if err != nil {
			color.Printf(color.Red, "Could not set period: %s", err.Error())
			return
		}
		color.Printf(color.Green, "State modified successfully with year %d and month %d.", year, month)
	},
}

func init() {
	SetCmd.AddCommand(periodCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// periodCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	periodCmd.Flags().IntVarP(&flagPeriodYear, "year", "y", time.Now().Year(), "Sets the year of the current period.")
	periodCmd.Flags().IntVarP(&flagPeriodMonth, "month", "m", int(time.Now().Month()), "Sets the month of the current period.")
	periodCmd.Flags().BoolVarP(&flagPeriodNow, "now", "n", false, "Sets the period based on current time.")

	periodCmd.MarkFlagsRequiredTogether("month", "year")

	// No need to specify month as well.
	periodCmd.MarkFlagsMutuallyExclusive("now", "year")
}
