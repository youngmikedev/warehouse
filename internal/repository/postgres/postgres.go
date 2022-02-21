package postgres

import (
	"context"
	"database/sql"
	"fmt"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/imranzahaev/warehouse/internal/repository/postgres/ent"

	_ "github.com/lib/pq"
)

// NewClient established connection to a mongoDb instance using provided URI and auth credentials
func NewClient(host, port, user, dbname, password string) (*ent.Client, error) {
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	sqlDB, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}

	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	drv := entsql.OpenDB("postgres", sqlDB)
	client := ent.NewClient(ent.Driver(drv))
	if err != nil {
		return nil, err
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, err
	}

	return client, nil
}
