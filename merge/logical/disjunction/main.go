package disjunction

import (
	"github.com/mijime/merje/merge"
	"reflect"
)

func init() {
	merge.Factory.Regist("disjunction", operator{})
}

type operator struct{}

func (o operator) Lookup(options interface{}) interface{} {
	op, ok := options.(merge.Options)

	if !ok {
		return nil
	}

	if op.Type == "or" || op.Type == "dis" {
		return o
	}

	return nil
}

func (o operator) Merge(curr, next interface{}) interface{} {
	return o.mergeStruct(curr, next)
}

func (o operator) mergeStruct(curr, next interface{}) interface{} {
	if curr == nil {
		return next
	}

	if next == nil {
		return curr
	}

	cVal := reflect.ValueOf(curr)
	nVal := reflect.ValueOf(next)

	if cVal.Kind() == reflect.Slice && nVal.Kind() == reflect.Slice {
		return o.mergeSlice(cVal, nVal)
	}

	if cVal.Kind() == reflect.Map && nVal.Kind() == reflect.Map {
		cMap, cMapOk := curr.(map[string]interface{})
		nMap, nMapOk := next.(map[string]interface{})

		if cMapOk && nMapOk {
			return o.mergeMap(cMap, nMap)
		}
	}

	return next
}

func (o operator) mergeMap(curr, next map[string]interface{}) map[string]interface{} {
	for k := range next {
		res := o.mergeStruct(curr[k], next[k])

		if res == nil {
			delete(curr, k)
		} else {
			curr[k] = res
		}
	}

	return curr
}

func (o operator) mergeSlice(curr, next reflect.Value) []interface{} {
	res := make([]interface{}, curr.Len()+next.Len())

	for i := 0; i < curr.Len(); i++ {
		res[i] = curr.Index(i).Interface()
	}

	for i := 0; i < next.Len(); i++ {
		res[i+curr.Len()] = next.Index(i).Interface()
	}

	return res
}
