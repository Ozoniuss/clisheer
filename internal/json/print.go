package json

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func PrintJsonByte(data []byte) error {
	var out bytes.Buffer
	if err := json.Indent(&out, data, "", "  "); err != nil {
		return fmt.Errorf("could not format json: %w", err)
	}
	fmt.Println(out.String())
	return nil
}
