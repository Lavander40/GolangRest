package store_test

import (
	"os"
	"testing"
)

var (
	databaseUrl string
)

func TestMain(m *testing.M) {
	databaseUrl = os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = "root:root@tcp(localhost:3306)/golang_users?multiStatements=true"
	}

	os.Exit(m.Run())
}
