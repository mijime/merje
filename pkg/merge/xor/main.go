package xor

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

	cMap, cMapOk := curr.(map[string]interface{})
	nMap, nMapOk := next.(map[string]interface{})

	if cMapOk && nMapOk {
		return mergeMap(cMap, nMap)
	}

	return nil
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
