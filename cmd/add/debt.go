package add

import (
	"fmt"
	"math"
	"os"
	"path/filepath"

	"github.com/Ozoniuss/clisheer/casheer"
	"github.com/Ozoniuss/clisheer/internal/format"
	"github.com/Ozoniuss/clisheer/internal/prompter"
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

		if addDebtFile != "" {
			var payload []byte
			fp, err := filepath.Abs(addDebtFile)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not find path: %s", err.Error())
				return
			}

			payload, err = os.ReadFile(fp)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Could not read file content: %s", err.Error())
				return
			}
			fmt.Println("feature not implemented yet.")
			fmt.Println(payload)
			return
		}

		debt := prompter.CreateDebtPrompter("please enter debt details below:")

		resp, err := casheer.Client.CreateDebt(debt.Person, debt.Details, int(math.Round(debt.Amount*100)), debt.Currency, -2)
		if err != nil {
			fmt.Printf("could not create debt: %s\n", err.Error())
			return
		}
		if showDebtResponse {
			format.DisplayRawResponse(resp)
			return
		}

		fmt.Printf("Added debt for %s with value %.2f %s (id %v).\n",
			resp.Data.Attributes.Person, float64(resp.Data.Attributes.Amount)*math.Pow10(resp.Data.Attributes.Exponent), resp.Data.Attributes.Currency, resp.Data.Id)
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
