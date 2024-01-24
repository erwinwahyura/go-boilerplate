package ulid

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid"
)

// Generate UlidID with unix timestamp + entropy e.g 01G65Z755AFWAKHE12NY0CQ9FH
func GenerateUlidID() string {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())

	id, _ := ulid.New(ms, entropy)
	return id.String()
}
