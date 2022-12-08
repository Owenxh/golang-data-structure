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

**排序算法对比**


| 算法名称 |     时间复杂度     | 空间复杂度 | 特殊用例                                         |
| :------: | :-----------------: | :--------: | :----------------------------------------------- |
| 选择排序 |     ${O(n^2)}$     |  ${O(1)}$  | 数组完全有序时，时间复杂度${O(n)}$               |
| 插入排序 |     ${O(n^2)}$     |  ${O(1)}$  |                                                  |
| 冒泡排序 |     ${O(n^2)}$     |  ${O(1)}$  |                                                  |
| 归并排序 |     $O(nlogN)$     |  ${O(n)}$  | 数组完全有序时，时间复杂度${O(n)}$               |
| 快速排序 |     $O(nlogN)*$     |  ${O(1)}$  | 数组元素完全相同时，三路快排时间复杂度为${O(n)}$ |
|  堆排序  |     $O(nlogN)$     |  ${O(1)}$  |                                                  |
| 希尔排序 | $O(nlogN)-O(n^{2})$ |  ${O(1)}$  |                                                  |

### 非比较排序

* [LSD](./sort/radix/lsd.go)
* [MSD](./sort/radix/msd.go)
* [桶排序](./sort/radix/bucket_sort.go)

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

### 二叉搜索树 - Binary Search Tree

* [二叉搜索树](./tree/binary_search_tree.go)
* [自平衡二叉树 AVL](./tree/avl_tree.go)

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

### 图 - Graph

* [图的数据结构表示](./graph/)

**无向无权图**

* [深度优先遍历 - DFS](./graph/dfs/)
* [连通分量](./graph/dfs/cc_count.go)
* [单一路径](./graph/dfs/path.go)
* [环检测](./graph/dfs/cycle_detection.go)
* [二分图检测](./graph/dfs/bipartition_detection.go)
* [广度优先遍历 - BFS](./graph/bfs/)
* [哈密尔顿回路 - Hamilton Loop](graph/path/hamilton_loop.go)
* [欧拉回路 - Euler Loop](graph/path/euler_loop.go)

**无向有权图**

* [最小生成树 - Kruskal 算法](./graph/mst/kruskal.go) $O(ElogE)$
* [最小生成树 - Prime 算法](./graph/mst/prime.go) $O(ElogE)$
* [迪杰斯特拉 - Dijkstra 算法](./graph/mst/prime.go) $O(ElogE)$

### 其他

* [二分查找 Binay Search](./search/binary_search.go)
* [前缀和 Pre-Sum](./util/presum.go)
