package prompter

import (
	"bufio"
	"fmt"
	"os"
)

type FieldType byte

const (
	INT_TYPE FieldType = iota
	FLOAT32_TYPE
	BOOL_TYPE
	STRING_TYPE
)

type TypedField struct {
	Name  string
	Ftype FieldType
}

// DefaultJSONPrompter interactively takes user input for the provided fields,
// and returns the data including the field names in json format.
func DefaultJSONPrompter(header string, fields ...TypedField) []byte {
	fmt.Println(header)
	s := bufio.NewScanner(os.Stdin)
	out := make([]byte, 0, 32)
	out = append(out, '{')

	for _, f := range fields {
		fmt.Print(f.Name + ": ")
		s.Scan()
		appendJsonString(&out, f.Name)
		out = append(out, ':')

		// TODO: type validations
		if f.Ftype == STRING_TYPE {
			appendJsonString(&out, s.Text())
		} else if f.Ftype == FLOAT32_TYPE {
			appendJson(&out, s.Text())
		} else if f.Ftype == INT_TYPE {
			appendJson(&out, s.Text())
		} else if f.Ftype == STRING_TYPE {
			appendJson(&out, s.Text())
		} else {
			panic("prompter encountered invalid type")
		}

		out = append(out, ',')
	}

	out = out[:len(out)-1]

	out = append(out, '}')
	return out
}

// appendString adds a string value to the JSON representation (thus requiring
// additional quotes).
func appendJsonString(out *[]byte, val string) {
	*out = append(*out, '"')
	*out = append(*out, []byte(val)...)
	*out = append(*out, '"')
}

// appendJson adds the byte representation of the value to the JSON
// representation.
func appendJson(out *[]byte, val string) {
	*out = append(*out, []byte(val)...)
}
