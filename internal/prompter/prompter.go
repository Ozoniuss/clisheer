package prompter

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Ozoniuss/clisheer/internal/color"
)

// scanStringUntilValid is a helper provided to get a string input from the
// terminal. It doesn't do any additional logic to simply retrieving the user
// input; it is only provided for consistency.
func scanStringUntilValid(s *bufio.Scanner, prefix string) string {
	fmt.Print(prefix)
	s.Scan()
	return s.Text()

}

// scanIntUntilValid is a helper provided to get an int input from the terminal.
// It will prompt the user for a value until the value is a valid integer.
func scanIntUntilValid(s *bufio.Scanner, prefix string) int {
	var valint int
	var err error
	for {
		fmt.Print(prefix)
		s.Scan()
		val := s.Text()
		if valint, err = strconv.Atoi(val); err != nil {
			fmt.Println("Invalid number, please try again.")
			continue
		}

		return valint
	}
}

// scanIntUntilValid is a helper provided to get a float input from the terminal.
// It will prompt the user for a value until the value is a valid floating
// point number.
func scanFloatUntilValid(s *bufio.Scanner, prefix string) float64 {
	var valfloat float64
	var err error
	for {
		fmt.Print(prefix)
		s.Scan()
		val := s.Text()
		if valfloat, err = strconv.ParseFloat(val, 64); err != nil {
			fmt.Println("Invalid number, please try again.")
			continue
		}
		return float64(valfloat)
	}
}

func promptYN(msg string) bool {
	fmt.Print(msg, " y/n: ")
	r := bufio.NewReader(os.Stdin)
	var ans byte

	for {
		var err error
		ans, err = r.ReadByte()
		if err != nil {
			color.Printf(color.Red, "Could not read answer: %s\n", err.Error())
			fmt.Print("Try again: ")
			continue
		}
		if (ans != 'y') && (ans != 'n') {
			color.Printf(color.Red, "Invalid answer %c.", ans)
			fmt.Print("Please select either y or n: ")
			continue
		}
		break
	}

	if ans == 'y' {
		return true
	}
	return false
}
