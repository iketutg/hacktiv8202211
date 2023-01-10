package helpers

import (
	"fmt"
	"testing"
)

func TestSum01(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		panic("result should be 20")
	}
	fmt.Println("Testing Panic  Success")
}

func TestSum02(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		t.Fail()
	}
	fmt.Println("Testing Fail Success ")
}

func TestSum03(t *testing.T) {
	result := Sum(10, 10)

	if result != 20 {
		t.Error("Result has to be 20")
	}

	fmt.Println("Testing Error Success ")
}

func TestSum04(t *testing.T) {

	param1 := 10
	param2 := 10
	result := Sum(param1, param2)

	expetasi := param1 + param2

	if expetasi != result {
		t.Errorf("Expetasi : %d but result %d ", expetasi, result)
	}

	fmt.Println("Testing Errorf Success ")
}

func TestSum05(t *testing.T) {
	nametest := "Testing Run Sum 04 kondisi ..."
	param1 := 10
	param2 := -10
	expetasi := 0

	t.Run(nametest, func(t *testing.T) {
		result := Sum(param1, param2)
		if result != expetasi {
			t.Errorf("Expetasi : %d but result %d ", expetasi, result)
		}
	})

	fmt.Println("Testing Errorf Success ")
}

func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Testing Normal / Nilai Positif", args: args{a: 10, b: 10}, want: 20},
		{name: "Testing Normal / Nilai Negatif", args: args{a: 10, b: -10}, want: 0},
		{name: "Testing Normal / Nilai 0", args: args{a: 10, b: 0}, want: 0},
		{name: "Testing Normal / kurang dari 10", args: args{a: 5, b: 5}, want: 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
