package commons

import "testing"

func TestIsEven(t *testing.T) {
	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"IsEvenTrue", args{val: 10}, true},
		{"IsEvenFalse", args{val: 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEven(tt.args.val); got != tt.want {
				t.Errorf("IsEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
