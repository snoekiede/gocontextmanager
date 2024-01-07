package contextmanager

type ContextManager[T any] struct {
	Value T
}

func (cm *ContextManager[T]) Dispose(dispose func(a T)) {
	dispose(cm.Value)
}

func WithResource[T any, O any](value T, action func(a T) (O, error), dispose func(element T)) (O, error) {
	resource := &ContextManager[T]{Value: value}
	defer resource.Dispose(dispose)
	return action(resource.Value)
}
