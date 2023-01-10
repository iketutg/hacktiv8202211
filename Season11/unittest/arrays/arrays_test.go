package arrays

import (
	"github.com/adamluzsi/testcase"
	"testing"
)

func Sum(number []int) int {
	sum := 0
	for i := 0; i < len(number); i++ {
		sum += number[i]
	}
	return sum
}

func TestSumArray(t *testing.T) {
	spec := testcase.NewSpec(t)

	//lay out dari input
	var (
		numbers = testcase.Let[[]int](spec, func(t *testcase.T) []int {
			return []int{}
		})
		actual = func(t *testcase.T) int {
			return Sum(numbers.Get(t))
		}
	)

	spec.When("there's a array of numbers", func(s *testcase.Spec) {
		//arrange
		numbers.Let(s, func(t *testcase.T) []int {
			return []int{1, 2, 3, 4, 5, 6}
		})

		//assert
		s.Then("it sum theme all return value ", func(t *testcase.T) {
			t.Must.Equal(21, actual(t))
		})
	})

	spec.When("diference input array of numbers", func(s *testcase.Spec) {
		//arrange
		numbers.Let(s, func(t *testcase.T) []int {
			return []int{1, 2}
		})

		//assert
		s.Then("it sum theme all return value ", func(t *testcase.T) {
			t.Must.Equal(3, actual(t))
		})
	})

}
