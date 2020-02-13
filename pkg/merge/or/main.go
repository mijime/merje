package or

import (
	"reflect"
)

func Merge(curr, next interface{}) interface{} {
	return mergeStruct(curr, next)
}

func mergeStruct(curr, next interface{}) interface{} {
	if curr == nil {
		return next
	}

	if next == nil {
		return curr
	}

	cVal := reflect.ValueOf(curr)
	nVal := reflect.ValueOf(next)

	if cVal.Kind() == reflect.Slice && nVal.Kind() == reflect.Slice {
		return mergeSlice(cVal, nVal)
	}

	if cVal.Kind() == reflect.Map && nVal.Kind() == reflect.Map {
		cMap, cMapOk := curr.(map[string]interface{})
		nMap, nMapOk := next.(map[string]interface{})

		if cMapOk && nMapOk {
			return mergeMap(cMap, nMap)
		}
	}

	return next
}

func mergeMap(curr, next map[string]interface{}) map[string]interface{} {
	for k := range next {
		res := mergeStruct(curr[k], next[k])

		if res == nil {
			delete(curr, k)
		} else {
			curr[k] = res
		}
	}

	return curr
}

func mergeSlice(curr, next reflect.Value) []interface{} {
	res := make([]interface{}, curr.Len()+next.Len())

	for i := 0; i < curr.Len(); i++ {
		res[i] = curr.Index(i).Interface()
	}

	for i := 0; i < next.Len(); i++ {
		res[i+curr.Len()] = next.Index(i).Interface()
	}

	return res
}
