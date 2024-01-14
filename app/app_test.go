package app

import "testing"

func Test_findIndex(t *testing.T) {
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
			name: "1",
			args: args{
				nums:   []int{0, 10, 20, 100},
				target: 100,
			},
			want: 3,
		},
		{
			name: "2",
			args: args{
				nums:   []int{0, 10, 20, 100},
				target: 10,
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				nums:   []int{0, 100, 200},
				target: 111,
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				nums:   []int{0, 1, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},
		{
			name: "5",
			args: args{
				nums:   []int{21},
				target: 20,
			},
			want: 0,
		},
		{
			name: "6",
			args: args{
				nums:   []int{0, 100, 200},
				target: 110,
			},
			want: 1,
		},
		{
			name: "7",
			args: args{
				nums:   []int{0, 10, 20, 100},
				target: 15,
			},
			want: -1,
		},
		{
			name: "8",
			args: args{
				nums:   []int{0, 10, 20, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200},
				target: 1150,
			},
			want: 13,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if got := findIndex(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("findIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
