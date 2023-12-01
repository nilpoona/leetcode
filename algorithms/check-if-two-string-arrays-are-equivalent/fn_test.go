package check_if_two_string_arrays_are_equivalent

import "testing"

func Test_arrayStringsAreEqual(t *testing.T) {
	type args struct {
		word1 []string
		word2 []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: struct {
				word1 []string
				word2 []string
			}{word1: []string{"ab", "c"}, word2: []string{"a", "bc"}},
			want: true,
		},
		{
			args: struct {
				word1 []string
				word2 []string
			}{word1: []string{"a", "cb"}, word2: []string{"ab", "c"}},
			want: false,
		},
		{
			args: struct {
				word1 []string
				word2 []string
			}{word1: []string{"abcd", "e", "fg"}, word2: []string{"abcdefg"}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayStringsAreEqual(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("arrayStringsAreEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
