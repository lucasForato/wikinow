package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func JsonPrettyPrint(in string) *error {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "  ")
	if err != nil {
		return &err
	}
	fmt.Print(out.String())
	return nil
}
