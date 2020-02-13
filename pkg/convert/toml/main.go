package toml

import (
	"io"

	"github.com/BurntSushi/toml"
)

func Decode(in io.Reader) (interface{}, error) {
	var data interface{}
	_, err := toml.DecodeReader(in, &data)

	return data, err
}

func Encode(out io.Writer, data interface{}) error {
	return toml.NewEncoder(out).Encode(&data)
}
