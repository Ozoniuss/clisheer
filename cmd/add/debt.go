package add

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Ozoniuss/casheer/pkg/casheerapi"
	"github.com/Ozoniuss/clisheer/internal/calls"
	"github.com/Ozoniuss/clisheer/internal/color"
	"github.com/Ozoniuss/clisheer/internal/format"
	"github.com/spf13/cobra"
)

var addDebtFile string
var showDebtResponse bool

// liteCmd represents the period command
var debtsCmd = &cobra.Command{
	Use:   "debt",
	Short: "Add a new debt interactively",
	Args:  cobra.NoArgs,
	Long: `
Fill in the prompter details or read the debt from a file.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		var payload []byte

		if addDebtFile != "" {
			fp, err := filepath.Abs(addDebtFile)
			if err != nil {
				color.Printf(color.Red, "Could not find path: %s", err.Error())
				return
			}

			payload, err = os.ReadFile(fp)
			if err != nil {
				color.Printf(color.Red, "Could not read file content: %s", err.Error())
				return
			}
		} else {
			// Use prompter for the user.
		}

		resp, errResp, err := calls.MakePOST[
			casheerapi.CreateDebtRequest,
			casheerapi.CreateDebtResponse]("http://localhost:8033/api/debts/", payload)
		if err != nil {
			color.Printf(color.Red, "Could not add debt: %s", err.Error())
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

		if errResp != nil {
			format.DisplayErrorResponse(*errResp)
			return
		}

		if resp != nil {
			color.Printf(color.Green, "Added debt for %s with value %f (id %v).\n",
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
	debtsCmd.Flags().StringVarP(&addDebtFile, "file", "f", "", "Read the debt from a file (json).")
	// debtsCmd.Flags().BoolVarP(&showDebtVerbose, "verbose", "v", false, "Display additional debt information")
}
