package format

import (
	"encoding/json"
	"fmt"
)

// DisplayRawResponse prints to the terminal the raw response received from the
// request.
func DisplayRawResponse(resp any) {
	out, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		fmt.Println("Response doesn't have json format. Printing raw response...")
		fmt.Println(resp)
		return
	}
	fmt.Printf("%s\n", out)
}
