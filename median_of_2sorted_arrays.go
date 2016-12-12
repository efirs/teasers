package teasers

import (
	"sort"
)

type arrctx struct {
	v             []int
	start, end    int
	median, image int
}

// calcMedianAndImage finds median of a and image of a's median in b
func calcMedianAndImage(a *arrctx, b *arrctx) {
	if a.end-a.start != 0 {
		a.median = a.start + (a.end-a.start)/2
		a.image = b.start + sort.SearchInts(b.v[b.start:b.end], a.v[a.median])
	}
}

func cutLeft(leftLen int, cur int, prev int, a *arrctx, b *arrctx) (int, int, int) {
	cutLen := (a.median - a.start) + (a.image - b.start)

	/* to ensure progress cut one element(median of the source), when
	nothing to cut. ex: single element arrays input []int{1}, []int{3}*/
	if cutLen == 0 {
		a.start = a.median + 1
		leftLen++
		prev = cur
		cur = a.v[a.median]
	} else if leftLen+cutLen > (len(a.v)+len(b.v))/2+1 {
		/* if sum of what we cut earlier and what we are going to cut now
		is greater then half of the combined length then we are discarding
		right parts, because median cannot be there */
		a.end = a.median
		b.end = a.image
	} else {
		/* remember previous in order element if we are cutting only one element.
		it was set by previous cut */
		prev = cur
		/* find maximum element in current cut and if we cut more then
		one element find previous in order also */
		if a.median-a.start > 0 && (a.image-b.start == 0 || b.v[a.image-1] < a.v[a.median-1]) {
			cur = a.v[a.median-1]
			if a.median-a.start > 1 && (a.image-b.start == 0 || b.v[a.image-1] < a.v[a.median-2]) {
				prev = a.v[a.median-2]
			} else if a.image-b.start != 0 {
				prev = b.v[a.image-1]
			}
		} else {
			cur = b.v[a.image-1]
			if a.image-b.start > 1 && (a.median-a.start == 0 || b.v[a.image-2] > a.v[a.median-1]) {
				prev = b.v[a.image-2]
			} else if a.median-a.start != 0 {
				prev = a.v[a.median-1]
			}
		}
		/* leave current cut on the left */
		a.start = a.median
		b.start = a.image
		leftLen += cutLen
	}

	return leftLen, cur, prev
}

/* log(N+M) solution */
func median(ao []int, bo []int) int {

	if len(ao) == 0 && len(bo) == 0 {
		return 0
	}

	leftLen, prev, cur := 0, 0, 0
	a, b := &arrctx{ao, 0, len(ao), 0, 0}, &arrctx{bo, 0, len(bo), 0, 0}

	for leftLen != (len(ao)+len(bo))/2+1 {
		calcMedianAndImage(a, b)
		calcMedianAndImage(b, a)

		/* cut left part based on the selected array with smaller median */
		if a.end-a.start != 0 && (b.end-b.start == 0 || a.v[a.median] < b.v[b.median]) {
			leftLen, cur, prev = cutLeft(leftLen, cur, prev, a, b)
		} else {
			leftLen, cur, prev = cutLeft(leftLen, cur, prev, b, a)
		}
	}

	if (len(ao)+len(bo))%2 == 1 {
		return cur
	}

	return (cur + prev) / 2
}

/*
func printrray(a *arrctx) {
	fmt.Print("[")
	for i, v := range a.v {
		if i == a.start {
			fmt.Print("<")
		}
		if i == a.end {
			fmt.Print(">")
		}
		if i == a.median {
			fmt.Print("|")
		}
		if i != a.end && i != a.median && i != len(a.v)-1 {
			fmt.Print(v, ",")
		}
	}
	if len(a.v) == a.start {
		fmt.Print("<")
	}
	if len(a.v) == a.end {
		fmt.Print(">")
	}
	if len(a.v) == a.median {
		fmt.Print("|")
	}
	fmt.Println("]")
}
*/
