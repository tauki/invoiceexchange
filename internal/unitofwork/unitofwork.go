package unitofwork

import "context"

// Work represents a transaction coordinator
type Work interface {
	Commit() error
	RollBack() error
	RollbackUnlessCommitted()
}

// Factory represents a Unit of Work factory
type Factory interface {
	New(context.Context) (Work, error)
}

type Option func() Work

func With(w Work) Option {
	return func() Work {
		return w
	}
}
