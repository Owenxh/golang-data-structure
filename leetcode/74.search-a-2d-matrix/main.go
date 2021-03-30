package main

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) < 1 {
		return false
	}
	return search(matrix, target, 0, len(matrix)-1)
}

func search(matrix [][]int, target int, h, t int) bool {
	if h > t {
		return false
	}
	m := (h + t) / 2
	if matrix[m][0] > target {
		return search(matrix, target, h, m-1)
	} else if matrix[m][len(matrix[m])-1] < target {
		return search(matrix, target, m+1, t)
	} else {
		return searchInRow(matrix[m], target, 0, len(matrix[m])-1)
	}
}

func searchInRow(arr []int, target, l, r int) bool {
	if l > r {
		return false
	}
	m := (l + r) / 2
	if arr[m] == target {
		return true
	} else if arr[m] > target {
		return searchInRow(arr, target, l, m-1)
	} else {
		return searchInRow(arr, target, m+1, r)
	}
}

func main() {
	searchMatrix([][]int{{-10, -10}, {-9, -9}, {-8, -6}, {-4, -2}, {0, 1}, {3, 3}, {5, 5}, {6, 8}}, 0)
}
