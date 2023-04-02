package format

import (
	"fmt"
)

// DisplayResponse displays a response to the terminal.
func DisplayResponse(resp any) {
	switch resp.(type) {

	default:
		fmt.Println(resp)
	}
}
