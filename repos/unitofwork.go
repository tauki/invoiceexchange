package repos

import (
	"context"
	"github.com/tauki/invoiceexchange/ent"
	"github.com/tauki/invoiceexchange/internal/errors"
	"github.com/tauki/invoiceexchange/internal/unitofwork"
)

type EntWork struct {
	Tx        *ent.Tx
	committed bool
}

func (u *EntWork) Commit() error {
	err := u.Tx.Commit()
	if err != nil {
		return errors.NewInfrastructureError("failed to commit transaction", err)
	}
	u.committed = true
	return nil
}

func (u *EntWork) RollBack() error {
	return u.Tx.Rollback()
}

func (u *EntWork) RollbackUnlessCommitted() {
	if !u.committed {
		u.RollBack()
	}
}

type EntUOWFactory struct {
	client *ent.Client
}

func NewEntUOWFactory(client *ent.Client) *EntUOWFactory {
	return &EntUOWFactory{client: client}
}

func (f *EntUOWFactory) New(ctx context.Context) (unitofwork.Work, error) {
	tx, err := f.client.Tx(ctx)
	if err != nil {
		return nil, errors.NewInfrastructureError("failed to create transaction", err)
	}
	return &EntWork{Tx: tx}, nil
}

func defaultClient(opts []unitofwork.Option, client *ent.Client) *ent.Client {
	for _, opt := range opts {
		work := opt()
		if entWork, ok := work.(*EntWork); ok {
			return entWork.Tx.Client()
		}
	}
	return client
}
