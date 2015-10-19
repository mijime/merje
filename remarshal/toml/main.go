package toml

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/mijime/merje/remarshal"
	"path/filepath"
)

type converter struct{}

func init() {
	remarshal.Factory.Regist("toml", converter{})
}

func (c converter) Lookup(options interface{}) interface{} {
	op, ok := options.(remarshal.Options)

	if !ok {
		return nil
	}

	if filepath.Ext(op.FileName) == ".toml" {
		return c
	}

	if op.Format == "toml" {
		return c
	}

	return nil
}

func (c converter) Unmarshal(buf []byte) (data interface{}, err error) {
	_, err = toml.Decode(string(buf), &data)

	if err != nil {
		return nil, err
	}

	return data, err
}

func (c converter) Marshal(data interface{}) (output []byte, err error) {
	buf := new(bytes.Buffer)
	err = toml.NewEncoder(buf).Encode(data)
	if err != nil {
		return nil, err
	}
	output = buf.Bytes()
	return output, err
}
