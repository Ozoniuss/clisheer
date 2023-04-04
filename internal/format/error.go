package format

import (
	"github.com/Ozoniuss/casheer/pkg/casheerapi"
	"github.com/Ozoniuss/clisheer/internal/color"
)

func DisplayErrorResponse(errResp casheerapi.ErrorResponse) {
	color.Printf(color.Red, "%s: %s\n", errResp.Error.Title, errResp.Error.Detail)
}
