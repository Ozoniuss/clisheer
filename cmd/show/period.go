package show

import (
	"time"

	"github.com/Ozoniuss/clisheer/internal/color"
	"github.com/Ozoniuss/clisheer/internal/state"
	"github.com/spf13/cobra"
)

// periodCmd represents the period command
var periodCmd = &cobra.Command{
	Use:   "period",
	Short: "Set current period",
	Long: `
Current period consists of month and year, and helps filtering out expenses to
avoid cluttering the terminal with irrelevant information.
`,
	Run: func(cmd *cobra.Command, args []string) {

		_state, err := state.ReadState()
		if err != nil {
			color.Printf(color.Red, "Could not open state file: %s\n", err.Error())
			return
		}

		if _state.Month == 0 {
			color.Printf(color.Yellow, "Month not set. Defaulting to %d.\n", int(time.Now().Month()))
		} else {
			color.Printf(color.Green, "Month set to %d.\n", _state.Month)
		}
		if _state.Year == 0 {
			color.Printf(color.Yellow, "Year not set. Defaulting to %d.\n", time.Now().Year())
		} else {
			color.Printf(color.Green, "Year set to %d.\n", _state.Year)
		}

		return
	},
}

func init() {
	ShowCmd.AddCommand(periodCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// periodCmd.PersistentFlags().String("foo", "", "A help for foo")

}
