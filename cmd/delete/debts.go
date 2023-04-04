package delete

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

// liteCmd represents the period command
var debtsCmd = &cobra.Command{
	Use:   "debt",
	Short: "Delete a debt",
	Args:  cobra.ExactArgs(1),
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		id := args[0]

		resp, errResp, err := calls.MakeDELETE[casheerapi.DeleteDebtResponse](fmt.Sprintf("http://localhost:8033/api/debts/%s", id))
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
			return
		}

		// Only using multiple ifs to capture unexpected behaviours and show
		// a more clear separation.
		if errResp != nil {
			format.DisplayErrorResponse(*errResp)
			return
		}

		if resp != nil {
			color.Printf(color.Green, "Deleted debt from %s with value %f (id %v).\n",
				resp.Data.Attributes.Person, resp.Data.Attributes.Amount, resp.Data.Id)
			return
		}
		return
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
	// debtsCmd.Flags().BoolVarP(&showDebtVerbose, "verbose", "v", false, "Display additional debt information")
}
