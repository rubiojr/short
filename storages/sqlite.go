package storages

import (
	"fmt"
	"path/filepath"

	"github.com/rs/xid"
	"github.com/rubiojr/kv"
	_ "modernc.org/sqlite"
)

type Sqlite struct {
	db kv.Database
}

func (s *Sqlite) Init(root string) error {
	var err error
	dbroot := filepath.Join(root, "links.db")
	s.db, err = kv.New("sqlite", dbroot)
	return err
}

func (s *Sqlite) Code() string {
	return xid.New().String()
}

func (s *Sqlite) Save(url string) (string, error) {
	code := s.Code()

	fmt.Println("saving", code, "with", url)
	err := s.db.Set(code, []byte(url), nil)

	return code, err
}

func (s *Sqlite) Load(code string) (string, error) {
	fmt.Println("loading", code)
	c, err := s.db.Get(code)
	return string(c), err
}
