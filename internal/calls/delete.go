package calls

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Ozoniuss/casheer/pkg/casheerapi"
)

// MakeDELETE makes a simple DELETE request to the target url, and returns
// either a typed response or an error response.
func MakeDELETE[T casheerapi.DeleteDebtResponse](url string) (*T, *casheerapi.ErrorResponse, error) {

	var respData T
	var respErr casheerapi.ErrorResponse

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("could not initiate DELETE request: %w", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, nil, fmt.Errorf("DELETE request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not read response body: %w", err)
	}

	// All those status codes are only issued in case of errors.
	if resp.StatusCode >= 400 {
		err = json.Unmarshal(body, &respErr)
		if err != nil {
			return nil, nil, fmt.Errorf("Could not unmarshal request body: %w", err)
		}
		return nil, &respErr, nil
	}

	err = json.Unmarshal(body, &respData)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not unmarshal request body: %w", err)
	}

	return &respData, nil, nil
}
