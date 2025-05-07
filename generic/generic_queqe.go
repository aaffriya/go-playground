package _generic
import . "fmt"
type Queue[T interface{ ~string | ~[]any | ~int }] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
    q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
    if len(q.items) == 0 {
        var zero T
        return zero, false
    }
    item := q.items[0]
    q.items = q.items[1:]
    return item, true
}

func (q *Queue[T]) IsEmpty() bool {
		return len(q.items) == 0
}

func (q *Queue[T]) Size() int {
		return len(q.items)
}

func (q *Queue[T]) Peek() (T, bool) {
		if len(q.items) == 0 {
				var zero T
				return zero, false
		}
		return q.items[0], true
}
func (q *Queue[T]) Clear() {
		q.items = []T{}
}
func (q *Queue[T]) Print() {
		for _, item := range q.items {
				println(item)
		}
}
func (q *Queue[T]) PrintWithIndex() {
		for i, item := range q.items {
				println(i, item)
		}
}
func (q *Queue[T]) PrintWithIndexAndType() {
		for i, item := range q.items {
				println(i, item, "Type:", Sprintf("%T", item))
		}
}
func (q *Queue[T]) PrintWithIndexAndTypeAndValue() {
		for i, item := range q.items {
				println(i, item, "Type:", Sprintf("%T", item), "Value:", item)
		}
}

func StringQueue() {
		q := Queue[string]{}
		q.Enqueue("Hello")
		q.Enqueue("World")
		q.Enqueue("!")
		q.Print()
		println("Size:", q.Size())
		println("IsEmpty:", q.IsEmpty())
		item, ok := q.Dequeue()
		if ok {
				println("Dequeued:", item)
		} else {
				println("Queue is empty")
		}
		q.Print()
		println("Size:", q.Size())
		println("IsEmpty:", q.IsEmpty())
}

func IntQueue() {
		q := Queue[int]{}
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)
		q.Print()
		println("Size:", q.Size())
		println("IsEmpty:", q.IsEmpty())
		item, ok := q.Dequeue()
		if ok {
				println("Dequeued:", item)
		} else {
				println("Queue is empty")
		}
		q.Print()
		println("Size:", q.Size())
		println("IsEmpty:", q.IsEmpty())
}