package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Array demo addition", func(t *testing.T) {
		numbers := []int{1,2,3,4,5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got: %d want: %d given: %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("Slices variation addition", func(t *testing.T) {
		nums := []int{1,1,1}
		nums1 := []int{6,78,8}

		got := SumAll(nums,nums1)
		want := []int{3, 92}

		if !slices.Equal(got, want) {
			t.Errorf("got: %v want: %v given: %v", got, want, nums)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func (t testing.TB, got, want []int)  {
		t.Helper()
		if !slices.Equal(got, want){
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Adds all slices except first", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0,9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})
	t.Run("passing empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3,4,5})
		want := []int{0,9}

		checkSums(t, got, want)
	})
}
