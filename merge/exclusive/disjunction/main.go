package disjunction

import (
	"github.com/mijime/merje/merge"
)

func init() {
	merge.Factory.Regist("exclusiveDisjunction", operator{})
}

type operator struct{}

func (this operator) Lookup(options interface{}) interface{} {
	op, ok := options.(merge.Options)

	if !ok {
		return nil
	}

	if op.Type == "xor" || op.Type == "ex-dis" {
		return this
	}

	return nil
}

func (this operator) Merge(curr, next interface{}) interface{} {
	return this.merge(curr, next)
}

func (this operator) merge(curr, next interface{}) interface{} {
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

	return this.mergeHash(cHash, nHash)
}

func (this operator) mergeHash(curr, next map[string]interface{}) map[string]interface{} {
	for k := range next {
		res := this.merge(curr[k], next[k])

		if res == nil {
			delete(curr, k)
		} else {
			curr[k] = res
		}
	}

	return curr
}
