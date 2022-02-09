package postgres

import (
	"fmt"

	"github.com/imranzahaev/warehouse/internal/repository/postgres/ent"

	_ "github.com/lib/pq"
)

// NewClient established connection to a mongoDb instance using provided URI and auth credentials
func NewClient(host, port, user, dbname, password string) (*ent.Client, error) {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		host, port, user, dbname, password))
	if err != nil {
		return nil, err
	}

	return client, nil
}
