package config

const (
	PostgresEntDriver = "pgx"
	SqliteEntDriver   = "sqlite"
)

type Server struct {
	Port           string `env:"SERVE_PORT" default:"8080" flag:"-"`
	TLS            bool   `env:"TLS" default:"false" flag:"-"`
	TLSPort        string `env:"TLS_PORT" default:"443" flag:"-"`
	CertPath       string `env:"CERT_PATH" default:"invoiceexchange.local.pem" flag:"-"`
	CertPrivateKey string `env:"CERT_PRIVATE_KEY" default:"invoiceexchange.local-key.pem" flag:"-"`
}

type Config struct {
	Environment string `env:"ENVIRONMENT" default:"dev" flag:"-"`
	PostgresDSN string `env:"POSTGRES_DSN" default:"user=your_user password=your_password host=your_host port=your_port dbname=your_dbname sslmode=disable" yaml:"postgres-dsn" flag:"-"`

	EntDriver string `env:"ENT_DRIVER" default:"postgres" yaml:"ent-driver" flag:"-"`

	Server Server
}
