package toml

import (
	"bytes"
	"github.com/BurntSushi/toml"
	"github.com/mijime/merje/remarshal"
	"path/filepath"
)

type converter struct{}

func init() {
	remarshal.Factory.Regist("toml", New())
}

func New() *converter {
	return &converter{}
}

func (this *converter) Lookup(option interface{}) interface{} {
	op, ok := option.(remarshal.Option)

	if !ok {
		return nil
	}

	if filepath.Ext(op.FileName) == ".toml" {
		return this
	}

	if op.Format == "toml" {
		return this
	}

	return nil
}

func (this *converter) Unmarshal(buf []byte) (data interface{}, err error) {
	_, err = toml.Decode(string(buf), &data)

	if err != nil {
		return nil, err
	}

	return data, err
}

func (this *converter) Marshal(data interface{}) (output []byte, err error) {
	buf := new(bytes.Buffer)
	err = toml.NewEncoder(buf).Encode(data)
	if err != nil {
		return nil, err
	}
	output = buf.Bytes()
	return output, err
}
