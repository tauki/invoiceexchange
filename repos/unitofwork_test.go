package repos_test

import (
	"context"
	"github.com/tauki/invoiceexchange/repos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntUOWFactory_Commit(t *testing.T) {
	sqlMock, entClient, close := setupSQLMock(t)
	defer close()

	sqlMock.ExpectBegin()
	sqlMock.ExpectCommit()

	factory := repos.NewEntUOWFactory(entClient)
	uow, err := factory.New(context.Background())
	assert.NoError(t, err)
	assert.IsType(t, new(repos.EntWork), uow)

	err = uow.Commit()
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestEntUOWFactory_Rollback(t *testing.T) {
	sqlMock, entClient, close := setupSQLMock(t)
	defer close()

	sqlMock.ExpectBegin()
	sqlMock.ExpectRollback()

	factory := repos.NewEntUOWFactory(entClient)
	uow, err := factory.New(context.Background())
	assert.NoError(t, err)
	assert.IsType(t, new(repos.EntWork), uow)

	err = uow.RollBack()
	assert.NoError(t, err)
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestEntUOWFactory_RollbackUnlessCommitted(t *testing.T) {
	sqlMock, entClient, close := setupSQLMock(t)
	defer close()

	sqlMock.ExpectBegin()
	sqlMock.ExpectRollback()

	factory := repos.NewEntUOWFactory(entClient)
	uow, err := factory.New(context.Background())
	assert.NoError(t, err)
	assert.IsType(t, new(repos.EntWork), uow)

	uow.RollbackUnlessCommitted()
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}

func TestEntUOWFactory_RollbackUnlessCommitted_Committed(t *testing.T) {
	sqlMock, entClient, close := setupSQLMock(t)
	defer close()

	sqlMock.ExpectBegin()
	sqlMock.ExpectCommit()

	factory := repos.NewEntUOWFactory(entClient)
	uow, err := factory.New(context.Background())
	assert.NoError(t, err)
	assert.IsType(t, new(repos.EntWork), uow)

	err = uow.Commit()
	assert.NoError(t, err)

	uow.RollbackUnlessCommitted()
	assert.NoError(t, sqlMock.ExpectationsWereMet())
}
