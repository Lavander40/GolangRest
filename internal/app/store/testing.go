package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseUrl string) (*Store, func(...string)) {
	t.Helper()

	config := NewConfig()
	config.DatabaseUrl = databaseUrl
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := s.db.Exec(fmt.Sprintf("SET FOREIGN_KEY_CHECKS=0; TRUNCATE %s ; SET FOREIGN_KEY_CHECKS=1;", strings.Join(tables, "; TRUNCATE "))); err != nil {
				t.Fatal(err)
			}
		}

		s.Close()
	}
}
