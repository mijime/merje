package remarshal

import (
	"encoding/json"
)

// ConvertMapsToStringMaps is
func ConvertMapsToStringMaps(item interface{}) (res interface{}, err error) {
	return convertMapsToStringMaps(item)
}

// ConvertNumbersToInt64 is
func ConvertNumbersToInt64(item interface{}) (res interface{}, err error) {
	return convertNumbersToInt64(item)
}

func convertMapsToStringMaps(item interface{}) (res interface{}, err error) {
	switch item.(type) {
	case map[interface{}]interface{}:
		res := make(map[string]interface{})
		for k, v := range item.(map[interface{}]interface{}) {
			res[k.(string)], err = convertMapsToStringMaps(v)
			if err != nil {
				return nil, err
			}
		}
		return res, nil
	case []interface{}:
		res := make([]interface{}, len(item.([]interface{})))
		for i, v := range item.([]interface{}) {
			res[i], err = convertMapsToStringMaps(v)
			if err != nil {
				return nil, err
			}
		}
		return res, nil
	default:
		return item, nil
	}
}

func convertNumbersToInt64(item interface{}) (res interface{}, err error) {
	switch item.(type) {
	case map[string]interface{}:
		res := make(map[string]interface{})
		for k, v := range item.(map[string]interface{}) {
			res[k], err = convertNumbersToInt64(v)
			if err != nil {
				return nil, err
			}
		}
		return res, nil
	case []interface{}:
		res := make([]interface{}, len(item.([]interface{})))
		for i, v := range item.([]interface{}) {
			res[i], err = convertNumbersToInt64(v)
			if err != nil {
				return nil, err
			}
		}
		return res, nil
	case json.Number:
		n, err := item.(json.Number).Int64()
		if err != nil {
			f, err := item.(json.Number).Float64()
			if err != nil {
				// Can't convert to Int64.
				return item, nil
			}
			return f, nil
		}
		return n, nil
	default:
		return item, nil
	}
}
