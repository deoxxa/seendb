package seendb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSeenDB(t *testing.T) {
	a := assert.New(t)

	db, err := New("./test.db")
	if !a.NoError(err) {
		return
	}

	if !a.NoError(db.Mark("testing")) {
		return
	}

	if !a.True(db.Seen("testing")) {
		return
	}
}
