package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func JsonToString(i interface{}) string {
	j, err := json.Marshal(i)

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	var buf bytes.Buffer
	jerr := json.Indent(&buf, j, "", " ")

	if jerr != nil {
		fmt.Println(jerr.Error())
		return ""
	}

	return buf.String()
}
