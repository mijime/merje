package and

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		curr interface{}
		next interface{}
	}

	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			args: args{},
		},
		{
			args: args{
				curr: map[string]interface{}{
					"str": "1",
					"int": 1,
					"arr": []string{"s1", "s2"},
				},
				next: map[string]interface{}{
					"str": "2",
					"flo": 2.0,
					"arr": []string{"s3", "s4"},
				},
			},
			want: map[string]interface{}{
				"str": "2",
				"arr": []interface{}{"s1", "s2", "s3", "s4"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.curr, tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
