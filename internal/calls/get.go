package calls

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/Ozoniuss/casheer/pkg/casheerapi"
)

// MakeGET makes a simple GET request to the target url, and returns either a
// typed response or an error response.
func MakeGET[T casheerapi.GetTotalResponse | casheerapi.ListDebtResponse](url string) (*T, *casheerapi.ErrorResponse, error) {

	var respData T
	var respErr casheerapi.ErrorResponse

	resp, err := client.Get(url)
	if err != nil {
		return nil, nil, fmt.Errorf("GET request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not read response body: %w", err)
	}

	err = json.Unmarshal(body, &respData)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not unmarshal request body: %w", err)
	}

	// All those status codes are only issued in case of errors.
	if resp.StatusCode >= 400 {
		err = json.Unmarshal(body, &respErr)
		if err != nil {
			return nil, nil, fmt.Errorf("Could not unmarshal request body: %w", err)
		}
		return nil, &respErr, nil
	}

	return &respData, nil, nil
}
