package arraysandslices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSumArr(t *testing.T) {
	testcases := []struct {
		numbers [5]int
		sum     int
	}{
		{[5]int{1, 2, 3, 4, 5}, 15},
		{[5]int{-3, 4, 1, 13, 42}, 57},
		{[5]int{9, 20, 0, 0, 5}, 34},
		{[5]int{1, -9, 71, -6, -3}, 54},
	}

	for _, test := range testcases {
		s := SumArr(test.numbers)
		if s != test.sum {
			t.Errorf("Sum(%d): %d. Expected: %d, Got: %d", test.numbers, s, test.sum, s)
		}
	}
}

func TestSumAll(t *testing.T) {
	testcases := []struct{ s1, s2, sum []int }{
		{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{6, 8, 10, 12}},
		{[]int{-5, 12, 45, -322}, []int{2, -4, -55, -222}, []int{-3, 8, -10, -544}},
	}

	for _, test := range testcases {
		s := SumAll(test.s1, test.s2)
		if reflect.DeepEqual(s, test.sum) {
			t.Errorf("Sum(%d, %d): %d. Expected: %d, Got: %d", test.s1, test.s2, s, test.sum, s)
		}
	}
}

func TestSumSlice(t *testing.T) {
	testcases := []struct {
		numbers []int
		sum     int
	}{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{-3, 4, 1, 13, 42, -33}, 24},
		{[]int{9, 20, 0, 0, 5, -6, -54}, -26},
		{[]int{1, -9, 71, -6, -3}, 54},
	}

	for _, test := range testcases {
		s := SumSlice(test.numbers)
		if s != test.sum {
			t.Errorf("Sum(%d): %d. Expected: %d, Got: %d", test.numbers, s, test.sum, s)
		}
	}
}

func ExampleSumArr() {
	sum := SumArr([5]int{3, -4, 22, -108, 330})
	fmt.Println(sum)
	// Output: 243
}

func ExampleSumSlice() {
	sum := SumSlice([]int{3, -4, 22, -108, 330, 3, 1998, -33, -1085})
	fmt.Println(sum)
	// Output: 1126
}

func BenchmarkSumArr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumArr([5]int{311, -45, 272, -108, 330})
	}
}

func BenchmarkSumSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSlice([]int{3, -4, 22, -108, 330, 3, 1998, -33, -1085})
	}
}