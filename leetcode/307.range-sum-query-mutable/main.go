package main

type NumArray struct {
	size int
	tree []int
}

func Constructor(nums []int) NumArray {
	size := len(nums)
	//h := math.Ceil(math.Log2(float64(size)))
	//sz := int(2*math.Pow(2, h) - 1)
	sz := 4 * size
	tree := make([]int, sz, sz)

	st := NumArray{
		size: size,
		tree: tree,
	}

	buildSegmentTree(&st, nums, 0, 0, size-1)
	return st
}

func leftChild(i int) int {
	return i*2 + 1
}

func rightChild(i int) int {
	return i*2 + 2
}

func buildSegmentTree(st *NumArray, nums []int, i int, l int, r int) {
	if l == r {
		st.tree[i] = nums[l]
		return
	}

	left, right := leftChild(i), rightChild(i)

	m := ((r - l) >> 1) + l

	buildSegmentTree(st, nums, left, l, m)
	buildSegmentTree(st, nums, right, m+1, r)

	st.tree[i] = st.tree[left] + st.tree[right]
}

func (st *NumArray) Update(index int, val int) {
	if index < 0 || index >= st.size {
		return
	}
	st.update(0, 0, st.size-1, index, val)
}

func (st *NumArray) update(treeIndex int, l int, r int, index int, val int) {
	if l == r && index == l {
		st.tree[treeIndex] = val
		return
	}

	left, right := leftChild(treeIndex), rightChild(treeIndex)

	m := ((r - l) >> 1) + l
	if m >= index {
		st.update(left, l, m, index, val)
	} else {
		st.update(right, m+1, r, index, val)
	}

	st.tree[treeIndex] = st.tree[left] + st.tree[right]
}

func (st *NumArray) SumRange(left int, right int) int {
	if left < 0 || right >= st.size || left > right {
		return 0
	}
	return st.sumRange(0, 0, st.size-1, left, right)
}

func (st *NumArray) sumRange(treeIndex int, l int, r int, qL int, qR int) int {
	if l >= qL && r <= qR {
		return st.tree[treeIndex]
	}

	left, right := leftChild(treeIndex), rightChild(treeIndex)

	m := ((r - l) >> 1) + l

	var sum int
	if qL <= m {
		sum += st.sumRange(left, l, m, qL, qR)
	}
	if qR > m {
		sum += st.sumRange(right, m+1, r, qL, qR)
	}

	return sum
}

func main() {

}
