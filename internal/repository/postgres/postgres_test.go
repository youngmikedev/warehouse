package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/youngmikedev/warehouse/internal/domain"
	"github.com/youngmikedev/warehouse/internal/repository/postgres/ent"

	_ "github.com/lib/pq"
)

const (
	dbuser        = "user"
	dbname        = "warehouse"
	password      = "example"
	dbhost        = "localhost"
	dbport        = "5432"
	dbExposedPort = "49194"

	testat = "access_token"
	testrt = "refresh_token"
)

var db *ent.Client

func TestMain(m *testing.M) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11-alpine",
		// ExposedPorts: []string{dbExposedPort},
		PortBindings: map[docker.Port][]docker.PortBinding{
			dbport + "/tcp": {{HostPort: dbExposedPort}},
		},
		Env: []string{
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_USER=" + dbuser,
			"POSTGRES_DB=" + dbname,
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		// set AutoRemove to true so that stopped container goes away by itself
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	hostAndPort := resource.GetHostPort(dbport + "/tcp")
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", dbuser, password, hostAndPort, dbname)

	if err = resource.Expire(120); err != nil { // Tell docker to hard kill the container in 120 seconds
		log.Fatalf("Could not set docker expire: %s", err)
	}
	var sqlDB *sql.DB

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	pool.MaxWait = 120 * time.Second
	if err = pool.Retry(func() error {
		sqlDB, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return sqlDB.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// init ent
	drv := entsql.OpenDB("postgres", sqlDB)
	db = ent.NewClient(ent.Driver(drv))

	// migrate schema
	ctx := context.TODO()
	if err := db.Schema.Create(ctx); err != nil {
		log.Fatalf("Failed creating schema resources: %v", err)
	}

	log.Println("Database connected")

	//Run tests
	code := m.Run()

	db.Close()

	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

// createUserAndGetTokens create account with default email and name if they are empty
func createUserAndGetTokens(t *testing.T, email, name string) (uid int, at, rt string) {
	if email == "" {
		email = "testemail@mail.com"
	}
	if name == "" {
		name = "Jon"
	}
	r := &UsersRepo{
		client: db,
	}
	id, err := r.Create(context.TODO(), domain.User{
		Name:  name,
		Email: email,
	}, "strongpass")
	if err != nil {
		t.Fatalf("createUserAndGetToken.UsersRepo.Create() error = %v", err)
	}

	_, err = r.CreateSession(context.TODO(), domain.Session{
		UserID:       id,
		AccessToken:  testat,
		RefreshToken: testrt,
		ExpiresAt:    time.Hour,
	})
	if err != nil {
		t.Fatalf("createUserAndGetToken.UsersRepo.CreateSession() error = %v", err)
	}

	return id, testat, testrt
}
