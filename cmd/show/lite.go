package show

import (
	"fmt"

	"github.com/Ozoniuss/clisheer/internal/calls"
	"github.com/Ozoniuss/clisheer/internal/color"
	"github.com/Ozoniuss/clisheer/internal/format"
	"github.com/Ozoniuss/clisheer/internal/state"
	"github.com/spf13/cobra"
)

// liteCmd represents the period command
var liteCmd = &cobra.Command{
	Use:   "lite",
	Short: "Set current period",
	Long: `
Current period consists of month and year, and helps filtering out expenses to
avoid cluttering the terminal with irrelevant information.
`,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: refactor to use helpers.
		year, month, err := state.GetValidPeriod()
		if err != nil {
			color.Printf("Could not get current period: %s\n", err.Error())
		}

		resp, err := calls.MakeGET(fmt.Sprintf("http://localhost:8033/api/totals/?year=%d&month=%d", year, month))
		if err != nil {
			color.Printf(color.Red, "GET request failed: %s", err.Error())
			return
		}

		format.DisplayGetTotalResponse(resp)
		return
	},
}

func init() {
	ShowCmd.AddCommand(liteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// periodCmd.PersistentFlags().String("foo", "", "A help for foo")

}
