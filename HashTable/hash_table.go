package hashtable

import "fmt"

// structure for holding hashtable
type Hashtable struct {
	array [10]*bucket
}

// structure for buckets in hashtable
type bucket struct {
	Head *bucketNode
}

// structure for each node
type bucketNode struct {
	key   string
	Value string
	next  *bucketNode
}

// inserts a key-value pair into hashtable
func (h *Hashtable) Insert(key string, value string) {
	index := Hash(key)
	h.array[index].insertIntoBucket(key, value)
}

// searches for a specific key in hashtable,
// returns a bool and correspoding value for key
func (h *Hashtable) Search(key string) (bool, string) {
	index := Hash(key)
	return h.array[index].searchInBucket(key)
}

// Deletes a key from hashtable
func (h *Hashtable) Delete(key string) {
	index := Hash(key)
	h.array[index].deleteFromBucket(key)
}

func (b *bucket) insertIntoBucket(key string, value string) {
	found, _ := b.searchInBucket(key)
	if !found {
		newNode := &bucketNode{key: key, Value: value}
		newNode.next = b.Head
		b.Head = newNode
	} else {
		fmt.Println(key + " already Exists")
	}

}

func (b *bucket) searchInBucket(k string) (bool, string) {
	current := b.Head
	for current != nil {
		if current.key == k {
			return true, current.Value
		}
		current = current.next
	}
	return false, ""
}

func (b *bucket) deleteFromBucket(key string) {
	if b.Head.key == key {
		b.Head = b.Head.next
		return
	}
	temp := b.Head
	for temp.next != nil {
		if temp.next.key == key {
			temp.next = temp.next.next
		}
		temp = temp.next
	}
}

// Hash function which uses division method for hashing
func Hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % 10
}

// initializer function which returns a new hash table instance
func Init() *Hashtable {
	result := &Hashtable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}
