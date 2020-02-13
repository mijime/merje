package yaml

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/mijime/merje/pkg/convert"
)

func Decode(in io.Reader) (interface{}, error) {
	buf, err := ioutil.ReadAll(in)
	if err != nil {
		return nil, fmt.Errorf("failed to read yaml: %w", err)
	}

	var data interface{}
	err = yaml.Unmarshal(buf, &data)

	if err != nil {
		return nil, fmt.Errorf("failed to decode yaml: %w", err)
	}

	return convert.MapsToStringMaps(data)
}

func Encode(out io.Writer, data interface{}) error {
	buf, err := yaml.Marshal(&data)
	if err != nil {
		return fmt.Errorf("failed to encode yaml: %w", err)
	}

	_, err = io.Copy(out, bytes.NewReader(buf))
	if err != nil {
		return fmt.Errorf("failed to write yaml: %w", err)
	}

	return nil
}
