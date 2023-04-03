package state

import (
	"fmt"
)

func SetPeriod(year, month int) error {

	// Validation is already performed, but better safe than sorry.
	if month < 1 || month > 12 {
		return fmt.Errorf("month must be between 1 and 12.")
	}

	if year < 2000 || year > 2200 {
		return fmt.Errorf("year must be between 2000 and 2200.")
	}

	// Reading the state in order to not modify potentially other information
	// from the state.
	_state, err := ReadState()
	if err != nil {
		return fmt.Errorf("could not read state: %w", err)
	}
	_state.Month = month
	_state.Year = year

	err = WriteState(_state)
	if err != nil {
		return fmt.Errorf("could not write state: %w", err)
	}
	return nil
}
