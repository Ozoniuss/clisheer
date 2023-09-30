package format

import (
	"fmt"

	"github.com/Ozoniuss/casheer/pkg/casheerapi"
)

// DisplayErrorResponse displays a jsonapi error response in a more compact
// format.
func DisplayErrorResponse(errResp casheerapi.ErrorResponse) {
	fmt.Printf("%s: %s\n", errResp.Err.Title, errResp.Err.Detail)
}
