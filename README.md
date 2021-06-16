# GOLRU
# Local Cache(LRU based) Implementation using golang.

About LRU:  Least Recently Used (LRU) Cache organizes items in order of use, allowing you to quickly identify which item hasn't been used for the longest amount of time.

# Commmand 

1. git clone https://github.com/Yadavshivpal/GOLRU

2. cd GOLRU

3. go run client.go

# Data structures :
1. Doubly linked list: for storing recent cache item.
2. Hash map: for maintaining key and node so we can access items in liked list in O(1) time.

space 	                       O(n)
get least recently used item 	O(1)
access item 	                O(1)

![alt text](https://www.interviewcake.com/images/svgs/lru_cache__vanilla_cake_recipe_request_response.svg?bust=209)
![alt text](https://www.interviewcake.com/images/svgs/lru_cache__doubly_linked_list.svg?bust=209)


# Reference
1. https://www.interviewcake.com/concept/java/lru-cache for LRU algo
2. https://www.youtube.com/watch?v=9yjQ2DV5rmY for LRU algo
3. https://www.sibis.dev/linked-lists-in-golang-explained-with-real-world-application for linked list
4. https://gobyexample.com/maps
