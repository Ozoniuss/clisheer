package format

import (
	"encoding/json"

	"github.com/Ozoniuss/clisheer/internal/color"
)

// DisplayRawResponse prints to the terminal the raw response received from the
// request.
func DisplayRawResponse(resp any) {
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		color.Println(color.Yellow, "Response doesn't have json format. Printing raw response...")
	}
	color.Printf(color.White, "%s\n", out)
	return
}
