package repos_test

import (
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/tauki/invoiceexchange/ent"
	"testing"
)

func setupSQLMock(t *testing.T) (sqlmock.Sqlmock, *ent.Client, func() error) {
	db, sqlMock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	assert.NoError(t, err)
	drv := entsql.OpenDB(dialect.Postgres, db)

	entClient := ent.NewClient(ent.Driver(drv) /*ent.Debug(), ent.Log(t.Log) */)

	return sqlMock, entClient, db.Close
}
