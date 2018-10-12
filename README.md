# Go Data Structure And Algorithm
[![Build Status](https://travis-ci.org/liuwill/go-data-structure.svg?branch=master)](https://travis-ci.org/liuwill/go-data-structure)
[![codecov](https://codecov.io/gh/liuwill/go-data-structure/branch/master/graph/badge.svg)](https://codecov.io/gh/liuwill/go-data-structure)

> some data structure and Algorithm

## List

* 基数排序 [radix sort](./radix_sort/sort.go)
* 基数树 [radix tree](./radix_trees/tree.go)
* 插入排序 [insertion tree](./sorts/insert_sort.go)

### radix sort

radix sort[^1] is a non-comparative integer sorting algorithm that sorts data with integer keys by grouping keys by the individual digits which share the same significant position and value.

### radix tree

a radix tree[^2] (also radix trie or compact prefix tree) is a data structure that represents a space-optimized trie (prefix tree) in which each node that is the only child is merged with its parent. The result is that the number of children of every internal node is at most the radix r of the radix tree, where r is a positive integer and a power x of 2, having x ≥ 1. Unlike in regular tries, edges can be labeled with sequences of elements as well as single elements. This makes radix trees much more efficient for small sets (especially if the strings are long) and for sets of strings that share long prefixes.

### insertion tree

Insertion sort[^3] is a simple sorting algorithm that builds the final sorted array (or list) one item at a time.

* Simple implementation: Jon Bentley shows a three-line C version, and a five-line optimized version
* Efficient for (quite) small data sets, much like other quadratic sorting algorithms
* More efficient in practice than most other simple quadratic (i.e., O(n2)) algorithms such as selection sort or bubble sort
* Adaptive, i.e., efficient for data sets that are already substantially sorted: the time complexity is O(nk) when each element in the input is no more than k places away from its sorted position
* Stable; i.e., does not change the relative order of elements with equal keys
* In-place; i.e., only requires a constant amount O(1) of additional memory space
* Online; i.e., can sort a list as it receives it

[^1]: https://en.wikipedia.org/wiki/Radix_sort  "Radix Sort"
[^2]: https://en.wikipedia.org/wiki/Radix_tree  "Radix Tree"
[^3]: https://en.wikipedia.org/wiki/Insertion_sort  "Insertion Sort"
