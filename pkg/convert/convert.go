package convert

import (
	"encoding/json"
)

func MapsToStringMaps(item interface{}) (res interface{}, err error) {
	return mapsToStringMaps(item)
}

func NumbersToInt64(item interface{}) (res interface{}, err error) {
	return numbersToInt64(item)
}

func mapsToStringMaps(item interface{}) (res interface{}, err error) {
	switch itype := item.(type) {
	case map[interface{}]interface{}:
		res := make(map[string]interface{})
		for k, v := range itype {
			res[k.(string)], err = mapsToStringMaps(v)
			if err != nil {
				return nil, err
			}
		}

		return res, nil
	case []interface{}:
		res := make([]interface{}, len(item.([]interface{})))
		for i, v := range item.([]interface{}) {
			res[i], err = mapsToStringMaps(v)
			if err != nil {
				return nil, err
			}
		}

		return res, nil
	default:
		return item, nil
	}
}

func numbersToInt64(item interface{}) (res interface{}, err error) {
	switch itype := item.(type) {
	case map[string]interface{}:
		res := make(map[string]interface{})
		for k, v := range itype {
			res[k], err = numbersToInt64(v)
			if err != nil {
				return nil, err
			}
		}

		return res, nil
	case []interface{}:
		res := make([]interface{}, len(item.([]interface{})))
		for i, v := range item.([]interface{}) {
			res[i], err = numbersToInt64(v)
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
