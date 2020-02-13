package json

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	type args struct {
		in io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			args: args{in: strings.NewReader(`{"str":"s"}`)},
			want: map[string]interface{}{"str": "s"},
		},
		{
			args: args{in: strings.NewReader(`{"int":1}`)},
			want: map[string]interface{}{"int": int64(1)},
		},
		{
			args: args{in: strings.NewReader(`{"flo":2.1}`)},
			want: map[string]interface{}{"flo": float64(2.1)},
		},
		{
			args: args{in: strings.NewReader(`{"bol":true}`)},
			want: map[string]interface{}{"bol": true},
		},
		{
			args: args{in: strings.NewReader(`{"arr":["s1",2]}`)},
			want: map[string]interface{}{"arr": []interface{}{"s1", int64(2)}},
		},
		{
			args:    args{in: strings.NewReader(`{"no-json"}`)},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decode(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	type args struct {
		data interface{}
	}

	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr bool
	}{
		{
			args:    args{data: map[string]interface{}{"str": "s"}},
			wantOut: `{"str":"s"}` + "\n",
		},
		{
			args:    args{data: map[string]interface{}{"int": int64(1)}},
			wantOut: `{"int":1}` + "\n",
		},
		{
			args:    args{data: map[string]interface{}{"flo": float64(2.1)}},
			wantOut: `{"flo":2.1}` + "\n",
		},
		{
			args:    args{data: map[string]interface{}{"bol": true}},
			wantOut: `{"bol":true}` + "\n",
		},
		{
			args:    args{data: map[string]interface{}{"arr": []interface{}{"s1", int64(2)}}},
			wantOut: `{"arr":["s1",2]}` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			if err := Encode(out, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOut := out.String(); gotOut != tt.wantOut {
				t.Errorf("Encode() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}
