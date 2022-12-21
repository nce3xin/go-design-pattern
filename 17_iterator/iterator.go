package iterator

type IIterator interface {
	HasNext() bool
	Next()                    // 用来将游标移动到下一个元素
	CurrentItem() interface{} // 用来获取当前游标指向的元素
}

type ArrayList []interface{}

func (a *ArrayList) Iterator() IIterator {
	return &ArrayListIterator{
		array:  a,
		cursor: 0,
	}
}

type ArrayListIterator struct {
	cursor int
	array  *ArrayList
}

func (a *ArrayListIterator) HasNext() bool {
	if a.cursor >= len(*a.array) {
		return false
	}
	return true
}

func (a *ArrayListIterator) Next() {
	a.cursor++
}

func (a *ArrayListIterator) CurrentItem() interface{} {
	return (*a.array)[a.cursor]
}
