package sum

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "sum 3 and 8",
			args: args{x: 3, y: 8},
			want: 11,
		},
		{
			name: "sum 3 and 8",
			args: args{x: 5, y: 5},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList(t *testing.T) {
	type args struct {
		list []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "list 1",
			args: args{[]int{2, 34, 5, 35, 23, 1}},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := List(tt.args.list); got != tt.want {
				t.Errorf("List() = %v, want %v", got, tt.want)
			}
		})
	}
}

//go test ./...
