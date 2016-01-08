package disjunction

import (
	"github.com/mijime/merje/merge"
)

func init() {
	merge.Factory.Regist("exclusiveDisjunction", operator{})
}

type operator struct{}

func (o operator) Lookup(options interface{}) interface{} {
	op, ok := options.(merge.Options)

	if !ok {
		return nil
	}

	if op.Type == "xor" || op.Type == "ex-dis" {
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

	cMap, cMapOk := curr.(map[string]interface{})
	nMap, nMapOk := next.(map[string]interface{})

	if cMapOk && nMapOk {
		return o.mergeMap(cMap, nMap)
	}

	return nil
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
