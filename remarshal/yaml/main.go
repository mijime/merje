package yaml

import (
	"github.com/mijime/merje/remarshal"
	"gopkg.in/yaml.v2"
	"path/filepath"
)

type converter struct{}

func init() {
	remarshal.Factory.Regist("yaml", converter{})
}

func (c converter) Lookup(options interface{}) interface{} {
	op, ok := options.(remarshal.Options)

	if !ok {
		return nil
	}

	if filepath.Ext(op.FileName) == ".yaml" ||
		filepath.Ext(op.FileName) == ".yml" {
		return c
	}

	if op.Format == "yaml" {
		return c
	}

	return nil
}

func (c converter) Unmarshal(buf []byte) (data interface{}, err error) {
	err = yaml.Unmarshal(buf, &data)

	if err != nil {
		return nil, err
	}

	return remarshal.ConvertMapsToStringMaps(data)
}

func (c converter) Marshal(data interface{}) (output []byte, err error) {
	return yaml.Marshal(&data)
}
