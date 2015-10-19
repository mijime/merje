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
	return o.merge(curr, next)
}

func (o operator) merge(curr, next interface{}) interface{} {
	if curr == nil {
		return next
	}

	if next == nil {
		return curr
	}

	cHash, cHashOk := curr.(map[string]interface{})

	if !cHashOk {
		return nil
	}

	nHash, nHashOk := next.(map[string]interface{})

	if !nHashOk {
		return nil
	}

	return o.mergeHash(cHash, nHash)
}

func (o operator) mergeHash(curr, next map[string]interface{}) map[string]interface{} {
	for k := range next {
		res := o.merge(curr[k], next[k])

		if res == nil {
			delete(curr, k)
		} else {
			curr[k] = res
		}
	}

	return curr
}
