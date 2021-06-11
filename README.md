# go-lru
Problem statement: Local Cache(LRU based) Implementation using golang.

About LRU:  Least Recently Used (LRU) Cache organizes items in order of use, allowing you to quickly identify which item hasn't been used for the longest amount of time.


Data structures :
1. Doubly linked list: for storing recent cache item.
2. Hash map: for maintaining key and node so we can access items in liked list in O(1) time.

space 	                       O(n)
get least recently used item 	O(1)
access item 	                O(1)

