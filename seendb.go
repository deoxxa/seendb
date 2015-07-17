package seendb // import "fknsrs.biz/p/seendb"

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type SeenDB struct {
	path string
	keys []string
}

func New(path string) (*SeenDB, error) {
	s := SeenDB{
		path: path,
	}

	return &s, s.load()
}

func (s *SeenDB) load() error {
	f, err := os.Open(s.path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}

		return err
	}
	defer f.Close()

	rd := bufio.NewReader(f)

	for {
		l, err := rd.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		l = strings.TrimSpace(l)

		if l == "" {
			continue
		}

		k, err := strconv.Unquote(strings.TrimSpace(l))
		if err != nil {
			return err
		}

		s.keys = append(s.keys, k)
	}

	return nil
}

func (s *SeenDB) Seen(key string) bool {
	for _, k := range s.keys {
		if k == key {
			return true
		}
	}

	return false
}

func (s *SeenDB) Mark(key string) error {
	if s.Seen(key) {
		return nil
	}

	f, err := os.OpenFile(s.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := fmt.Fprintf(f, "%q\n", key); err != nil {
		return err
	}

	s.keys = append(s.keys, key)

	return nil
}
