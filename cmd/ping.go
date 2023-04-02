/*
Copyright © 2023 Ozoniuss
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"

	color "github.com/Ozoniuss/clisheer/internal/color"
	ijson "github.com/Ozoniuss/clisheer/internal/json"

	"github.com/spf13/cobra"
)

// pingCmd represents the ping command
var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Checks if the casheer backend is active.",
	Run: func(cmd *cobra.Command, args []string) {
		color.Println(color.Green + "Pinging casheer server...")
		resp, err := http.Get("http://localhost:8033")
		if err != nil {
			color.Printf(color.Red, "Could not reach server: %s\n", err.Error())
			color.Println(color.Red + "Ping failed. Casheer server unreachable.")
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Could not read response body: %s", err.Error())
		}
		// No need to do anything special here, just print the response in a
		// nice format.

		fmt.Print(color.White)
		ijson.PrintJsonByte(body)
		fmt.Print(color.Reset)
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
