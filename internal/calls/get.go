package calls

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Ozoniuss/casheer/pkg/casheerapi"
)

// MakeGET makes a simple GET request to the target url, and returns a response
// of the provided struct's type, if possible.
func MakeGET[T casheerapi.GetTotalResponse | casheerapi.ListDebtResponse](url string) (T, error) {

	var respData T

	resp, err := http.Get(url)
	if err != nil {
		return respData, fmt.Errorf("GET request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return respData, fmt.Errorf("Could not read response body: %w", err)
	}

	err = json.Unmarshal(body, &respData)
	if err != nil {
		return respData, fmt.Errorf("Could not unmarshal request body: %w", err)
	}

	return respData, nil
}