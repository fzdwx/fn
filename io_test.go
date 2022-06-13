package fn

import (
	"bytes"
	"reflect"
	"testing"
)

func TestWrap(t *testing.T) {
	type args struct {
		bys []byte
	}
	tests := []struct {
		name string
		args args
		want *bytes.Buffer
	}{
		{name: "有值", args: args{bys: []byte("hello")}, want: bytes.NewBuffer([]byte("hello"))},
		{name: "null", args: args{bys: nil}, want: bytes.NewBuffer(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Wrap(tt.args.bys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrapStr(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want *bytes.Buffer
	}{
		{name: "有值", args: args{s: "hello"}, want: bytes.NewBufferString("hello")},
		{name: "null", args: args{s: ""}, want: bytes.NewBufferString("")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WrapStr(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
