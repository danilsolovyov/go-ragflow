package goragflow

import (
	"bytes"
	"encoding/json"
	"io"
)

func ioReaderFromStruct(v any) (io.Reader, error) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
