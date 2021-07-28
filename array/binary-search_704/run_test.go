package binary_search_704

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				nums:   []int{-1, 2, 5, 8, 9, 20, 24, 36, 67, 99, 100},
				target: 20,
			},
			want: 5,
		},
		{
			name: "test1",
			args: args{
				nums:   []int{-1, 2, 5, 8, 9, 20, 24, 36, 67, 99, 100},
				target: 21,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
