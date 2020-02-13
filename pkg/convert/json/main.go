package json

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/mijime/merje/pkg/convert"
)

func Decode(in io.Reader) (interface{}, error) {
	var data interface{}

	decoder := json.NewDecoder(in)
	decoder.UseNumber()

	err := decoder.Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}

	return convert.NumbersToInt64(data)
}

func Encode(out io.Writer, data interface{}) error {
	return json.NewEncoder(out).Encode(data)
}
