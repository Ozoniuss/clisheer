package show

import (
	"fmt"

	"github.com/Ozoniuss/casheer/pkg/casheerapi"
	"github.com/Ozoniuss/clisheer/internal/calls"
	"github.com/Ozoniuss/clisheer/internal/color"
	"github.com/Ozoniuss/clisheer/internal/format"
	"github.com/Ozoniuss/clisheer/internal/state"
	"github.com/spf13/cobra"
)

// liteCmd represents the period command
var liteCmd = &cobra.Command{
	Use:   "lite",
	Short: "Show running totals for current period",
	Long: `
Includes expected and actual income, as well as expected and actual money spent.
`,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO: refactor to use helpers.
		year, month, err := state.GetValidPeriod()
		if err != nil {
			color.Printf("Could not get current period: %s\n", err.Error())
		}

		resp, err := calls.MakeGET[casheerapi.GetTotalResponse](fmt.Sprintf("http://localhost:8033/api/totals/?year=%d&month=%d", year, month))
		if err != nil {
			color.Printf(color.Red, "Could not retrieve current month status: %s", err.Error())
			return
		}

		format.DisplayGetTotalResponse(resp)
		return
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// periodCmd.PersistentFlags().String("foo", "", "A help for foo")

}
