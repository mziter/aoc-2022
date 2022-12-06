package ringbuf

type Ringbuf[T any] struct {
	Data []T
	next int
}

func New[T any](capacity int) *Ringbuf[T] {
	return &Ringbuf[T]{
		Data: make([]T, capacity, capacity),
	}
}

func (b *Ringbuf[T]) Add(item T) T {
	oldVal := b.Data[b.next]
	b.Data[b.next] = item
	b.setNext()
	return oldVal
}

func (b *Ringbuf[T]) Length() int {
	return len(b.Data)
}

func (b *Ringbuf[T]) setNext() {
	b.next++
	if b.next == cap(b.Data) {
		b.next = 0
	}
}
