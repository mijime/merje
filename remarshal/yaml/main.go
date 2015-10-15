package yaml

import (
	"github.com/mijime/merje/remarshal"
	"gopkg.in/yaml.v2"
	"path/filepath"
)

type converter struct{}

func init() {
	remarshal.Factory.Regist("yaml", New())
}

func New() *converter {
	return &converter{}
}

func (this *converter) Lookup(option interface{}) interface{} {
	roption, ok := option.(remarshal.Option)

	if !ok {
		return nil
	}

	if filepath.Ext(roption.FileName) == ".yaml" ||
		filepath.Ext(roption.FileName) == ".yml" {
		return this
	}

	if roption.Format == "yaml" {
		return this
	}

	return nil
}

func (this *converter) Unmarshal(buf []byte) (data interface{}, err error) {
	err = yaml.Unmarshal(buf, &data)

	if err != nil {
		return nil, err
	}

	return remarshal.ConvertMapsToStringMaps(data)
}

func (this *converter) Marshal(data interface{}) (output []byte, err error) {
	return yaml.Marshal(&data)
}
