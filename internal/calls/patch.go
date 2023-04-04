package calls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Ozoniuss/casheer/pkg/casheerapi"
)

// MakePATCH makes a simple PATCH request to the target url, and returns
// either a typed response or an error response.
func MakePATCH[R casheerapi.UpdateDebtRequest, T casheerapi.UpdateDebtResponse](url string, payload []byte) (*T, *casheerapi.ErrorResponse, error) {

	var req R
	err := json.Unmarshal(payload, &req)
	if err != nil {
		return nil, nil, fmt.Errorf("Invalid payload format: %w", err)
	}

	var respData T
	var respErr casheerapi.ErrorResponse

	reqBody := bytes.NewReader(payload)

	request, err := http.NewRequest("PATCH", url, reqBody)
	if err != nil {
		return nil, nil, fmt.Errorf("could not initiate DELETE request: %w", err)
	}
	request.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, fmt.Errorf("could not initiate PATCH request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not read response body: %w", err)
	}

	// All those status codes are only issued in case of errors.
	if resp.StatusCode >= 400 {
		err = json.Unmarshal(respBody, &respErr)
		if err != nil {
			return nil, nil, fmt.Errorf("Could not unmarshal request body: %w", err)
		}
		return nil, &respErr, nil
	}

	err = json.Unmarshal(respBody, &respData)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not unmarshal request body: %w", err)
	}

	return &respData, nil, nil
}
