package sort

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func __sortData() [][]int {
	return [][]int{
		randomIntSliceWithBound(100000, 100000),
		randomIntSliceWithBound(1000000, 1000000),
		randomIntSliceWithBound(2000000, 2000000),
		randomIntSliceWithBound(5000000, 5000000),
		randomIntSliceWithBound(10000000, 10000000),
		randomIntSlice(10000000),
		orderedIntSlice(10000000, 10000000),
		randomIntSliceWithBound(10000000, 1),
	}
}

func Test__QuickSort3Ways(t *testing.T) {
	data := __sortData()
	applySortTest(__QuickSort3Ways, t, data)
}

func __QuickSort3Ways(arr []int) {
	__quickSort3Ways(arr, 0, len(arr)-1)
}

func __quickSort3Ways(arr []int, l, r int) {
	if l >= r {
		return
	}

	lt, gt := __partition3Ways(arr, l, r)

	__quickSort3Ways(arr, l, lt-1)
	__quickSort3Ways(arr, gt, r)
}

// loop invariant
// ① arr[l] 标定点元素
// ② arr[l+1, lt] < v && arr[lt+1, i-1] == v && [gt, r] > v
// ③ [i, gt-1] 待处理的区间
func __partition3Ways(arr []int, l, r int) (int, int) {
	k := l + rand.Intn(r-l+1)
	arr[l], arr[k] = arr[k], arr[l]

	v := arr[l]

	// 初始变量阶段
	// ① arr[l+1, lt] 为 [l+1, l] 左闭 > 右闭的空区间，满足循环不变量
	// ② arr[gt,   r] 为 [r+1, r] 左闭 > 右闭的空区间，满足循环不变量
	lt, gt := l, r+1

	for i := lt + 1; i < gt; {
		if arr[i] < v {
			lt++
			arr[i], arr[lt] = arr[lt], arr[i]
			i++
		} else if arr[i] > v {
			gt--
			arr[i], arr[gt] = arr[gt], arr[i]
		} else {
			i++
		}
	}

	// 循环之后
	// ① [l+1, lt] < v
	// ② [lt + 1, gt - 1] == v
	// ③ [gt, r] > v

	arr[l], arr[lt] = arr[lt], arr[l]
	// ① arr[l, lt-1] < v
	// ② arr[lt, gt-1] == v
	// ③ [gt, r] > v
	return lt, gt
}

func Test__QuickSort2Ways(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())
	applySortTest(__QuickSort2Ways, t, [][]int{randomIntSliceWithBound(8, 100)})
}

func __QuickSort2Ways(arr []int) {
	__quickSort2Ways(arr, 0, len(arr)-1)
}

func __quickSort2Ways(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := __partition2Ways(arr, l, r)
	__quickSort2Ways(arr, l, p-1)
	__quickSort2Ways(arr, p+1, r)
}

// loop invariant
// ① arr[l] 标定点元素
// ② arr[l+1, i-1] < v && [j+1, r] > v
// ③ i <= j
//
// 初始阶段
// i = l, j = r
// arr[l+1, l] 空区间满足 arr[l+1, i-1]
// arr[r+1, r]空区间满足 [j+1, r] > v
func __partition2Ways(arr []int, l, r int) int {
	fmt.Println("start partition")
	fmt.Println(arr[l : r+1])

	k := l + rand.Intn(r-l+1)
	arr[l], arr[k] = arr[k], arr[l]

	fmt.Println("after random")
	fmt.Println(arr[l : r+1])

	v := arr[l]
	i, j := l+1, r
	for {
		for i <= j && arr[i] < v {
			i++
		}
		for j >= i && arr[j] > v {
			j--
		}

		if i >= j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--

		fmt.Println("after swap i & j")
		fmt.Println(arr[l : r+1])
	}
	arr[l], arr[j] = arr[j], arr[l]

	fmt.Println("after partition")
	fmt.Println(arr[l : r+1])
	return j
}

func Test__MergeSort(t *testing.T) {
	data := __sortData()
	applySortTest(__MergeSort, t, data)
}

