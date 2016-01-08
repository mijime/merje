package disjunction

import (
	"reflect"
	"testing"
)

func TestHashMergeHash(t *testing.T) {
	op := operator{}

	v1 := map[string]interface{}{
		"kc_number": 1,
		"kc_array": []interface{}{
			1,
			"v1",
			map[string]interface{}{"v1": 1},
		},
		"k1_number": 1,
		"k1_array": []interface{}{
			1,
			"v1",
			map[string]interface{}{"v1": 1},
		},
		"kc_hash": map[string]interface{}{
			"kc_number": 1,
			"kc_array": []interface{}{
				1,
				"v1",
				map[string]interface{}{"v1": 1},
			},
			"k1_number": 1,
			"k1_array": []interface{}{
				1,
				"v1",
				map[string]interface{}{"v1": 1},
			},
		},
	}

	v2 := map[string]interface{}{
		"kc_number": 2,
		"kc_array": []interface{}{
			2,
			"v2",
			map[string]interface{}{"v2": 2},
		},
		"k2_number": 2,
		"k2_array": []interface{}{
			2,
			"v2",
			map[string]interface{}{"v2": 2},
		},
		"kc_hash": map[string]interface{}{
			"kc_number": 2,
			"kc_array": []interface{}{
				2,
				"v2",
				map[string]interface{}{"v2": 2},
			},
			"k2_number": 2,
			"k2_array": []interface{}{
				2,
				"v2",
				map[string]interface{}{"v2": 2},
			},
		},
	}

	res := op.Merge(v1, v2)

	ans := map[string]interface{}{
		"kc_number": 2,
		"kc_array": []interface{}{
			1,
			"v1",
			map[string]interface{}{"v1": 1},
			2,
			"v2",
			map[string]interface{}{"v2": 2},
		},
		"k1_number": 1,
		"k1_array": []interface{}{
			1,
			"v1",
			map[string]interface{}{"v1": 1},
		},
		"k2_number": 2,
		"k2_array": []interface{}{
			2,
			"v2",
			map[string]interface{}{"v2": 2},
		},
		"kc_hash": map[string]interface{}{
			"kc_number": 2,
			"kc_array": []interface{}{
				1,
				"v1",
				map[string]interface{}{"v1": 1},
				2,
				"v2",
				map[string]interface{}{"v2": 2},
			},
			"k1_number": 1,
			"k1_array": []interface{}{
				1,
				"v1",
				map[string]interface{}{"v1": 1},
			},
			"k2_number": 2,
			"k2_array": []interface{}{
				2,
				"v2",
				map[string]interface{}{"v2": 2},
			},
		},
	}

	if !reflect.DeepEqual(res, ans) {
		t.Error("failed", "res <", res, "> ans <", ans, ">")
	}
}

func TestArrayMergeArray(t *testing.T) {
	op := operator{}

	v1 := []interface{}{
		1,
		"v1",
		map[string]interface{}{"v1": 1},
	}

	v2 := []interface{}{
		2,
		"v2",
		map[string]interface{}{"v2": 2},
	}

	res := op.Merge(v1, v2)

	ans := []interface{}{
		1,
		"v1",
		map[string]interface{}{"v1": 1},
		2,
		"v2",
		map[string]interface{}{"v2": 2},
	}

	if !reflect.DeepEqual(res, ans) {
		t.Error("failed", "res <", res, "> ans <", ans, ">")
	}
}

func TestHashMergeArray(t *testing.T) {
	op := operator{}

	v1 := map[string]interface{}{
		"kc_number": 1,
		"kc_array": []interface{}{
			1,
			"v1",
			map[string]interface{}{"v1": 1},
		},
		"k1_number": 1,
		"k1_array": []interface{}{
			1,
			"v1",
			map[string]interface{}{"v1": 1},
		},
		"kc_hash": map[string]interface{}{
			"kc_number": 1,
			"kc_array": []interface{}{
				1,
				"v1",
				map[string]interface{}{"v1": 1},
			},
			"k1_number": 1,
			"k1_array": []interface{}{
				1,
				"v1",
				map[string]interface{}{"v1": 1},
			},
		},
	}

	v2 := []interface{}{
		2,
		"v2",
		map[string]interface{}{"v2": 2},
	}

	res := op.Merge(v1, v2)

	if !reflect.DeepEqual(res, v2) {
		t.Error("failed", "res <", res, "> ans <", v2, ">")
	}
}
