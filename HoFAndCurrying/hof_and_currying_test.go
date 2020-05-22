package hofandcurrying

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSelectStrings(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"no strings given",
			args{nil},
			nil,
		},
		{
			"strings given with length over 3",
			args{[]string{"a", "ab", "abc", "abcd", "abcde"}},
			[]string{"abcd", "abcde"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectStrings(tt.args.strs)
			fmt.Println("got --- ", got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSelectStringsHof(t *testing.T) {
	strs := []string{"a", "ab", "bbc", "dbcd", "abcde"}

	type args struct {
		strs []string
		f    ChooseFnc
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"length over 3",
			args{
				strs,
				LengthOver(3),
			},
			[]string{"dbcd", "abcde"},
		},
		{
			"starts with 'a'",
			args{
				strs,
				StartsWith("a"),
			},
			[]string{"a", "ab", "abcde"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SelectStringsHof(tt.args.strs, tt.args.f)
			fmt.Println("SelectStringsHof -- got ", got)
			if  !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectStringsHof() = %v, want %v", got, tt.want)
			}
		})
	}
}
