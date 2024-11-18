package main
import (
    "fmt"
    "container/heap"
)

type Task struct {
    name string
    priority int
}

type PriorityQueue []Task

func (pq PriorityQueue) Len() int {return len(pq)}
func (pq PriorityQueue) Less(i,j int) bool {return pq[i].priority > pq[j].priority}
func (pq PriorityQueue) Swap(i,j int) {pq[i], pq[j]=pq[j],pq[i]}

func (pq *PriorityQueue) Push(x interface{}) {
    *pq = append(*pq, x.(Task))
}

func (pq *PriorityQueue) Pop() interface{} {
    old:=*pq
    n:=len(old)
    task:=old[n-1]
    *pq=old[0:n-1]
    return task
}

type TrieNode struct {
    children map[rune]*TrieNode
    isEnd bool
}

type Trie struct {
    root *TrieNode
}

func NewTrie() *Trie {
    return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (t *Trie) Insert(word string) {
    node:=t.root
    for _, char:=range word {
        if node.children[char]==nil {
            node=node.children[char]
        }
        node.isEnd=true
    }
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if node.children[char] == nil {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}


type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	queue    *list.List
}

type CacheItem struct {
	key   int
	value int
}


func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		queue:    list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	if element, exists := c.cache[key]; exists {
		c.queue.MoveToFront(element)
		return element.Value.(*CacheItem).value
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	if element, exists := c.cache[key]; exists {
		element.Value.(*CacheItem).value = value
		c.queue.MoveToFront(element)
		return
	}

	if len(c.cache) >= c.capacity {
		back := c.queue.Back()
		delete(c.cache, back.Value.(*CacheItem).key)
		c.queue.Remove(back)
	}

	newElement := c.queue.PushFront(&CacheItem{key: key, value: value})
	c.cache[key] = newElement
}


func main() {
    pq:=make(PriorityQueue,0)
    heap.Init(&pq)
	heap.Push(&pq, Task{"High Priority Task", 3})
	heap.Push(&pq, Task{"Low Priority Task", 1})
	trie := NewTrie()
	trie.Insert("golang")
	trie.Insert("python")
	fmt.Println(trie.Search("golang"))  
	fmt.Println(trie.Search("java"))    


	cache := NewLRUCache(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))  
}
