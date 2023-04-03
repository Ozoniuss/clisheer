package state

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/rogpeppe/go-internal/lockedfile"
)

// State stores relevant state of the client.
type State struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

// ReadState reads content of the state file. It automatically ensures that the
// file is closed afterwards.
func ReadState() (State, error) {

	var s State
	content, err := lockedfile.Read("state")

	if err != nil {
		// Not having a state file is a perfectly acceptable state.
		if os.IsNotExist(err) {
			return State{}, nil
		} else {
			return State{}, fmt.Errorf("could not read state file: %w", err)
		}
	}

	err = json.Unmarshal(content, &s)
	if err != nil {
		return State{}, fmt.Errorf("invalid state file format: %w", err)
	}

	return s, nil
}

// WriteState writes the state content to a file, creating the file if it
// doesn't exists. It automatically ensures the file is closed afterwards.
func WriteState(state State) error {

	content, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("could not marshal state content to bytes: %w", err)
	}

	wr := bytes.NewReader(content)

	err = lockedfile.Write("state", wr, 644)
	if err != nil {
		return fmt.Errorf("could not write state to file: %w", err)
	}

	return nil
}
