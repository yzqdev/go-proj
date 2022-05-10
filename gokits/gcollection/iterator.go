package gcollection

// Iterator interface for collection
// see LinkedBlockDequeIterator
type Iterator interface {
	HasNext() bool
	Next() interface{}
	Remove()
}

// InterruptedErr when deque block method bean interrupted will return this err
type InterruptedErr struct {
}

// NewInterruptedErr return new error instance
func NewInterruptedErr() *InterruptedErr {
	return &InterruptedErr{}
}

func (err *InterruptedErr) Error() string {
	return "Interrupted"
}