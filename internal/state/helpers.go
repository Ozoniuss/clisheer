package state

import (
	"fmt"
	"time"
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

// GetPeriod returns the actual period stored by the state, first the year and
// then the month.
func GetPeriod() (int, int, error) {
	_state, err := ReadState()
	if err != nil {
		return 0, 0, fmt.Errorf("Could not open state file: %w", err)
	}

	return _state.Year, _state.Month, nil
}

// GetPeriod returns a valid period, either what is stored in the state or
// the current year and month.
func GetValidPeriod() (int, int, error) {

	var year, month int

	_state, err := ReadState()
	if err != nil {
		return 0, 0, fmt.Errorf("Could not open state file: %w", err)
	}
	if _state.Month == 0 {
		month = int(time.Now().Month())
	}
	if _state.Year == 0 {
		year = time.Now().Year()
	}
	return year, month, nil
}
