# Play Data Structure & Algorithm with Golang

使用 `Golang` 实现常用的数据结构和算法。

### 排序算法
* [选择排序](./sort/selectsort.go)
* [插入排序](./sort/insertionsort.go)
* [冒泡排序](./sort/bubblesort.go)
* [归并排序](./sort/mergesort.go)
* [希尔排序](./sort/shellsort.go)

**归并排序**
* [自顶向下](./sort/mergesort.go)
* [自底向上](./sort/mergesortbu.go)

**快速排序**
* [一路快排](./sort/quicksort1way.go)
* [二路快排](./sort/quicksort2ways.go)
* [三路速排](./sort/quicksort3ways.go)

### 排序算法对比
$$
\begin{array}{|c|c|c|c|}
\hline
{算法名称} & {时间复杂度} & {空间复杂度} & {特殊用例} \\
\hline
{选择排序} & {O(n^2)} & {O(1)} & {} \\
\hline
{插入排序} & {O(n^2)} & {O(1)} & {数组完全有序时，时间复杂度 O(n) } \\
\hline
{冒泡排序} & {O(n^{2})} & {O(1)} & {} \\
\hline
{归并排序} & {O(nlogN)} & {O(n)} & {数组完全有序时，时间复杂度 O(n)} \\
\hline
{快速排序} & {O(nlogN)*} & {O(1)} & {数组元素完全相同时，三路快排时间复杂度为 O(n)} \\
\hline
{堆排序} & {O(nlogN)} & {O(1)} & {} \\
\hline
{希尔排序} & {O(nlogN)-O(n^{2})} & {O(1)} & {} \\
\hline
\end{array}
$$

### 非比较排序
* [LSD](./sort/radix/lsd.go)
* [MSD](./sort/radix/msd.go)
* [桶排序](./sort/radix/bucket_sort.go)

### 二分搜索 - Binary Search
* [二分搜索](./search/binary_search.go)

### 动态数组 - Dynamic Array
* [动态数组](./array/array.go)

### 链表 - Linked List

### 栈 - Stack
* [动态数组实现栈](./stack/array_stack.go)

### 队列 - Queue
* [动态数组实现队列](./queue/array_queue.go)
* [循环数组实现队列](./queue/loop_queue.go)

### 堆 - Heap
* [最大堆](./tree/max_heap.go)
* [最小堆](./tree/min_heap.go)

### 线段树 - Segment Tree
* [线段树](./tree/segment_tree.go)
* [懒更新线段树](./tree/lazy_segment_tree.go)

### 字典树 - Trie
* [字典树](./tree/trie.go)

### 并查集 - Union Find
* [Union Find](./tree/union_find.go)

### 树状数组 - Binary Indexed Tree
* [树状数组](./tree/binary_indexed_tree.go)


### 字符匹配 - String Match
* [暴力破解法](./strings/match/brute_force.go)
* [Rabin Karp](./strings/match/rabin_karp.go)
* [KMP](./strings/match/kmp.go)

### 图 Graph
* [图的数据结构表示](./graph/)
* [深度优先遍历 DFS](./graph/dfs/)
* [广度优先遍历 BFS](./graph/bfs/)

### 其他
* [前缀和](./util/presum/)