func __MergeSort(arr []int) {
	temp := make([]int, len(arr), len(arr))
	CopyArray(arr, 0, len(arr)-1, temp)
	__mergeSort(arr, 0, len(arr)-1, temp)
}

func __mergeSort(arr []int, l, r int, temp []int) {
	if l >= r {
		return
	}

	m := (r-l)/2 + l
	__mergeSort(arr, l, m, temp)
	__mergeSort(arr, m+1, r, temp)
	if arr[m] > arr[m+1] {
		__merge(arr, l, m, r, temp)
	}
}

// 循环不变量
// ① l <= i <= m
// ② m < j <= r
// ③ i <= k <= r
// ④ arr[k - 1] <= arr[k]
func __merge(arr []int, l, m, r int, temp []int) {
	CopyArray(arr, l, r, temp)
	i, j, k := l, m+1, l
	for ; i <= m || j <= r; k++ {
		if j > r || (i <= m && temp[i] <= temp[j]) {
			arr[k] = temp[i]
			i++
		} else {
			arr[k] = temp[j]
			j++
		}
	}
}

func Test__bubbleSort(t *testing.T) {
	applySortTest(__bubbleSort, t, [][]int{randomIntSliceWithBound(8, 100)})
}

func __bubbleSort(arr []int) {
	fmt.Println(arr)

	sorted := false

	for i := 0; i < len(arr) && !sorted; i++ {
		sorted = true
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
				sorted = false
			}
		}
		fmt.Println(arr)
	}
}

func Test__selectSort(t *testing.T) {
	applySortTest(__selectSort, t, [][]int{randomIntSliceWithBound(8, 100)})
}

func __selectSort(arr []int) {
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		p := i

		fmt.Println(arr[:p])

		for j := i; j < len(arr); j++ {
			if arr[j] < arr[p] {
				p = j
			}
		}
		arr[i], arr[p] = arr[p], arr[i]
		fmt.Println(arr)
	}
}

func Test__insertSort(t *testing.T) {
	applySortTest(__insertSort, t, [][]int{randomIntSliceWithBound(8, 100)})
}

// 循环不变量
// arr[0, i-1] 是有序的，将arr[i] 向 arr[0, i-1] 中插入，使 arr[0, i] 有序
func __insertSort(arr []int) {
	fmt.Println(arr)

	for i := 1; i < len(arr); i++ {
		// 已排好序的
		fmt.Println(arr[:i])

		v, p := arr[i], i
		for ; p > 0 && v < arr[p-1]; p-- {
			arr[p] = arr[p-1]
		}
		// 找到 p 在排序后的数组中的位置
		arr[p] = v

		fmt.Println(arr)
	}
}

func Test__QuickSort(t *testing.T) {
	rand.Seed(time.Now().UnixMilli())
	applySortTest(__QuickSort, t, [][]int{randomIntSliceWithBound(8, 100)})
}

func __QuickSort(arr []int) {
	__quickSort(arr, 0, len(arr)-1)
}

func __quickSort(arr []int, l, r int) {
	if l >= r {
		return
	}
	p := __partition(arr, l, r)
	__quickSort(arr, l, p-1)
	__quickSort(arr, p+1, r)
}

// loop invariant
// [l+1, p] < v && [p + 1, i-1] > v
// [i, r] 待处理的区间
func __partition(arr []int, l, r int) int {
	fmt.Println("start partition")
	fmt.Println(arr[l : r+1])

	k := l + rand.Intn(r-l+1)
	arr[l], arr[k] = arr[k], arr[l]

	fmt.Println("after random")
	fmt.Println(arr[l : r+1])

	v, p := arr[l], l
	for i := l + 1; i <= r; i++ {
		if arr[i] < v {
			p++
			arr[p], arr[i] = arr[i], arr[p]
		}
	}
	arr[l], arr[p] = arr[p], arr[l]

	fmt.Println("after partition")
	fmt.Println(arr[l : r+1])
	return p
}
