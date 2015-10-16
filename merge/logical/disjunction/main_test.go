package disjunction

import (
	"log"
	"testing"
)

func TestRun__main(t *testing.T) {
	op := operator{}

	a := map[string]interface{}{
		"a": 1,
		"b": map[string]interface{}{
			"d": 1,
			"e": 1,
		},
		"f": []int{1, 2, 3, 4},
	}

	b := map[string]interface{}{
		"a": 2,
		"c": 1,
		"b": map[string]interface{}{
			"d": 2,
		},
	}

	log.Print(op.Merge(a, b))
}
