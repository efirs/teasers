/* Write a function which takes two sorted int arrays as an input and returns
the median of combined array.

Ex.:
	Input: {1, 2, 5}, {4, 8}, result 4
	Input: {1, 2}, {4, 5}, result 3

Function signature:
	func median(ao []int, bo []int) int
*/

package teasers

import (
	"math/rand"
	"sort"
	"testing"
)

type test struct {
	median int
	a      []int
	b      []int
}

var tests = []test{
	{0, []int{}, []int{}},

	{1, []int{1}, []int{}},
	{1, []int{}, []int{1}},

	{2, []int{1, 3}, []int{}},
	{2, []int{}, []int{1, 3}},

	{2, []int{1}, []int{3}},
	{2, []int{3}, []int{1}},

	{3, []int{2, 1}, []int{4, 5}},
	{3, []int{4, 5}, []int{2, 1}},

	{3, []int{2, 1, 5}, []int{3, 8}},
	{3, []int{3, 8}, []int{2, 1, 5}},

	{4, []int{1, 100}, []int{2, 3, 5, 6}},
	{4, []int{2, 3, 5, 6}, []int{1, 100}},

	{4, []int{1, 100}, []int{2, 3, 4, 5, 6}},
	{4, []int{2, 3, 4, 5, 6}, []int{1, 100}},

	{4, []int{1, 3, 5, 7}, []int{2, 4, 6, 8}},
	{4, []int{2, 4, 6, 8}, []int{1, 3, 5, 7}},

	{5, []int{1, 3, 4, 5, 6, 7, 9}, []int{2, 8}},
	{5, []int{2, 8}, []int{1, 3, 4, 5, 6, 7, 9}},

	{1500, []int{1, 1000, 2000, 4000}, []int{500, 1500, 2500}},

	{8, []int{2, 1, 10, 100}, []int{6, 20}},

	{1, []int{1, 1}, []int{1, 1, 1}},
}

func TestMedian(t *testing.T) {
	for _, v := range tests {
		sort.Ints(v.a)
		sort.Ints(v.b)

		if v.median != median(v.a, v.b) {
			t.Fail()
		}
	}
}

func TestMedianSlow(t *testing.T) {
	for _, v := range tests {
		sort.Ints(v.a)
		sort.Ints(v.b)

		if v.median != medianSlow(v.a, v.b) {
			t.Fail()
		}
	}
}

func fillRandomArray(i int, j int) (*[]int, *[]int) {
	a := make([]int, i, i)
	b := make([]int, j, j)
	for k := 0; k < len(a); k++ {
		a[k] = 50000 - rand.Intn(100000)
	}
	for k := 0; k < len(b); k++ {
		b[k] = 50000 - rand.Intn(100000)
	}

	sort.Ints(a)
	sort.Ints(b)

	return &a, &b
}

func TestMedianHuge(t *testing.T) {
	a, b := fillRandomArray(100000, 10000)
	if medianSlow(*a, *b) != median(*a, *b) {
		t.Fail()
	}
}

func TestMedianBrute(t *testing.T) {
	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			a, b := fillRandomArray(i, j)
			if medianSlow(*a, *b) != median(*a, *b) {
				t.Fail()
			}
		}
	}
}

func BenchmarkMedianSlow(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := make([]int, 0, n)
		b := make([]int, 0, n)
		for i := 0; i < len(a); i++ {
			a[i] = i //rand.Intn(100000)
			b[i] = i //rand.Intn(100000)
		}
		medianSlow(a, b)
	}
}

func BenchmarkMedian(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := make([]int, 0, n)
		b := make([]int, 0, n)
		for i := 0; i < len(a); i++ {
			a[i] = i //rand.Intn(100000)
			b[i] = i //rand.Intn(100000)
		}
		median(a, b)
	}
}

/* linear soulution */
func medianSlow(ao []int, bo []int) int {
	if len(ao) == 0 && len(bo) == 0 {
		return 0
	}

	ptrA, ptrB, leftLen, prev, cur := 0, 0, 0, 0, 0

	/* merge like iterations till we reach half of the combined array */
	for leftLen != (len(ao)+len(bo))/2+1 {
		prev = cur
		if ptrA < len(ao) && (ptrB >= len(bo) || ao[ptrA] < bo[ptrB]) {
			cur = ao[ptrA]
			ptrA++
		} else {
			cur = bo[ptrB]
			ptrB++
		}
		leftLen++
	}

	if (len(ao)+len(bo))%2 == 1 {
		return cur
	}

	return (prev + cur) / 2
}
