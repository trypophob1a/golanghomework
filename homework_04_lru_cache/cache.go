package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	value, exists := l.items[key]

	if exists {
		if t, ok := value.Value.(cacheItem); ok {
			l.queue.MoveToFront(l.getNode(key))
			return t.value, exists
		}
	}
	return nil, exists
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	_, exists := l.Get(key)
	if exists {
		l.queue.Remove(l.getNode(key))
		l.items[key] = l.queue.PushFront(cacheItem{key, value})
		return exists
	}

	if l.capacity < l.queue.Len() {
		tail := l.queue.Back()
		l.queue.Remove(tail)

		if t, ok := tail.Value.(cacheItem); ok {
			delete(l.items, t.key)
		}

		l.items[key] = l.queue.PushFront(cacheItem{key, value})
		return exists
	}

	l.items[key] = l.queue.PushFront(cacheItem{key, value})

	return exists
}

func (l lruCache) getNode(key Key) *ListItem {
	value, exists := l.items[key]
	if exists {
		return value
	}
	return nil
}

func (l *lruCache) Clear() {
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
