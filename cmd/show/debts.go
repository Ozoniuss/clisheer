package show

import (
	"encoding/json"
	"fmt"

	"github.com/Ozoniuss/casheer/pkg/casheerapi"
	"github.com/Ozoniuss/clisheer/internal/calls"
	"github.com/Ozoniuss/clisheer/internal/color"
	"github.com/Ozoniuss/clisheer/internal/format"
	"github.com/spf13/cobra"
)

var showDebtResponse bool
var showDebtVerbose bool

// liteCmd represents the period command
var debtsCmd = &cobra.Command{
	Use:   "debts",
	Short: "Show current debts",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		resp, err := calls.MakeGET[casheerapi.ListDebtResponse](fmt.Sprintf("http://localhost:8033/api/debts/"))
		if err != nil {
			color.Printf(color.Red, "Could not retrieve debts: %s", err.Error())
			return
		}

		if showDebtResponse {
			out, err := json.MarshalIndent(resp, "", "  ")
			if err != nil {
				color.Printf(color.Red, "Could not print response: %s\n", err.Error())
				return
			}
			color.Println(color.White, string(out))
		} else {
			format.DisplayListDebtResponse(resp, showDebtVerbose)
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// periodCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	debtsCmd.Flags().BoolVarP(&showDebtResponse, "response", "", false, "Display raw json response")
	debtsCmd.Flags().BoolVarP(&showDebtVerbose, "verbose", "v", false, "Display additional debt information")
}
