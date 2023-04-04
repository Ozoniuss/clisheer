/*
Copyright © 2023 Ozoniuss
*/
package add

import (
	"github.com/spf13/cobra"
)

// AddCmd represents the set command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add something",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	AddCmd.AddCommand(debtsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
