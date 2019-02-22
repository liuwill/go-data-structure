# Go Data Structure And Algorithm
[![Build Status](https://travis-ci.org/liuwill/go-data-structure.svg?branch=master)](https://travis-ci.org/liuwill/go-data-structure)
[![codecov](https://codecov.io/gh/liuwill/go-data-structure/branch/master/graph/badge.svg)](https://codecov.io/gh/liuwill/go-data-structure)

> Some Data Structure and Algorithm

## List

* 基数排序 [radix sort](./radix_sort/sort.go)
* 基数树 [radix tree](./radix_trees/tree.go)
* 插入排序 [insertion tree](./sorts/insert_sort.go)

* Promise [Promise](./promise/future.go)
* Bloom filter [Bloom filter](./bloom_filter/filter.go)

### radix sort

radix sort[1] is a non-comparative integer sorting algorithm that sorts data with integer keys by grouping keys by the individual digits which share the same significant position and value.

### radix tree

a radix tree[2] (also radix trie or compact prefix tree) is a data structure that represents a space-optimized trie (prefix tree) in which each node that is the only child is merged with its parent. The result is that the number of children of every internal node is at most the radix r of the radix tree, where r is a positive integer and a power x of 2, having x ≥ 1. Unlike in regular tries, edges can be labeled with sequences of elements as well as single elements. This makes radix trees much more efficient for small sets (especially if the strings are long) and for sets of strings that share long prefixes.

### insertion tree

Insertion sort[3] is a simple sorting algorithm that builds the final sorted array (or list) one item at a time.

* Simple implementation: Jon Bentley shows a three-line C version, and a five-line optimized version
* Efficient for (quite) small data sets, much like other quadratic sorting algorithms
* More efficient in practice than most other simple quadratic (i.e., O(n2)) algorithms such as selection sort or bubble sort
* Adaptive, i.e., efficient for data sets that are already substantially sorted: the time complexity is O(nk) when each element in the input is no more than k places away from its sorted position
* Stable; i.e., does not change the relative order of elements with equal keys
* In-place; i.e., only requires a constant amount O(1) of additional memory space
* Online; i.e., can sort a list as it receives it

### skip list

Skip lists[4] are a data structure that can be used in place of balanced trees.
Skip lists use probabilistic balancing rather than strictly enforced balancing
and as a result the algorithms for insertion and deletion in skip lists are
much simpler and significantly faster than equivalent algorithms for
balanced trees.

skip list[5] allows fast search within an ordered sequence of elements. Fast search is made possible by maintaining a linked hierarchy of subsequences, with each successive subsequence skipping over fewer elements than the previous one (see the picture below on the right). Searching starts in the sparsest subsequence until two consecutive elements have been found, one smaller and one larger than or equal to the element searched for. Via the linked hierarchy, these two elements link to elements of the next sparsest subsequence, where searching is continued until finally we are searching in the full sequence. The elements that are skipped over may be chosen probabilistically  or deterministically, with the former being more common.

### Promise

The Promise[6] object represents the eventual completion (or failure) of an asynchronous operation, and its resulting value.

### Bloom filter

A Bloom filter[7] is a space-efficient probabilistic data structure, conceived by Burton Howard Bloom in 1970, that is used to test whether an element is a member of a set. False positive matches are possible, but false negatives are not – in other words, a query returns either "possibly in set" or "definitely not in set". Elements can be added to the set, but not removed (though this can be addressed with a "counting" filter); the more elements that are added to the set, the larger the probability of false positives.


[1]: https://en.wikipedia.org/wiki/Radix_sort  "Radix Sort"
[2]: https://en.wikipedia.org/wiki/Radix_tree  "Radix Tree"
[3]: https://en.wikipedia.org/wiki/Insertion_sort  "Insertion Sort"
[4]: https://www.cl.cam.ac.uk/teaching/0506/Algorithms/skiplists.pdf  "Skip List"
[5]: https://en.wikipedia.org/wiki/Skip_list  "Skip List"
[6]: https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise  "Promise"
[7]: https://en.wikipedia.org/wiki/Bloom_filter  "Bloom filter"
