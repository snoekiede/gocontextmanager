package contextmanager

type ContextManager[T any] struct {
	Context T
}

func (cm *ContextManager[T]) Dispose(dispose func(a T)) {
	dispose(cm.Context)
}

func WithResource[T any, O any](context T, action func(a T) (O, error), dispose func(element T)) (O, error) {
	resource := &ContextManager[T]{Context: context}
	defer resource.Dispose(dispose)
	return action(resource.Context)
}
