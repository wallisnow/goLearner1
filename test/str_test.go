package test

import (
	"base/src/Utils"
	"strconv"
	"testing"
)

//定义测试, 这里类似Junit中的外部提供的 Parameterized test
func TestJoin(t *testing.T) {
	type args struct {
		str []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//这里给定, 枚举值, 和期望值
		{"test1", args{[]string{"a", "b"}}, "ab"},
		{"test2", args{[]string{"c", "c"}}, "cc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Utils.Join(tt.args.str...); got != tt.want {
				t.Errorf("Join() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Utils.Join(strconv.Itoa(i))
	}
}

func BenchmarkGoJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Utils.GoJoin(strconv.Itoa(i))
	}
}

func BenchmarkBufferJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Utils.BufferJoin(strconv.Itoa(i))
	}
}
