package cmd

import (
	"encoding/json"
	"fmt"
)

func formatResult(format string, data interface{}, formatter func() string) (string, error) {
	var output string
	var err error

	switch format {
	case "":
		output = formatter()
	case "json":
		o, e := json.Marshal(data)
		if e != nil {
			err = fmt.Errorf("Failed to Marshal JSON output: %s", e)
		} else {
			output = string(o)
		}
	default:
		err = fmt.Errorf("Unknown output format \"%s\"", format)
	}
	return output, err
}
