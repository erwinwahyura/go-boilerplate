package mongodb

import (
	"testing"
)

func TestNewMessageDB(t *testing.T) {
	NewMongoDsn("mongodb://localhost:27017", "test")
}
