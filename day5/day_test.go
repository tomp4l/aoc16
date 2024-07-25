package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPassword(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Parallel()
	assert.Equal(t, "18f47a30", password("abc"))
}

func TestPassword2(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	t.Parallel()
	assert.Equal(t, "05ace8e3", password2("abc"))
}
