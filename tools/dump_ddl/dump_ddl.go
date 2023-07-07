package main

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/tauki/invoiceexchange/ent"
	"go.uber.org/zap"
	"log"
	"os"
)

var (
	PostgresDSN string = "user=user password=password host=host port=port dbname=dbname sslmode=disable"
)

// Generates migration sql ddl based on the schema's current state in the
// specified database
func main() {
	log.SetOutput(os.Stderr)

	flag.StringVar(&PostgresDSN, "postgres-dsn", LookupEnvOrString("POSTGRES_DSN", PostgresDSN), "postgres_dsn")
	flag.Parse()

	log.Printf("app.config %v\n", getConfig(flag.CommandLine))

	conCfg, err := pgx.ParseConfig(PostgresDSN)
	if err != nil {
		log.Fatal("error parsing dsn", zap.Error(err))
	}

	db, err := sql.Open("pgx", conCfg.ConnString())
	if err != nil {
		log.Fatal("error opening db", zap.Error(err))
	}

	drv := entsql.OpenDB(dialect.Postgres, db)
	entClient := ent.NewClient(ent.Driver(drv) /*ent.Debug(), ent.Log(t.Log) */)
	defer entClient.Close()
	ctx := context.Background()
	// Dump migration changes to stdout.
	if err := entClient.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
}

func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func getConfig(fs *flag.FlagSet) []string {
	cfg := make([]string, 0, 10)
	fs.VisitAll(func(f *flag.Flag) {
		cfg = append(cfg, fmt.Sprintf("%s:%q", f.Name, f.Value.String()))
	})

	return cfg
}
