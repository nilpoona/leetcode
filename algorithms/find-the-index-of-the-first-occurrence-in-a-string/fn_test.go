package find_the_index_of_the_first_occurrence_in_a_string

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				haystack: "sadbutsad",
				needle:   "sad",
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				haystack: "leetcode",
				needle:   "leeto",
			},
			want: -1,
		},
		{
			name: "3",
			args: args{
				haystack: "a",
				needle:   "b",
			},
			want: -1,
		},
		{
			name: "4",
			args: args{
				haystack: "a",
				needle:   "a",
			},
			want: 0,
		},
		{
			name: "5",
			args: args{
				haystack: "bkehikdaaaa",
				needle:   "a",
			},
			want: 7,
		},
		{
			name: "6",
			args: args{
				haystack: "abc",
				needle:   "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			},
			want: -1,
		},
		{
			name: "7",
			args: args{
				haystack: "bdekeaabc",
				needle:   "abc",
			},
			want: 6,
		},
		{
			name: "8",
			args: args{
				haystack: "bdekeaabc",
				needle:   "",
			},
			want: -1,
		},
		{
			name: "9",
			args: args{
				haystack: "mississippi",
				needle:   "issip",
			},
			want: 4,
		},
		{
			name: "10",
			args: args{
				haystack: "bfledekeababbabbabcedf",
				needle:   "abc",
			},
			want: 16,
		},
		{
			name: "11",
			args: args{
				haystack: "abbbbbaabbaabaabbbaaaaabbabbbabbbbbaababaabbaabbbbbababaababbbbaaabbbbabaabaaaabbbbabbbaabbbaabbaaabaabaaaaaaaa",
				needle:   "abbbaababbbabbbabbbbbabaaaaaaabaabaabbbbaabab",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
