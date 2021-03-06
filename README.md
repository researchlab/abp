# Algorithms Best Practices (ABP™)


### sorting
| algorithms| en-wiki|zh-wiki|
|---|---|---|
| [BubbleSort](https://github.com/researchlab/algorithms-cs/tree/master/sorting/bubble)| [en-wiki](https://en.wikipedia.org/wiki/Bubble_sort) |[zh-wiki](https://zh.wikipedia.org/wiki/%E5%86%92%E6%B3%A1%E6%8E%92%E5%BA%8F)|
|[SelectionSort](https://github.com/researchlab/algorithms-cs/tree/master/sorting/selection)| [en-wiki](https://en.wikipedia.org/wiki/Selection_sort) |[zh-wiki](https://zh.wikipedia.org/wiki/%E9%80%89%E6%8B%A9%E6%8E%92%E5%BA%8F)|
|[InsertSort](https://github.com/researchlab/algorithms-cs/tree/master/sorting/insert)| [en-wiki](https://en.wikipedia.org/wiki/Insertion_sort) |[zh-wiki](https://zh.wikipedia.org/wiki/%E6%8F%92%E5%85%A5%E6%8E%92%E5%BA%8F)|
|[MergeSort](https://github.com/researchlab/algorithms-cs/tree/master/sorting/merge)| [en-wiki](https://en.wikipedia.org/wiki/Merge_sort) |[zh-wiki](https://zh.wikipedia.org/wiki/%E5%BD%92%E5%B9%B6%E6%8E%92%E5%BA%8F)|
|[QuickSort](https://github.com/researchlab/algorithms-cs/tree/master/sorting/quicksort)| [en-wiki](https://en.wikipedia.org/wiki/Quicksort) |[zh-wiki](https://zh.wikipedia.org/wiki/%E5%BF%AB%E9%80%9F%E6%8E%92%E5%BA%8F)|
|[HeapSort](https://github.com/researchlab/algorithms-cs/tree/master/sorting/heap)| [en-wiki](https://en.wikipedia.org/wiki/Heapsort) |[zh-wiki](https://zh.wikipedia.org/wiki/%E5%A0%86%E6%8E%92%E5%BA%8F)|
|[BucketSort](https://github.com/researchlab/algorithms-cs/tree/master/sorting/bucketsort)| [en-wiki](https://en.wikipedia.org/wiki/Bucket_sort) |[zh-wiki](https://zh.wikipedia.org/wiki/%E6%A1%B6%E6%8E%92%E5%BA%8F)
|[RadixSort](https://github.com/researchlab/algorithms-cs/tree/master/sorting/radixsort)| [en-wiki](https://en.wikipedia.org/wiki/Radix_sort) |[zh-wiki](https://zh.wikipedia.org/wiki/%E5%9F%BA%E6%95%B0%E6%8E%92%E5%BA%8F)|


### Cache replacement policies
  - theory: [Cache replacement policies](https://en.wikipedia.org/wiki/Cache_replacement_policies#Least_recently_used_(LRU))
  - algorithms:
     - [lru/arc/2q](https://github.com/hashicorp/golang-lru)
     - [lru from groupcache src](https://github.com/golang/groupcache)
     - [lru/lfu/arc](https://github.com/bluele/gcache)
     - [lfu](https://github.com/mtchavez/lfu)
     - [lfu](https://github.com/dgrijalva/lfu-go)
		 - [http://dhruvbird.com/lfu.pdf](http://dhruvbird.com/lfu.pdf)
