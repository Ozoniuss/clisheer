package calls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Ozoniuss/casheer/pkg/casheerapi"
)

// MakePOST makes a simple POST request to the target url, and returns
// either a typed response or an error response.
func MakePOST[R casheerapi.CreateDebtRequest, T casheerapi.CreateDebtResponse](url string, payload []byte) (*T, *casheerapi.ErrorResponse, error) {

	var req R
	err := json.Unmarshal(payload, &req)
	if err != nil {
		return nil, nil, fmt.Errorf("Invalid payload format: %w", err)
	}

	var respData T
	var respErr casheerapi.ErrorResponse

	reqBody := bytes.NewReader(payload)

	resp, err := http.Post(url, "application/json", reqBody)
	if err != nil {
		return nil, nil, fmt.Errorf("could not initiate POST request: %w", err)
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
