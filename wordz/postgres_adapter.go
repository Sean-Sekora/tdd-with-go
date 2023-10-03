package wordz

import (
	"context"
	"database/sql"
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "github.com/lib/pq"
)

func StartPostgres() (testcontainers.Container, error) {
	port := "5432"
	_, filename, _, _ := runtime.Caller(0)
	wd := filepath.Dir(filename)
	req := testcontainers.ContainerRequest{
		Image:        "postgres:16-alpine",
		ExposedPorts: []string{port},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "passwd",
		},
		Mounts: testcontainers.ContainerMounts{
			testcontainers.ContainerMount{
				Source: testcontainers.GenericBindMountSource{
					HostPath: fmt.Sprintf("%s/testdata/", wd),
				},
				Target: "/docker-entrypoint-initdb.d",
			},
		},
		Cmd:        []string{"postgres", "-c", "fsync=off"},
		WaitingFor: wait.ForExposedPort(),
	}

	ctx := context.Background()
	return testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
}

func ConnectToDB(port string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost", port, "postgres", "passwd", "postgres")
	return sql.Open("postgres", psqlInfo)
}

func GetOrmConection(t *testing.T, port string) *pg.DB {
	return pg.Connect(&pg.Options{
		Addr:     "localhost:" + port,
		User:     "postgres",
		Password: "passwd",
		Database: "postgres",
	})
}
