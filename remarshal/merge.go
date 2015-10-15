package remarshal

func Merge(curr, next interface{}) interface{} {
	return merge(curr, next)
}

func merge(curr, next interface{}) interface{} {
	if curr == nil {
		return next
	}

	if next == nil {
		return curr
	}

	switch curr.(type) {
	case map[string]interface{}:
		switch next.(type) {
		case map[string]interface{}:
			cHash := curr.(map[string]interface{})
			nHash := next.(map[string]interface{})
			return mergeHash(cHash, nHash)
		default:
			return next
		}
	default:
		return next
	}
}

func mergeHash(curr, next map[string]interface{}) map[string]interface{} {
	for k := range next {
		curr[k] = merge(curr[k], next[k])
	}

	return curr
}
