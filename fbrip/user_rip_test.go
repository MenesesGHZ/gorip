package fbrip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSenseParameters(t *testing.T) {
	user := NewUserRip("fake_email", "fake_pass")
	assert.Equal(t, "fake_email", user.Parameters["email"])
	assert.Equal(t, "fake_pass", user.Parameters["pass"])
	assert.Equal(t, "", user.Parameters["lsd"])
	assert.Equal(t, "", user.Parameters["jazoest"])
	assert.Equal(t, "", user.Parameters["m_ts"])
	assert.Equal(t, "", user.Parameters["li"])
	assert.Equal(t, "", user.Parameters["try_number"])
	assert.Equal(t, "", user.Parameters["unrecognized_tries"])
	assert.Equal(t, "", user.Parameters["login"])
	user.Sense()
	assert.Equal(t, "fake_email", user.Parameters["email"])
	assert.Equal(t, "fake_pass", user.Parameters["pass"])
	assert.NotEqual(t, "", user.Parameters["lsd"])
	assert.NotEqual(t, "", user.Parameters["jazoest"])
	assert.NotEqual(t, "", user.Parameters["m_ts"])
	assert.NotEqual(t, "", user.Parameters["li"])
	assert.NotEqual(t, "", user.Parameters["try_number"])
	assert.NotEqual(t, "", user.Parameters["unrecognized_tries"])
	assert.NotEqual(t, "", user.Parameters["login"])
}
