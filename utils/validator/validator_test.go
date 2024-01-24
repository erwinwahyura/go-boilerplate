package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	assert.NotNil(t, GetValidatorController())
}
