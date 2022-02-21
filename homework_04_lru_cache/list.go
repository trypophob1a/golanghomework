package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	front *ListItem
	back  *ListItem
	size  int
}

func (l list) Len() int {
	return l.size
}

func (l list) IsEmpty() bool {
	return l.size < 1
}

func (l list) Front() *ListItem {
	return l.front
}

func (l list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	newListItem := ListItem{v, nil, nil}

	if l.IsEmpty() {
		l.back = &newListItem
	} else {
		currentItem := l.front
		newListItem.Next = currentItem
		currentItem.Prev = &newListItem
	}

	l.front = &newListItem
	l.size++

	return l.front
}
func (l *list) PushBack(v interface{}) *ListItem {
	newListItem := ListItem{v, nil, nil}

	if l.IsEmpty() {
		l.front = &newListItem
	} else {
		l.back.Next = &newListItem
		newListItem.Prev = l.back
	}

	l.back = &newListItem
	l.size++

	return l.back
}
func (l *list) Remove(i *ListItem) {
	if l.IsEmpty() {
		return
	}

	if l.size < 2 {
		l.front = nil
		l.back = nil
		l.size--
		return
	}

	if i.Prev == nil {
		l.RemoveFront()
		return
	}
	if i.Next == nil {
		l.RemoveBack()
		return
	}

	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev

	l.size--
}

func (l *list) RemoveFront() {
	if l.IsEmpty() {
		return
	}

	l.front = l.front.Next
	l.front.Prev = nil

	l.size--
}

func (l *list) RemoveBack() {
	if l.IsEmpty() {
		return
	}

	l.back.Prev.Next = nil
	l.back = l.back.Prev

	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev == nil {
		return
	}

	if i.Next == nil {
		front := l.front
		i.Prev.Next = nil
		l.back = i.Prev
		i.Prev = nil
		front.Prev = i
		i.Next = front
		l.front = i
		return

	}

	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
	i.Prev = nil
	front := l.front
	i.Next = front
	l.front = i
	front.Prev = l.front

}

func NewList() List {
	return &list{size: 0, back: nil, front: nil}
}
