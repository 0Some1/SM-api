package muxlib

import (
	"bytes"
	"encoding/json"
)

func ConvertToJsonBytes(payload interface{}) ([]byte, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(payload)
	return buffer.Bytes(), err
}
