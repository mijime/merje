package json

import (
	"bytes"
	"encoding/json"
	"github.com/mijime/merje/remarshal"
	"path/filepath"
)

type converter struct{}

func init() {
	remarshal.Factory.Regist("json", converter{})
}

func (c converter) Lookup(options interface{}) interface{} {
	op, ok := options.(remarshal.Options)

	if !ok {
		return nil
	}

	if filepath.Ext(op.FileName) == ".json" {
		return c
	}

	if op.Format == "json" {
		return c
	}

	return nil
}

func (c converter) Unmarshal(buf []byte) (data interface{}, err error) {
	decoder := json.NewDecoder(bytes.NewReader(buf))
	decoder.UseNumber()
	err = decoder.Decode(&data)

	if err != nil {
		return nil, err
	}

	return remarshal.ConvertNumbersToInt64(data)
}

func (c converter) Marshal(data interface{}) (output []byte, err error) {
	output, err = json.Marshal(&data)

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	err = json.Indent(buf, output, "", "  ")

	if err != nil {
		return nil, err
	}

	output = buf.Bytes()

	return output, err
}
