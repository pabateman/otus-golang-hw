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
	Size      int
	FrontItem *ListItem
	BackItem  *ListItem
}

func (l *list) Len() int {
	return l.Size
}

func (l *list) Front() *ListItem {
	return l.FrontItem
}

func (l *list) Back() *ListItem {
	return l.BackItem
}

func (l *list) PushFront(value interface{}) *ListItem {
	newFront := &ListItem{
		Value: value,
		Prev:  nil,
		Next:  l.FrontItem,
	}

	l.FrontItem = newFront

	if l.FrontItem.Next == nil {
		l.BackItem = l.FrontItem
	} else {
		l.FrontItem.Next.Prev = l.FrontItem
	}

	l.Size++
	return l.FrontItem
}

func (l *list) PushBack(value interface{}) *ListItem {
	newBack := &ListItem{
		Value: value,
		Prev:  l.BackItem,
		Next:  nil,
	}

	l.BackItem = newBack

	if l.BackItem.Prev == nil {
		l.FrontItem = l.BackItem
	} else {
		l.BackItem.Prev.Next = l.BackItem
	}

	l.Size++
	return l.BackItem
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	l.Size--

	if l.Size == 0 {
		l.FrontItem = nil
		l.BackItem = nil
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		return
	}

	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	oldFront := l.FrontItem
	oldFront.Prev = i

	i.Prev = nil
	i.Next = oldFront

	l.FrontItem = i
}

func NewList() List {
	return new(list)
}
