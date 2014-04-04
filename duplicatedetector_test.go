package duplicatedetector

import (
	"math/rand"
	"testing"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

const testServer = "localhost:11211"

func TestDeduper(t *testing.T) {
	c := initChecker()
	// adding first entry
	ok, _ := c.IsDuplicate("abc")
	if ok {
		t.Error("Cannot add entry #1 to cache. It was supposed to be the first time.")
	}

	// adding second entry
	ok, _ = c.IsDuplicate("xyz")
	if ok {
		t.Error("Cannot add entry #2 to cache. It was supposed to be the first time.")
	}

	// add same entry again, expected error
	ok, _ = c.IsDuplicate("abc")
	if !ok {
		t.Error("The checker failed to detect the duplicate entry #1")
	}

	// add same entry again, expected error
	ok, _ = c.IsDuplicate("xyz")
	if !ok {
		t.Error("The checker failed to detect the duplicate entry #2")
	}
}

func initChecker() *Checker {
	mc := memcache.New(testServer)

	rand.Seed(int64(time.Now().Nanosecond()))
	s := randStrings(5)
	prefix := s[0] + "."

	return NewChecker(mc, prefix, 3)
}

func randStrings(N int) []string {
	Maxlen := 10
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890abcdefghijklmnopqrstuvwxyz"

	r := make([]string, N)
	ri := 0
	buf := make([]byte, Maxlen)
	known := map[string]bool{}

	for i := 0; i < N; i++ {
	retry:
		l := rand.Intn(Maxlen)
		for j := 0; j < l; j++ {
			buf[j] = chars[rand.Intn(len(chars))]
		}
		s := string(buf[0:l])
		if known[s] {
			goto retry
		}
		known[s] = true
		r[ri] = s
		ri++
	}
	return r
}
